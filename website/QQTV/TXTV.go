package QQTV

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

var (
	web         = "TX"
	getPlayInfo = "https://tv.aiseet.atianqi.com/i-tvbin/qtv_video/get_play_info?cid=%s&Q-UA=PT%%3DTVMORE%%26CHID%%3D10009%%26VN%%3D1.9.0"
	client      = sesssion.NewClient(config.Conf.WebConfig[web])
)

type playInfo struct {
	Result struct {
		Ret      int    `json:"ret"`
		Msg      string `json:"msg"`
		CostTime int    `json:"cost_time"`
		Code     int    `json:"code"`
	} `json:"result"`
	Data struct {
		CID            string `json:"c_id"`
		CType          int    `json:"c_type"`
		CTitle         string `json:"c_title"`
		CSTitle        string `json:"c_s_title"`
		CHorizontalPic string `json:"c_horizontal_pic"`
		CVerticalPic   string `json:"c_vertical_pic"`
		CPayStatus     int    `json:"c_pay_status"`
		Imgtag         string `json:"imgtag"`
		Videos         []struct {
			VID          string            `json:"v_id"`
			VType        int               `json:"v_type"`
			VTitle       string            `json:"v_title"`
			VSTitle      string            `json:"v_s_title"`
			IsTrailer    int               `json:"is_trailer"`
			HeadTime     int               `json:"head_time"`
			TailTime     int               `json:"tail_time"`
			PlayStatus   int               `json:"play_status"`
			Tips         string            `json:"tips"`
			VPayStatus   int               `json:"v_pay_status"`
			UhdFlag      int               `json:"uhd_flag"`
			Pictures     map[string]string `json:"pictures"`
			NewPayStatus int               `json:"new_pay_status"`
			SquareTags   any               `json:"square_tags"`
		} `json:"videos"`
		CopyrightStatus int    `json:"copyright_status"`
		Tips            string `json:"tips"`
		QrcodeURL       string `json:"qrcode_url"`
		UhdFlag         int    `json:"uhd_flag"`
		Paid            int    `json:"paid"`
		VideoUIInfo     struct {
			VideoUIType       int  `json:"videoUIType"`
			VideoDataListType int  `json:"videoDataListType"`
			FreshPage         bool `json:"freshPage"`
		} `json:"videoUIInfo"`
	} `json:"data"`
}

func GetMateInfo(ctx context.Context, sharerUrl string) (r *server.Data, code int) {
	var err error
	requestId := ctx.Value("requestId").(string)
	log.Debugf("site: %s, start GetMateInfo", web)
	r = &server.Data{}
	cid := ExtractCid(sharerUrl)
	if cid == "" {
		err = fmt.Errorf("url is invalid")
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		code = 2
		return
	}
	getPlayInfoUrl := fmt.Sprintf(getPlayInfo, cid)
	res, err := client.Do("get", getPlayInfoUrl, nil)
	if err != nil {
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		code = 1
		return
	}
	var data playInfo
	err = json.Unmarshal(res.Content, &data)
	if err != nil {
		return
	}
	if data.Result.Ret != 0 {
		err = fmt.Errorf(data.Result.Msg)
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		code = 1
		return
	}
	r.SeriesTitle = data.Data.CTitle
	r.SeriesId = data.Data.CID
	for _, v := range data.Data.Videos {
		r.VideoList = append(r.VideoList, &server.Video{
			EpisodeId:    v.VID,
			EpisodeTitle: v.VTitle + " " + v.VSTitle,
			IsTrailer:    v.IsTrailer == 1,
			Extra: map[string]string{
				"uhd_flag": strconv.Itoa(v.UhdFlag),
			},
		})
	}
	return
}
func ExtractCid(u string) (cid string) {
	ud, err := url.Parse(u)
	if err != nil {
		log.Error("QQTV", "ExtractCid", err)
		return
	}
	cid = ud.Params.Get("cid")
	if cid == "" {
		ids := strings.Split(ud.Path, "/")
		if len(ids) < 4 {
			return
		}
		cid = ids[3]
	}
	return
}
