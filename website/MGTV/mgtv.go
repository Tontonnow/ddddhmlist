package MGTV

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Tontonnow/ddddhmlist/config"
	"github.com/Tontonnow/ddddhmlist/server"
	"github.com/Tontonnow/ddddhmlist/utils/sesssion"
	"github.com/wangluozhe/requests/url"
	"strings"
	"trpc.group/trpc-go/trpc-go/log"
)

var (
	web    = "MGTV"
	params = map[string]interface{}{
		"uuid":      "mgtvmac",
		"version":   "6.2.511.390.3.DANGBEI_TVAPP.0.0_Release",
		"device_id": "0",
		"mac_id":    strings.Replace(config.Conf.Mac, ":", "-", -1),
	}
	headers = map[string]string{
		"Accept-Encoding": "gzip",
		"Connection":      "Keep-Alive",
		"User-Agent":      "Dalvik/2.1.0 (Linux; U; Android 7.1.2; EU9-hi3716M Build/N2G48C)",
	}
	api    = "https://ott.bz.mgtv.com/ott/v2/video/info?uuid=mgtvmac&version=9&device_id=0&mac_id=" + strings.Replace(config.Conf.Mac, ":", "-", -1) + "&partId="
	client = sesssion.NewClient(config.Conf.WebConfig[web])
)

