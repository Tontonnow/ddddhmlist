package Litv

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Tontonnow/ddddhmlist/config"
	"github.com/Tontonnow/ddddhmlist/server"
	"github.com/Tontonnow/ddddhmlist/utils/sesssion"
	"github.com/wangluozhe/requests/url"
	"strconv"
	"strings"
	"trpc.group/trpc-go/trpc-go/log"
)

type RpcData struct {
	Jsonrpc string                 `json:"jsonrpc,omitempty"`
	Id      int                    `json:"id"`
	Method  string                 `json:"method"`
	Params  interface{}            `json:"params"`
	Data    map[string]interface{} `json:"data,omitempty"`
}
type RpcRet struct {
	Jsonrpc string    `json:"jsonrpc"`
	Result  RpcResult `json:"result"`
	Error   RpcError  `json:"error"`
	Id      int       `json:"id"`
}
type RpcError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type RpcResult struct {
	DataVersion string `json:"data_version"`
	Data        struct {
		Assets []struct {
			AssetId   string        `json:"asset_id"`
			Bitrate   string        `json:"bitrate"`
			Quality   string        `json:"quality"`
			Subtitles []interface{} `json:"subtitles"`
		} `json:"assets"`
		HasSeasons            bool        `json:"has_seasons"`
		Seasons               []Season    `json:"seasons"`
		ProgramPublishPicType string      `json:"program_publish_pic_type"`
		Season                string      `json:"season"`
		Episode               string      `json:"episode"`
		Title                 string      `json:"title"`
		Picture               string      `json:"picture"`
		Score                 string      `json:"score"`
		Quality               string      `json:"quality"`
		Description           string      `json:"description"`
		OriginalTitle         interface{} `json:"original_title"`
		ContentId             string      `json:"content_id"`
		ContentType           string      `json:"content_type"`
		IsSeries              bool        `json:"is_series"`
		SeriesId              string      `json:"series_id"`
		SeasonName            string      `json:"season_name"`
		IsFinale              bool        `json:"is_finale"`
		EpisodeCount          string      `json:"episode_count"`
		ChargeMode            string      `json:"charge_mode"`
		VideoType             string      `json:"video_type"`
		VideoImage            string      `json:"video_image"`
		SeoKeyword            string      `json:"seo_keyword"`
		SeoDescription        string      `json:"seo_description"`
		CronDesc              string      `json:"cron_desc"`
	} `json:"data"`
}
type Season struct {
	Season        string   `json:"season"`
	Title         string   `json:"title"`
	SeasonName    string   `json:"season_name"`
	ContentId     string   `json:"content_id"`
	IsFinale      bool     `json:"is_finale"`
	EpisodeCount  string   `json:"episode_count"`
	PosterBanners []string `json:"poster_banners"`
	Episodes      []struct {
		Episode       string      `json:"episode"`
		EpisodeName   string      `json:"episode_name"`
		ContentId     string      `json:"content_id"`
		ChargeMode    string      `json:"charge_mode"`
		PosterBanners []string    `json:"poster_banners"`
		Copyright     []string    `json:"copyright"`
		VideoType     string      `json:"video_type"`
		VideoImage    string      `json:"video_image"`
		SecondaryMark string      `json:"secondary_mark"`
		OriginalDate  interface{} `json:"original_date"`
		GroupId       string      `json:"group_id"`
	} `json:"episodes"`
}

var (
	rpcApi = "https://proxy.svc.litv.tv/cdi/v2/rpc"
	web    = "LITV"
	client = sesssion.NewClient(config.Conf.WebConfig[web])
)

func NewRpcData(method string, params map[string]string) *RpcData {
	p := map[string]string{
		"version":     "3.0",
		"project_num": "LTAGP02",
		"client_id":   "",
		"device_id":   "",
		"swver":       "",
		"conditions":  "",
	}
	for k, v := range params {
		p[k] = v
	}
	data := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  method,
		"params":  p,
	}
	return &RpcData{
		Jsonrpc: "2.0",
		Id:      1,
		Method:  method,
		Params:  p,
		Data:    data,
	}
}

func (r *RpcData) Do() (data RpcRet, err error) {
	req := url.NewRequest()
	req.Headers = url.NewHeaders()
	req.Headers.Add("Content-Type", "application/json")
	req.Json = r.Data
	resp, err := client.Do("post", rpcApi, req)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp.Content, &data)
	if err != nil {
		return
	}
	if data.Error.Code != 0 {
		err = fmt.Errorf("code: %d, message: %s", data.Error.Code, data.Error.Message)
		return
	}
	return
}
func GetSeriesTree(seriesId string) (data []Season, err error) {
	r := NewRpcData("CCCService.GetSeriesTree", map[string]string{
		"series_id": seriesId,
	})
	ret, err := r.Do()
	if err != nil {
		return
	}
	data = ret.Result.Data.Seasons
	return
}

func GetProgramInformation(contentId string) (data RpcResult, err error) {
	r := NewRpcData("CCCService.GetProgramInformation", map[string]string{
		"content_id": contentId,
	})
	d, err := r.Do()
	if err != nil {
		return
	}
	if d.Result.Data.IsSeries {
		seasons, err := GetSeriesTree(d.Result.Data.SeriesId)
		if err != nil {
			return data, err
		}
		d.Result.Data.Seasons = seasons
	}
	data = d.Result
	return
}

func GetMateInfo(ctx context.Context, sharerUrl string) (r *server.Data, code int) {
	requestId := ctx.Value("requestId").(string)
	log.Debugf("site: %s, start GetMateInfo", web)
	r = &server.Data{}
	us := strings.Split(sharerUrl, "/")
	contentId := us[len(us)-1]
	d, err := GetProgramInformation(contentId)
	if err != nil {
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		code = 1
		return
	}
	data := d.Data
	Title := data.Title
	r.SeriesTitle = Title
	r.SeriesId = data.SeriesId
	r.IsVip = data.ChargeMode == "F"
	if data.IsSeries {
		r.IsSeries = true
		seas := data.Seasons
		for _, sea := range seas {
			SeasonName := sea.SeasonName
			episodes := sea.Episodes
			for _, episode := range episodes {
				SecondaryMark := episode.SecondaryMark
				e, _ := strconv.Atoi(episode.Episode)
				r.VideoList = append(r.VideoList, &server.Video{
					EpisodeTitle: SeasonName + " " + SecondaryMark,
					EpisodeId:    episode.ContentId,
					IsVip:        episode.ChargeMode == "F",
					Episode:      uint32(e),
				})
			}
		}
	} else {
		b, _ := json.Marshal(data.Assets)
		r.Extra = map[string]string{
			"assets": string(b),
		}
		r.IsSeries = false
	}
	return

}
