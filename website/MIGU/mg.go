package MIGU

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

var headers = map[string]string{
	"APP-VERSION-CODE": "81170308",
	"appVersion":       "0",
	"areaId":           "280",
	"carrierCode":      "CT",
	"channel":          "0148",
	"clientCityId":     "0280",
	"User-Agent":       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36",
}
var (
	web    = "migu"
	client = sesssion.NewClient(config.Conf.WebConfig[web])
)

type MgRsp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Body    struct {
		Data struct {
			Actor        string `json:"actor"`
			Area         string `json:"area"`
			AssetID      string `json:"assetID"`
			EpsAssetID   string `json:"epsAssetID"`
			EpsID        string `json:"epsID"`
			Year         string `json:"year"`
			ContentStyle string `json:"contentStyle"`
			Director     string `json:"director"`
			Detail       string `json:"detail"`
			EpsCount     int    `json:"epsCount"`
			Name         string `json:"name"`
			ProgramType  string `json:"programType"`
			Score        string `json:"score"`
			PublishTime  string `json:"publishTime"`
			StateTime    string `json:"stateTime"`
			UpdateEP     string `json:"updateEP"`
			Label4K      string `json:"label_4K"`
			Tip          struct {
				Code string `json:"code"`
				Msg  string `json:"msg"`
			} `json:"tip"`
			DownloadTip struct {
				Code string `json:"code"`
				Msg  string `json:"msg"`
			} `json:"downloadTip"`
			Playing struct {
				PID          string `json:"pID"`
				Name         string `json:"name"`
				Duration     string `json:"duration"`
				Index        string `json:"index"`
				LimitFreeTip struct {
					FreeLimitDates []struct {
						BeginDate string `json:"beginDate"`
						EndDate   string `json:"endDate"`
					} `json:"freeLimitDates"`
				} `json:"limitFreeTip"`
				LimitedTimeTips []struct {
					LimitedTime []struct {
						BeginDate string `json:"beginDate"`
						EndDate   string `json:"endDate"`
					} `json:"limitedTime"`
				} `json:"limitedTimeTips"`
				Label4K string `json:"label_4K"`
				ReqType string `json:"reqType"`
				CdnData string `json:"cdnData"`
			} `json:"playing"`
			Datas []struct {
				LimitFreeTip struct {
					FreeLimitDates []struct {
						BeginDate string `json:"beginDate"`
						EndDate   string `json:"endDate"`
					} `json:"freeLimitDates"`
				} `json:"limitFreeTip"`
				LimitedTimeTips []struct {
					LimitedTime []struct {
						BeginDate string `json:"beginDate"`
						EndDate   string `json:"endDate"`
					} `json:"limitedTime"`
				} `json:"limitedTimeTips"`
				Tip struct {
					Code string `json:"code"`
					Msg  string `json:"msg"`
				} `json:"tip"`
				DownloadTip struct {
					Code string `json:"code"`
					Msg  string `json:"msg"`
				} `json:"downloadTip"`
				Duration      string `json:"duration"`
				Name          string `json:"name"`
				Detail        string `json:"detail"`
				PID           string `json:"pID"`
				ProgramTypeV2 string `json:"programTypeV2"`
				Pics          struct {
					LowResolutionH    string `json:"lowResolutionH"`
					LowResolutionV    string `json:"lowResolutionV"`
					LowResolutionV34  string `json:"lowResolutionV34"`
					HighResolutionH   string `json:"highResolutionH"`
					HighResolutionV   string `json:"highResolutionV"`
					HighResolutionV34 string `json:"highResolutionV34"`
				} `json:"pics"`
				H5Pics struct {
					LowResolutionH    string `json:"lowResolutionH"`
					LowResolutionV    string `json:"lowResolutionV"`
					LowResolutionV34  string `json:"lowResolutionV34"`
					HighResolutionH   string `json:"highResolutionH"`
					HighResolutionV   string `json:"highResolutionV"`
					HighResolutionV34 string `json:"highResolutionV34"`
				} `json:"h5pics"`
				AssetID    string `json:"assetID"`
				Label4K    string `json:"label_4K"`
				VideoType  string `json:"videoType"`
				Resolution struct {
					MediaWidth  int `json:"mediaWidth"`
					MediaHeight int `json:"mediaHeight"`
				} `json:"resolution"`
				Index       string `json:"index"`
				CopyRightVo struct {
					Terminal  string `json:"terminal"`
					Area      string `json:"area"`
					BeginDate string `json:"beginDate"`
					EndDate   string `json:"endDate"`
					Way       string `json:"way"`
				} `json:"copyRightVo"`
				Way           string `json:"way"`
				DisplayType   string `json:"displayType"`
				IsUFC         string `json:"isUFC"`
				IsPrevue      string `json:"isPrevue"`
				ReqType       string `json:"reqType"`
				CdnData       string `json:"cdnData"`
				IsHDRVivid    string `json:"isHDRVivid"`
				ContentStatus string `json:"contentStatus"`
				KEYWORDS      string `json:"KEYWORDS"`
			} `json:"datas"`
			Star []struct {
				Career string `json:"career"`
				Img    string `json:"img"`
				Name   string `json:"name"`
				StarId string `json:"starId"`
				Action struct {
					Type              string `json:"type"`
					IosMinVersion     string `json:"iosMinVersion"`
					AndroidMinVersion string `json:"androidMinVersion"`
					Params            struct {
						FrameID   string `json:"frameID"`
						PageID    string `json:"pageID"`
						ContentID string `json:"contentID"`
						Location  string `json:"location"`
						Extra     struct {
							StarID string `json:"starID"`
						} `json:"extra"`
					} `json:"params"`
				} `json:"action"`
			} `json:"star"`
			PrdPackId string `json:"prdPackId"`
			Pics      struct {
				LowResolutionH    string `json:"lowResolutionH"`
				LowResolutionV    string `json:"lowResolutionV"`
				LowResolutionV34  string `json:"lowResolutionV34"`
				HighResolutionH   string `json:"highResolutionH"`
				HighResolutionV   string `json:"highResolutionV"`
				HighResolutionV34 string `json:"highResolutionV34"`
			} `json:"pics"`
			Way          string `json:"way"`
			TotalPage    string `json:"totalPage"`
			TotalCount   string `json:"totalCount"`
			PageSize     string `json:"pageSize"`
			PricingStage string `json:"pricing_stage"`
			H5Pics       struct {
				LowResolutionH    string `json:"lowResolutionH"`
				LowResolutionV    string `json:"lowResolutionV"`
				LowResolutionV34  string `json:"lowResolutionV34"`
				HighResolutionH   string `json:"highResolutionH"`
				HighResolutionV   string `json:"highResolutionV"`
				HighResolutionV34 string `json:"highResolutionV34"`
			} `json:"h5pics"`
			MpId          string `json:"mpId"`
			ProgramTypeV2 string `json:"programTypeV2"`
			CopyRightVo   struct {
				Terminal  string `json:"terminal"`
				Area      string `json:"area"`
				BeginDate string `json:"beginDate"`
				EndDate   string `json:"endDate"`
				Way       string `json:"way"`
			} `json:"copyRightVo"`
			Directors []struct {
				Img    string `json:"img"`
				Name   string `json:"name"`
				StarId string `json:"starId"`
				Action struct {
					Type              string `json:"type"`
					IosMinVersion     string `json:"iosMinVersion"`
					AndroidMinVersion string `json:"androidMinVersion"`
					Params            struct {
						FrameID   string `json:"frameID"`
						PageID    string `json:"pageID"`
						ContentID string `json:"contentID"`
						Location  string `json:"location"`
						Extra     struct {
							StarID string `json:"starID"`
						} `json:"extra"`
					} `json:"params"`
				} `json:"action"`
			} `json:"directors"`
			VideoType      string `json:"videoType"`
			DisplayType    string `json:"displayType"`
			IsUFC          string `json:"isUFC"`
			Recommendation string `json:"recommendation"`
			ProgramForm    string `json:"programForm"`
			Language       string `json:"language"`
			IsHDRVivid     string `json:"isHDRVivid"`
			IsEnd          string `json:"isEnd"`
			KEYWORDS       string `json:"KEYWORDS"`
		} `json:"data"`
	} `json:"body"`
	TimeStamp   int64 `json:"timeStamp"`
	TimeTraceVO struct {
		RequestTime  string `json:"requestTime"`
		ResponseTime string `json:"responseTime"`
	} `json:"timeTraceVO"`
}