type MgRsp struct {
	Msg   string `json:"msg"`
	Code  int    `json:"code"`
	Data  MgData `json:"data"`
	Seqid string `json:"seqid"`
}
type MgData struct {
	Total      int `json:"total"`
	PageNo     int `json:"pageNo"`
	TotalPage  int `json:"totalPage"`
	PageSize   int `json:"pageSize"`
	ExtraCount int `json:"extraCount"`
	Rows       []struct {
		ClipId           string `json:"clipId"`
		CornerLabelStyle struct {
			Color string `json:"color"`
			Font  string `json:"font"`
		} `json:"cornerLabelStyle"`
		Desc                  string `json:"desc"`
		Duration              int    `json:"duration"`
		Image                 string `json:"image"`
		Index                 int    `json:"index"`
		IsPositive            int    `json:"isPositive"`
		Name                  string `json:"name"`
		OttFileFormatDuration string `json:"ottFileFormatDuration"`
		PlId                  string `json:"plId"`
		Title                 string `json:"title"`
		Type                  int    `json:"type"`
		VideoId               string `json:"videoId"`
	} `json:"rows"`
	AbrConfig struct {
		CheckInterval int    `json:"checkInterval"`
		SelectAbr     bool   `json:"selectAbr"`
		EnhanceAPI    string `json:"enhanceAPI"`
		Fts           string `json:"fts"`
		SwithLimit    struct {
			Up   int `json:"up"`
			Down int `json:"down"`
		} `json:"swithLimit"`
		EnableAbr        bool `json:"enableAbr"`
		MaxIdleQuality   int  `json:"maxIdleQuality"`
		MinCheckInterval int  `json:"minCheckInterval"`
		MinIdleQuality   int  `json:"minIdleQuality"`
		Algorithm        int  `json:"algorithm"`
	} `json:"abrConfig"`
	AbrSources []interface{} `json:"abrSources"`
	Attach     struct {
		Defs []struct {
			StandardName string `json:"standardName"`
			Framerate    string `json:"framerate"`
			Name         string `json:"name"`
			DefLevel     int    `json:"defLevel"`
			Filebitrate  string `json:"filebitrate"`
			Type         int    `json:"type"`
			Playable     int    `json:"playable"`
		} `json:"defs"`
		Default  int   `json:"default"`
		Force    int   `json:"force"`
		Displays []int `json:"displays"`
		Points   []struct {
			PointTitle  string `json:"pointTitle"`
			PointPic    string `json:"pointPic"`
			PointStart  int    `json:"pointStart"`
			PartPointId int    `json:"partPointId"`
			PointId     int    `json:"pointId"`
			PartId      int    `json:"partId"`
			HasMppFile  int    `json:"hasMppFile"`
			PointType   int    `json:"pointType"`
			HasOttFile  int    `json:"hasOttFile"`
			ImgHash     string `json:"imgHash"`
		} `json:"points"`
	} `json:"attach"`
	CategoryList []struct {
		Code       string `json:"code"`
		Datatype   int    `json:"datatype"`
		Isrefresh  int    `json:"isrefresh"`
		More       string `json:"more"`
		ObjectType int    `json:"objectType"`
		Playorder  int    `json:"playorder"`
		Position   int    `json:"position"`
		Showtype   int    `json:"showtype"`
		Title      string `json:"title"`
		Url        string `json:"url"`
	} `json:"categoryList"`
	ClipId            string   `json:"clipId"`
	ClipImage         string   `json:"clipImage"`
	ClipName          string   `json:"clipName"`
	ClipStatus        string   `json:"clipStatus"`
	ClipVerticalImage string   `json:"clipVerticalImage"`
	Detail            []string `json:"detail"`
	Duration          int      `json:"duration"`
	FitAge            string   `json:"fitAge"`
	FstlvlId          string   `json:"fstlvlId"`
	FstlvlName        string   `json:"fstlvlName"`
	Index             int      `json:"index"`
	Info              string   `json:"info"`
	IsIntact          int      `json:"isIntact"`
	KeepPlay          struct {
		Duration  int    `json:"duration"`
		Info      string `json:"info"`
		Vid       string `json:"vid"`
		WatchTime int    `json:"watchTime"`
	} `json:"keepPlay"`
	KeepPlayType     int    `json:"keepPlayType"`
	Kind             string `json:"kind"`
	ModuleDetailInfo struct {
		UModuleId         string `json:"uModuleId"`
		ProgressBarEffect string `json:"progressBarEffect"`
		WhiteBgUrl        string `json:"whiteBgUrl"`
		ProgressBarColor  string `json:"progressBarColor"`
		ChannelAllModules []struct {
			ModuleType string `json:"moduleType"`
			OrderNum   int    `json:"orderNum"`
			ModuleId   string `json:"moduleId"`
		} `json:"channelAllModules"`
		HitType                  string `json:"hitType"`
		QiJingBrandImg           string `json:"qiJingBrandImg"`
		BlackBgUrl               string `json:"blackBgUrl"`
		VclassId                 string `json:"vclassId"`
		QiJingFeatureFlag        string `json:"qiJingFeatureFlag"`
		RectangleProgressBarLogo string `json:"rectangleProgressBarLogo"`
		CategoryList             []struct {
			Datatype string `json:"datatype"`
			Position string `json:"position"`
		} `json:"categoryList"`
		TopVclassId        string `json:"topVclassId"`
		ProgressBarLogoUrl string `json:"progressBarLogoUrl"`
	} `json:"moduleDetailInfo"`
	PlId            string `json:"plId"`
	PlImage         string `json:"plImage"`
	PlName          string `json:"plName"`
	PlVerticalImage string `json:"plVerticalImage"`
	Play            string `json:"play"`
	PublishYear     string `json:"publishYear"`
	RecState        string `json:"recState"`
	RecommendInfo   struct {
		ModuleNum  string `json:"moduleNum"`
		ModuleName string `json:"moduleName"`
		ModuleId   string `json:"moduleId"`
	} `json:"recommendInfo"`
	RelatedPlay interface{} `json:"relatedPlay"`
	Serialno    string      `json:"serialno"`
	SerieTabs   []struct {
		ClipID   string `json:"clipId"`
		Current  int    `json:"current"`
		SeasonID string `json:"seasonId"`
		Title    string `json:"title"`
		URL      string `json:"url"`
	} `json:"serieTabs"`
	SeriesId           string `json:"seriesId"`
	ShowMode           int    `json:"showMode"`
	ShowTitle          string `json:"showTitle"`
	TotalNumber        string `json:"totalNumber"`
	TotalSize          int    `json:"totalSize"`
	UpdateDesc         string `json:"updateDesc"`
	UpdateInfo         string `json:"updateInfo"`
	VclassId           string `json:"vclassId"`
	VideoId            string `json:"videoId"`
	VideoImage         string `json:"videoImage"`
	VideoName          string `json:"videoName"`
	VideoVerticalImage string `json:"videoVerticalImage"`
	VipInfo            struct {
		Preview     int           `json:"preview"`
		VipDefs     []int         `json:"vip_defs"`
		PRange      int           `json:"p_range"`
		ChargeFirms []interface{} `json:"charge_firms"`
		Hdcp        int           `json:"hdcp"`
		Mark        int           `json:"mark"`
	} `json:"vipInfo"`
	VipInfoOtt struct {
		Preview     int           `json:"preview"`
		VipDefs     []int         `json:"vip_defs"`
		PRange      int           `json:"p_range"`
		ChargeFirms []interface{} `json:"charge_firms"`
		Hdcp        int           `json:"hdcp"`
		Mark        int           `json:"mark"`
	} `json:"vipInfoOtt"`
	VipMark struct {
		VipDefs     []int         `json:"vip_defs"`
		ChargeFirms []interface{} `json:"charge_firms"`
		Mark        int           `json:"mark"`
	} `json:"vipMark"`
	VipMarkOtt struct {
		VipDefs     []int         `json:"vip_defs"`
		ChargeFirms []interface{} `json:"charge_firms"`
		Mark        int           `json:"mark"`
	} `json:"vipMarkOtt"`
	WatchTime int `json:"watchTime"`
}

