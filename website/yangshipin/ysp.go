package yangshipin

import (
	"context"
	"fmt"
	"github.com/Tontonnow/ddddhmlist/config"
	proto2 "github.com/Tontonnow/ddddhmlist/server"
	"github.com/Tontonnow/ddddhmlist/utils/sesssion"
	"github.com/wangluozhe/requests/url"
	"google.golang.org/protobuf/proto"
	"reflect"
	"strconv"
	"strings"
	"time"
	"trpc.group/trpc-go/trpc-go/log"
)

var (
	web       = "ysp"
	pageModel = "https://capi.yangshipin.cn/api/oms/pc/page/PG00000004"
	playInfo  = "https://csapi.yangshipin.cn/voapi/omsot/album/playinfo"
	yspEpg    = "https://capi.yangshipin.cn/api/yspepg/program/%s/%s"
	headers   = map[string]string{
		"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36 Edg/119.0.0.0",
	}
	Headers = url.ParseHeaders(headers)
	client  = sesssion.NewClient(config.Conf.WebConfig[web])
)

func doRequest(u string, respProto proto.Message, req *url.Request) error {
	if req == nil {
		req = url.NewRequest()
	}
	if req.Headers == nil {
		req.Headers = Headers
	}
	r, err := client.Do("get", u, req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	err = proto.Unmarshal(r.Content, respProto)
	if err != nil {
		return fmt.Errorf("unmarshal failed: %w", err)
	}

	respCode := reflect.ValueOf(respProto).Elem().FieldByName("Code").Uint() // Assumes the response has a Code field
	if respCode != 200 {
		respMessage := reflect.ValueOf(respProto).Elem().FieldByName("Message").String() // Assumes the response has a Message field
		return fmt.Errorf("error response: %d %s", respCode, respMessage)
	}

	return nil
}

func GetEpg(pid string, date ...string) []*CnYangshipinOmstvCommonProtoProgramModel_Program {
	if len(date) == 0 {
		date = append(date, time.Now().Format("20060102"))
	}
	u := fmt.Sprintf(yspEpg, pid, date[0])

	epg := &CnYangshipinOmstvCommonProtoEpgProgramModel_Response{}
	if err := doRequest(u, epg, nil); err != nil {
		log.Error(err)
		return nil
	}

	return epg.DataList
}

func GetPlayInfo(cid string) *CnYangshipinOmsCommonProtoAlbumPlayInfoModel_Data {
	tm := time.Now().Unix() / 5
	params := map[string]string{
		"cid": cid,
		"ts":  strconv.FormatInt(tm, 10),
	}

	req := url.NewRequest()
	req.Params = url.ParseParams(params)
	playInfoResp := &CnYangshipinOmsCommonProtoAlbumPlayInfoModel_Response{}
	if err := doRequest(playInfo, playInfoResp, req); err != nil {
		log.Error(err)
		return nil
	}

	return playInfoResp.Data
}
func GetEpgId() *CnYangshipinOmsCommonProtoPageModel_Data {
	pageModelResp := &CnYangshipinOmsCommonProtoPageModel_Response{}
	if err := doRequest(pageModel, pageModelResp, nil); err != nil {
		log.Error(err)
		return nil
	}

	return pageModelResp.Data
}
func GetMateInfo(ctx context.Context, sharerUrl string) (r *proto2.Data, code int) {
	var err error
	requestId := ctx.Value("requestId").(string)
	log.Debugf("site: %s, start GetMateInfo", web)
	r = &proto2.Data{}
	u := strings.Replace(sharerUrl, "/#/", "/o/", 1)
	ud, err := url.Parse(u)
	if err != nil {
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		code = 2
		return
	}
	cid := ud.Params.Get("cid")
	if cid == "" {
		err = fmt.Errorf("cid is empty")
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		code = 2
		return
	}
	playInfo := GetPlayInfo(cid)
	if playInfo == nil {
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		code = 1
		return
	}
	r.SeriesTitle = playInfo.Title
	r.SeriesId = playInfo.Cid
	r.IsVip = playInfo.IsVip
	for _, v := range playInfo.VideoList {
		r.VideoList = append(r.VideoList, &proto2.Video{
			EpisodeId:    v.Vid,
			EpisodeTitle: v.Title,
			Episode:      v.Episode,
			IsVip:        v.IsVip,
		})
	}
	return
}
