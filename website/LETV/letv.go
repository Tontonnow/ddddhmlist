package LETV

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Tontonnow/ddddhmlist/config"
	"github.com/Tontonnow/ddddhmlist/server"
	"github.com/Tontonnow/ddddhmlist/utils/sesssion"
	"github.com/wangluozhe/requests/url"
	"regexp"
	"strconv"
	"trpc.group/trpc-go/trpc-go/log"
)

var (
	web     = "letv"
	headers = map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.0.0",
		"referer":    "https://www.le.com",
	}
	ApiList      = "https://d-api-m.le.com/card/dynamic?platform=pc&callback=&vid=%s&cid=%s&pagesize=100&type=episode&isvip=1&page=1"
	ApiVideoInfo = "https://www.le.com/service/getVideoInfo/"
	client       = sesssion.NewClient(config.Conf.WebConfig[web])
)

type VideoInfo struct {
	Status string `json:"status"`
	Data   struct {
		IsPay int `json:"isPay"`
		Video struct {
			Vid           string `json:"vid"`
			Pid           int    `json:"pid"`
			Cid           string `json:"cid"`
			Title         string `json:"title"`
			Duration      string `json:"duration"`
			VideotypeName string `json:"videotypeName"`
			Description   string `json:"description"`
			CreateTime    string `json:"createTime"`
			VideoFrom     string `json:"video_from"`
			VidPlayCount  int    `json:"vid_Play_Count"`
		} `json:"video"`
		Album struct {
			Title        string `json:"title"`
			Description  string `json:"description"`
			PidPlayCount int    `json:"pid_Play_Count"`
			PlistScore   string `json:"plist_score"`
			VarietyShow  int    `json:"varietyShow"`
		} `json:"album"`
	} `json:"data"`
}

type ListInfo struct {
	Code string `json:"code"`
	Data struct {
		Position    int           `json:"position"`
		Location    int           `json:"location"`
		Periodpoint []interface{} `json:"periodpoint"`
		Otherlist   []interface{} `json:"otherlist"`
		Relalbum    []interface{} `json:"relalbum"`
		Relvideo    []interface{} `json:"relvideo"`
		Episode     struct {
			Cnt         int           `json:"cnt"`
			IsEnd       int           `json:"isEnd"`
			Now         int           `json:"now"`
			Allcnt      int           `json:"allcnt"`
			Upinfo      string        `json:"upinfo"`
			Yugao       []interface{} `json:"yugao"`
			Currentpage int           `json:"currentpage"`
			Videolist   []struct {
				Duration      string        `json:"duration"`
				Episode       string        `json:"episode"`
				Guest         interface{}   `json:"guest"`
				IsFirstLook   int           `json:"isFirstLook"`
				Key           int           `json:"key"`
				PayPlatforms  interface{}   `json:"payPlatforms"`
				Pic           string        `json:"pic"`
				ReleaseDate   string        `json:"releaseDate"`
				Singer        []interface{} `json:"singer"`
				Title         string        `json:"title"`
				Url           string        `json:"url"`
				Vid           int           `json:"vid"`
				VideoType     int           `json:"videoType"`
				SubTitle      string        `json:"subTitle"`
				WatchingFocus []interface{} `json:"watchingFocus"`
				Ispay         int           `json:"ispay"`
				Nextvid       interface{}   `json:"nextvid"`
			} `json:"videolist"`
		} `json:"episode"`
		Period []interface{} `json:"period"`
		List   []interface{} `json:"list"`
	} `json:"data"`
	Msg string `json:"msg"`
}

func GetVideoInfo(vid string) (info VideoInfo, err error) {
	req := url.NewRequest()
	req.Headers = url.ParseHeaders(headers)
	u := ApiVideoInfo + vid
	rsp, err := client.Do("GET", u, req)
	if err != nil {
		return
	}
	data, err := rsp.Json()
	if err != nil {
		return
	}
	switch data["data"].(type) {
	case string:
		err = fmt.Errorf(data["data"].(string))
		return
	default:
		err = json.Unmarshal(rsp.Content, &info)
		if err != nil {
			return
		}
	}
	return
}
func GetListInfo(vid, cid string) (info ListInfo, err error) {
	req := url.NewRequest()
	req.Headers = url.ParseHeaders(headers)
	u := fmt.Sprintf(ApiList, vid, cid)
	rsp, err := client.Do("GET", u, req)
	if err != nil {
		return
	}
	err = json.Unmarshal(rsp.Content, &info)
	return
}
func GetVId(url string) string {
	re := regexp.MustCompile(`/(\d+).`)
	productId := re.FindStringSubmatch(url)
	if len(productId) <= 1 {
		return ""
	}
	return productId[1]
}
func GetMateInfo(ctx context.Context, sharerUrl string) (r *server.Data, code int) {
	var (
		err error
	)
	r = &server.Data{}
	requestId := ctx.Value("requestId").(string)
	log.Debugf("site: %s, start GetMateInfo", web)
	vid := GetVId(sharerUrl)
	if vid == "" {
		code = 2
		err = fmt.Errorf("Invalid URL")
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		return
	}
	info, err := GetVideoInfo(vid)
	if err != nil {
		code = 1
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		return
	}
	r.SeriesTitle = info.Data.Album.Title
	r.SeriesId = strconv.Itoa(info.Data.Video.Pid)
	r.IsVip = info.Data.IsPay == 1
	r.IsSeries = false
	if info.Data.Video.Cid != "1" {
		r.IsSeries = true
		vid = info.Data.Video.Vid
		cid := info.Data.Video.Cid
		var listInfo ListInfo
		listInfo, err = GetListInfo(vid, cid)
		if err != nil {
			code = 1
			log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
			return
		}
		for _, v := range listInfo.Data.Episode.Videolist {
			r.VideoList = append(r.VideoList, &server.Video{
				EpisodeTitle: v.Title,
				EpisodeId:    strconv.Itoa(v.Vid),
				Episode:      uint32(v.Key),
				IsVip:        v.Ispay == 1,
			})
		}

	}
	return
}
