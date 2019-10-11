package main

import (
	"dumper/model"
	"dumper/parser"
	"dumper/tool"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const DIR string = "C:/tmp/00/" // 保存路径
const ThreadNum int = 5         // 线程数
//const MaxFileSize int = 100 * 1024 * 1024 // 文件最大内存（100MB）
const MaxFileSize int = 3 * 1024 * 1024 * 1024                                  // 文件最大内存（3GB）
const FFMPEG string = "C:/environment/ffmpeg-4.1.1-win64-static/bin/ffmpeg.exe" // FFMPEG命令地址

type Way struct {
	Id     string
	Name   string
	Parser parser.Parser
}

var ways = map[string]Way{
	"1": {Id: "1", Name: "ATV", Parser: new(parser.Atv)},
}

func init() {
	// 文件是否存在，不存在则创建
	if _, err := os.Stat(DIR); err != nil || os.IsNotExist(err) {
		if err = os.MkdirAll(DIR, os.ModePerm); err != nil {
			panic(err)
		}
	}
}

func main() {
	// 拿到输入数据
	pr, data := prepare()

	// 线程等待
	var wg sync.WaitGroup
	wg.Add(ThreadNum) // 增加线程数
	log.Println("开始任务")
	if chunks, ok := pr.Parse(data); ok {
		total := len(chunks)
		if chunkSize := int(math.Ceil(float64(total) / float64(ThreadNum))); chunkSize > 0 {
			for i := 0; i < ThreadNum; i++ { // 开启线程
				start := i * chunkSize
				end := start + chunkSize
				if end > total {
					end = total
				}
				go func(chunks *[]model.Chunk, start int, end int) {
					if start < total {
						for _, chunk := range (*chunks)[start:end] {
							_ = chunk.Download(DIR)
						}
					}
					wg.Done() // 操作完成，减少一个计数
				}(&chunks, start, end)
			}
		}
	}

	wg.Wait() // 等待，直到计数为0
	log.Println("完成下载")

	log.Println("开始合并")
	finish() // 生成清单文件
	ffmpeg() // 合并文件
	log.Println("完成合并")
}

func prepare() (pr parser.Parser, data string) {
	fmt.Println("选择类型：")
	for _, value := range ways {
		fmt.Println(value.Id + "）" + value.Name)
	}

	var way string
	for {
		fmt.Print("序号：")
		if _, err := fmt.Scanln(&way); err != nil {
			log.Println(err)
		}
		if value, ok := ways[way]; ok {
			pr = value.Parser
			break
		}
	}
	for {
		fmt.Printf("数据：")
		if _, err := fmt.Scanln(&data); err == nil {
			break
		}
	}
	return pr, data
}

func finish() {
	var totalSize int64
	var fileNames []string
	if dir, err := ioutil.ReadDir(DIR); err == nil {
		for _, fileInfo := range dir {
			if !fileInfo.IsDir() {
				totalSize += fileInfo.Size()
				fileNames = append(fileNames, fileInfo.Name())
			}
		}
	}

	sort.Sort(tool.FileNameList(fileNames))
	if fileCount := len(fileNames); fileCount > 0 {
		// 分片
		chunkNum := int(math.Ceil(float64(totalSize) / float64(MaxFileSize)))
		chunkSize := int(math.Ceil(float64(fileCount) / float64(chunkNum)))
		for i := 1; i <= chunkNum; i++ {
			// 创建新的文件
			if file, err := os.Create(DIR + "f" + strconv.Itoa(i) + ".txt"); err == nil {
				start := (i - 1) * chunkSize
				end := start + chunkSize
				if end > fileCount {
					end = fileCount
				}

				for _, value := range fileNames[start:end] {
					if _, err = file.WriteString("file '" + value + "'\n"); err != nil {
						panic(nil)
					}
				}
			}
		}
	}
}

func ffmpeg() {
	var fileNames []string
	if dir, err := ioutil.ReadDir(DIR); err == nil {
		for _, fileInfo := range dir {
			if !fileInfo.IsDir() && strings.HasSuffix(fileInfo.Name(), ".txt") {
				fileNames = append(fileNames, fileInfo.Name())
			}
		}
	}

	sort.Sort(tool.FileNameList(fileNames))
	for _, value := range fileNames {
		// 执行命令
		name := value[0:strings.Index(value, ".")]
		cmd := exec.Command("cmd", "/C", FFMPEG, "-f", "concat", "-i", DIR+value, "-c", "copy", DIR+name+".mp4")
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
		}
	}
}
