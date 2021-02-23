package parser

import (
	"dumper/model"
	"fmt"
	"math"
	"net/url"
	"strconv"
	"strings"
)

// MHPPTV类型数据解析器
//http://api.gy940830.com/ckx1/ppv.html?fr=31q&w=600&h=500&ppvdur=360.2&ppvsh=aliyun.vod.pptv.com&ppvrid=2b869057de8b418b38679e27a72ddc2b.mp4&ppvk=114d831855ea2781e93c864cb5f7bacc-a7db-1614063673%26segment%3D7d860603_7d860602_1614049273%26bppcataid%3D38&ppvvvid=EC74AB6E-8E90-4797-9D39-AA6E9EF86C04
// 数据格式：
//fr: 31q
//w: 600
//h: 500
//ppvdur: 360.2
//ppvsh: aliyun.vod.pptv.com
//ppvrid: 2b869057de8b418b38679e27a72ddc2b.mp4
//ppvk: 114d831855ea2781e93c864cb5f7bacc-a7db-1614063673&segment=7d860603_7d860602_1614049273&bppcataid=38
//ppvvvid: EC74AB6E-8E90-4797-9D39-AA6E9EF86C04

type MHPPTV struct{}

func (pv *MHPPTV) Parse(params string) (chunks []model.Chunk, ok bool) {
	if parseUrl, err := url.Parse(params); err == nil {
		ppvdur:=parseUrl.Query().Get("ppvdur")
		ppvsh:=parseUrl.Query().Get("ppvsh")
		ppvrid:=parseUrl.Query().Get("ppvrid")
		ppvk:=parseUrl.Query().Get("ppvk")
		ppvvvid:=parseUrl.Query().Get("ppvvvid")

		if ppvdur, err := strconv.ParseFloat(ppvdur, 64); err == nil {
			dur := int(math.Ceil(ppvdur))
			for i := 0; i <= dur; i++ {
				chunk := model.Chunk{
					Duration: "0",
					FileSize: "0",
					FilePath: fmt.Sprintf("http://%s/%d/%s?type=mhpptv&k=%s&vvid=%s", ppvsh, i, ppvrid, ppvk, ppvvvid),
					FileName: fmt.Sprintf("%d.%s", i, strings.Split(parseUrl.Query().Get("ppvrid"), ".")[1]),
				}
				chunks = append(chunks, chunk)
			}
			return chunks, true
		}
	}

	return chunks, false
}