func GetMateInfo(ctx context.Context, sharerUrl string) (r *server.Data, code int) {
	requestId := ctx.Value("requestId").(string)
	log.Debugf("site: %s, start GetMateInfo", web)
	r = &server.Data{}
	b, _, _ := strings.Cut(sharerUrl, "?")
	bs := strings.Split(b, "/")
	cid := bs[len(bs)-1]
	u := fmt.Sprintf("https://v2-sc.miguvideo.com/program/v3/cont/content-info/%s", cid)
	Req := url.NewRequest()
	Req.Headers = url.NewHeaders()
	for k, v := range headers {
		Req.Headers.Add(k, v)
	}
	re, err := client.Do("get", u, Req)
	if err != nil {
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		return
	}
	var rsp MgRsp
	err = json.Unmarshal(re.Content, &rsp)
	if err != nil {
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		code = 1
		return
	}
	if rsp.Code != 200 {
		err = fmt.Errorf(rsp.Message)
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		code = 1
		return
	}
	r.SeriesTitle = rsp.Body.Data.Name
	r.SeriesId = rsp.Body.Data.AssetID
	if rsp.Body.Data.Datas != nil {
		r.IsSeries = true
		for _, v := range rsp.Body.Data.Datas {
			e, _ := strconv.Atoi(v.Index)
			r.VideoList = append(r.VideoList, &server.Video{
				EpisodeId:    v.PID,
				EpisodeTitle: v.Name,
				Extra: map[string]string{
					"label_4K": rsp.Body.Data.Label4K,
					"HDRVivid": rsp.Body.Data.IsHDRVivid,
				},
				Episode:   uint32(e),
				IsTrailer: v.IsPrevue == "1",
				IsVip:     v.Tip.Code == "VIP" || v.Tip.Code == "USE_TICKET",
			})
		}
	} else {
		r.IsSeries = false
		r.VideoList = append(r.VideoList, &server.Video{
			EpisodeId:    rsp.Body.Data.Playing.PID,
			EpisodeTitle: rsp.Body.Data.Playing.Name,
			Extra: map[string]string{
				"label_4K": rsp.Body.Data.Label4K,
				"HDRVivid": rsp.Body.Data.IsHDRVivid,
			},
		})
	}
	return
}