func doRequest(u string, data map[string]interface{}) (re MgData, err error) {
	Req := url.NewRequest()
	Req.Headers = url.NewHeaders()
	for k, v := range headers {
		Req.Headers.Add(k, v)
	}
	Req.Json = data
	r, err := client.Do("get", u, Req)
	if err != nil {
		log.Error("MGTV", "doRequest", err)
		return
	}
	var rsp MgRsp
	err = json.Unmarshal(r.Content, &rsp)
	if err != nil {
		log.Error("MGTV", "doRequest", err)
		return
	}
	if rsp.Code != 200 {
		err = fmt.Errorf(rsp.Msg)
		log.Error("MGTV", "doRequest", err)
		return
	}
	re = rsp.Data
	return
}

func GetMateInfo(ctx context.Context, sharerUrl string) (ret *server.Data, code int) {
	requestId := ctx.Value("requestId").(string)
	log.Debugf("site: %s, start GetMateInfo", web)
	ret = &server.Data{}
	var pageSize, pageNo int
	pageSize = 100
	pageNo = 1
	b, _, _ := strings.Cut(sharerUrl, ".html")
	bs := strings.Split(b, "/")
	cid := bs[len(bs)-1]
	u := api + cid
	r, err := doRequest(u, nil)
	if err != nil {
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		code = 1
		return
	}
	ret.SeriesTitle = r.ClipName
	ret.SeriesId = r.ClipId
	categoryList := r.CategoryList
	/*如果只需要多季正片可以取serieTabs
		for _, v := range r.SerieTabs {
		u := v.URL + "&pageSize=" + fmt.Sprint(pageSize) + "&pageNo=" + fmt.Sprint(pageNo)
	}
	*/
	for _, v := range categoryList {
		if v.Code == "formal" { //默认只取正片 其余的可以自行添加 可能需要修改结构体
			u := v.Url + "&pageSize=" + fmt.Sprint(pageSize) + "&pageNo=" + fmt.Sprint(pageNo)
			r, err = doRequest(u, nil)
			if err != nil {
				log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
				code = 1
				return
			}
			ret.Extra = map[string]string{
				"total":     fmt.Sprint(r.Total),
				"totalPage": fmt.Sprint(r.TotalPage)}

			for _, v := range r.Rows {
				if v.Type == 8 || v.VideoId == "" {
					continue
				}
				ret.VideoList = append(ret.VideoList, &server.Video{
					EpisodeTitle: v.Name,
					EpisodeId:    v.VideoId,
					IsVip:        v.CornerLabelStyle.Font == "VIP",
					IsTrailer:    v.Type != 1,
				})
			}
		}
	}
	return
}
