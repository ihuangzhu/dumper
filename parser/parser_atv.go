package parser

import (
	"dumper/model"
	"encoding/json"
	"net/url"
	"strings"
)

// ATV类型数据解析器
// 数据格式：
// 	{
// 		"autoplay":true,
// 		"video":[{
// 			"video":[
// 				{
// 					"file":"http://183.2.199.27/0/c256526e9f834ee5f0229e08e778f8e1.mp4?type=atv&k=aa03100bae81e1f633820f7a17a8a24c-5433-1554829225%26segment%3Db7eedbf1_b7eedadd_1554814825%26bppcataid%3D38",
// 					"duration":9.76,
// 					"bytesTotal":8952118
// 				},
// 				...
// 			]
// 		}]
// 	}

type Video struct {
	File       string  // 文件路径
	Duration   string // 时长（秒）
	BytesTotal string  // 文件大小
}

type Atv struct {
	Autoplay bool
	Video    []map[string][]Video
}

func (pv *Atv) Parse(params string) (chunks []model.Chunk, ok bool) {
	var atv Atv
	decoder := json.NewDecoder(strings.NewReader(params))
	if err := decoder.Decode(&atv); err == nil {
		for _, valueL1 := range atv.Video {
			for _, valueL2 := range valueL1 {
				for _, valueL3 := range valueL2 {
					var fileName string
					if parseUrl, err := url.Parse(valueL3.File); err == nil {
						if pathArr := strings.Split(strings.Trim(parseUrl.Path, "/"), "/"); len(pathArr) > 1 {
							var suffix string = ""
							if i := strings.LastIndex(parseUrl.Path, "."); i != -1 {
								suffix = parseUrl.Path[i:]
							}
							fileName = pathArr[0] + suffix
						}
					} else {
						continue
					}

					chunk := model.Chunk{
						Duration: valueL3.Duration,
						FileSize: valueL3.BytesTotal,
						FilePath: valueL3.File,
						FileName: fileName,
					}
					chunks = append(chunks, chunk)
				}
			}
		}
		ok = true
	} else {
		panic(err)
	}

	return chunks, ok
}
