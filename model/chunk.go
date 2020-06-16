package model

import (
	"io"
	"net/http"
	"os"
)

type Chunk struct {
	Duration string // 时长（秒）
	FileSize string  // 文件大小
	FilePath string  // 文件路径
	FileName string  // 文件名称
}

func (c *Chunk) Download(target string) error {
	// 发起GET请求
	resp, err := http.Get(c.FilePath)
	if err != nil || resp.StatusCode != 200 {
		return err
	}
	defer resp.Body.Close()

	// 创建文件
	out, err := os.Create(target + c.FileName)
	if err != nil {
		return err
	}
	defer out.Close()

	// 读取写入
	_, err = io.Copy(out, resp.Body)
	return err
}
