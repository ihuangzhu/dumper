package parser

import (
	"bufio"
	"dumper/model"
	"net/http"
	"strings"
)

// VOD类型数据解析器
//https://apd-8a581ef6630150e0d2fb0a11906a307e.v.smtcdns.com/sportsts.tc.qq.com/AG3Xf2UjoKESrlhyoN94Ham-up5hDEaNxmUE6TcW5BWc/uwMROfz2r5zAoaQXGdGlC2df64jGQbD41Y-gLg4o93LRZ4It/cse0hX-qahFdJTyFaFkAlMKWke7UFNbP-O5ofUVLbiu9Q57RKP11Lsaud-VXauwZOygb5ooiR_05pY7BgLIdgF5H9y_c7KqLPB-u_oMcuoT0eioh_cHwlcHU0cPpXwJl9GFkJ3dEntEyc0syFgJOG7O1bV03Vxuy9uCni4J4IqY/k00363fx4lw.321004.ts.m3u8?ver=4
//https://apd-8a581ef6630150e0d2fb0a11906a307e.v.smtcdns.com/sportsts.tc.qq.com/AG3Xf2UjoKESrlhyoN94Ham-up5hDEaNxmUE6TcW5BWc/uwMROfz2r5zAoaQXGdGlC2df64jGQbD41Y-gLg4o93LRZ4It/cse0hX-qahFdJTyFaFkAlMKWke7UFNbP-O5ofUVLbiu9Q57RKP11Lsaud-VXauwZOygb5ooiR_05pY7BgLIdgF5H9y_c7KqLPB-u_oMcuoT0eioh_cHwlcHU0cPpXwJl9GFkJ3dEntEyc0syFgJOG7O1bV03Vxuy9uCni4J4IqY/01_k00363fx4lw.321004.1.ts?index=1&start=12000&end=24000&brs=5086340&bre=9927339&ver=4&token=d2543289c78146bc1b36b3da84e74895
// 数据格式：
//#EXTM3U
//#EXT-X-VERSION:3
//#EXT-X-MEDIA-SEQUENCE:0
//#EXT-X-TARGETDURATION:13
//#EXT-X-PLAYLIST-TYPE:VOD
//#EXTINF:12.000,
//00_k00363fx4lw.321004.1.ts?index=0&start=0&end=12000&brs=0&bre=5086339&ver=4&token=7d85c1cdc4d6046c7e64dd3e09fa3d27
//#EXTINF:12.000,
//01_k00363fx4lw.321004.1.ts?index=1&start=12000&end=24000&brs=5086340&bre=9927339&ver=4&token=d2543289c78146bc1b36b3da84e74895
//#EXTINF:12.000,
//01_k00363fx4lw...

type Vod struct{}

func (pv *Vod) Parse(params string) (chunks []model.Chunk, ok bool) {
	// 下载VOD文件
	resp, err := http.Get(params)
	if err != nil || resp.StatusCode != 200 {
		return chunks, false
	}
	defer resp.Body.Close()

	// 主链接
	var paths []string
	paths = strings.Split(params, "/")
	paths = paths[:len(paths) - 1]

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		if !strings.HasPrefix(scanner.Text(), "#") {
			pathLocal := append(paths, scanner.Text())

			chunk := model.Chunk{
				Duration: "0",
				FileSize: "0",
				FilePath: strings.Join(pathLocal, "/"),
				FileName: strings.Split(scanner.Text(), "?")[0],
			}
			chunks = append(chunks, chunk)
		}
	}

	return chunks, true
}
