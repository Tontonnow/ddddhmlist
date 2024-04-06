package QQTV

import (
	"bytes"
	"context"
	"fmt"
	proto2 "github.com/Tontonnow/ddddhmlist/server"
	"github.com/wangluozhe/requests/url"
	"google.golang.org/protobuf/proto"
	"strconv"
	"strings"
	"trpc.group/trpc-go/trpc-go/log"
)

type TxTrpc struct {
	QQLiveHead *RequestHead
	RequestId  int
	Caller     string
	Headers    map[string]string
	videoList  []*VideoItemData
}

var (
	CountryCode = []int32{153514, 153505, 153564, 153513, 153548, 153560, 153569, 153512}
)

func NewTxTrpc() *TxTrpc {
	return &TxTrpc{
		RequestId: 1,
		Caller:    "com.tencent.qqlivei18n",
		Headers: map[string]string{
			"User-Agent": "okhttp/3.12.13",
		},
		QQLiveHead: &RequestHead{
			RequestId: 0,
			VersionInfo: &VersionInfo{
				AppId: "1200009",
			},
			LocationInfo: &LocationInfo{
				CountryCode: 153505,
				LangCode:    1491963,
			},
		},
	}
}

func (tx *TxTrpc) GenerateRequest(pbRspBody []byte, callee []byte, funcName []byte) ([]byte, error) {
	r := Request{
		RequestId: int32(tx.RequestId),
		Caller:    []byte(tx.Caller),
		Callee:    callee,
		Func:      funcName,
	}
	qqliveHead := tx.QQLiveHead
	qqliveHead.RequestId = int32(tx.RequestId)
	qqliveHead.Callee = string(callee)
	qqliveHead.Func = string(funcName)
	tx.RequestId += 1
	bytesQQLiveHead, err := proto.Marshal(qqliveHead)
	if err != nil {
		log.Error("wetv", "GenerateRequest", err)
		return nil, err
	}
	r.TransInfo = map[string][]byte{
		"qqlive_head": bytesQQLiveHead,
	}

	pbRspHead, err := proto.Marshal(&r)
	if err != nil {
		log.Error("wetv", "GenerateRequest", err)
		return nil, err
	}

	var bytesRspBody bytes.Buffer
	bytesRspBody.Write([]byte("\x09\x30\x00\x00"))
	rspPkgTotalLen := len(pbRspBody) + len(pbRspHead) + 16
	bytesRspBody.Write([]byte{byte(rspPkgTotalLen >> 24), byte(rspPkgTotalLen >> 16), byte(rspPkgTotalLen >> 8), byte(rspPkgTotalLen)})
	pbRspHeadLen := len(pbRspHead)
	bytesRspBody.Write([]byte{byte(pbRspHeadLen >> 8), byte(pbRspHeadLen), 0, 0, 0, 0, 0, 0})
	bytesRspBody.Write(pbRspHead)
	bytesRspBody.Write(pbRspBody)

	return bytesRspBody.Bytes(), nil
}
func (tx *TxTrpc) parseResponse(data []byte) ([]byte, error) {
	if len(data) < 16 {
		return nil, nil
	}
	pbRspHeadLen := int(data[8])<<8 + int(data[9])
	pbRspHead := data[16 : 16+pbRspHeadLen]
	Response := &Response{}
	err := proto.Unmarshal(pbRspHead, Response)
	if err != nil {
		return nil, err
	}
	if Response.ErrorMsg != "" {
		return nil, fmt.Errorf(Response.ErrorMsg)
	}
	pbRspBody := data[16+pbRspHeadLen:]
	return pbRspBody, nil
}
func (tx *TxTrpc) GetVideoList(cid string, page int, dataKey string, pageContext string) error {
	if dataKey == "" {
		dataKey = "cid=" + cid + "&list=1"

	}
	if pageContext == "" {
		pageContext = "page_index=" + strconv.Itoa(page) + "&page_size=30"
	}
	request := &DetailMoreListReq{
		DataKey:     dataKey,
		PageContext: pageContext,
	}
	data, err := proto.Marshal(request)
	if err != nil {
		log.Error("wetv", "GenerateRequest", err)
		return err
	}
	payload, err := tx.GenerateRequest(data, []byte("trpc.video_app_international.trpc_detail_list.VideoDetail"), []byte("/trpc.video_app_international.trpc_detail_list.VideoDetail/GetDetailVideoList"))
	if err != nil {
		log.Error("wetv", "GenerateRequest", err)
		return err
	}
	req := url.NewRequest()
	req.Body = string(payload)
	r, err := client.Do("post", "https://pbacc.wetvinfo.com/trpc.video_app_international.trpc_detail_list", req)
	if err != nil {
		log.Error("wetv", "GenerateRequest", err)
		return err
	}
	RspBody, err := tx.parseResponse(r.Content)
	if err != nil {
		log.Error("wetv", "GenerateRequest", err)
		return err
	}
	DetailVideoListRsp := &DetailVideoListRsp{}
	err = proto.Unmarshal(RspBody, DetailVideoListRsp)
	if err != nil {
		log.Error("wetv", "GenerateRequest", err)
		return err
	}
	nextPageInfo := DetailVideoListRsp.NextPageInfo
	hasNextPage := nextPageInfo.HasNextPage
	dK := nextPageInfo.DataKey
	pC := nextPageInfo.PageContext
	videoList := DetailVideoListRsp.VideoList
	tx.videoList = append(tx.videoList, videoList...)
	if hasNextPage != nil && *hasNextPage {
		return tx.GetVideoList(cid, page+1, *dK, *pC)
	}
	return nil
}
func (tx *TxTrpc) GetDetailPage(vid string, cid string) (d *DetailPageRsp, err error) {
	d = &DetailPageRsp{}
	request := &DetailPageReq{
		Vid: vid,
		Cid: cid,
	}
	data, err := proto.Marshal(request)
	if err != nil {
		log.Error("wetv", "GenerateRequest", err)
		return
	}
	payload, err := tx.GenerateRequest(data, []byte("trpc.video_app_international.trpc_video_detail.VideoDetail"), []byte("/trpc.video_app_international.trpc_video_detail.VideoDetail/GetDetailPage"))
	if err != nil {
		log.Error("wetv", "GenerateRequest", err)
		return
	}
	req := url.NewRequest()
	req.Body = string(payload)
	r, err := client.Do("post", "https://pbacc.wetvinfo.com/trpc.video_app_international.trpc_video_detail", req)
	if err != nil {
		log.Error("wetv", "GenerateRequest", err)
		return
	}
	RspBody, err := tx.parseResponse(r.Content)
	if err != nil {
		log.Error("wetv", "GenerateRequest", err)
		return
	}
	err = proto.Unmarshal(RspBody, d)
	if err != nil {
		log.Error("wetv", "GenerateRequest", err)
		return
	}
	return
}
func AddVideoList(ret *proto2.Data, videoList []*VideoItemData) {
	for _, v := range videoList {
		if v.Vid == nil || v.Title == nil {
			continue
		}
		IsDrm := "false"
		if v.IsDrm != nil {
			IsDrm = strconv.FormatBool(*v.IsDrm)
		}
		if v.EpisodeId == nil {
			v.EpisodeId = proto.Int32(-1)
		}
		ret.VideoList = append(ret.VideoList, &proto2.Video{
			EpisodeId:    *v.Vid,
			EpisodeTitle: *v.Title,
			Episode:      uint32(*v.EpisodeId),
			Extra: map[string]string{
				"isDrm": IsDrm,
			},
		})
	}
}
func GetWeTvMateInfo(ctx context.Context, sharerUrl string) (ret *proto2.Data, code int) {
	var err error
	requestId := ctx.Value("requestId").(string)
	log.Debugf("site: %s, start GetMateInfo", web)
	ret = &proto2.Data{}
	cid := ExtractCid(sharerUrl)
	if cid == "" {
		err = fmt.Errorf("url is invalid")
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		code = 2
		return
	}
	if strings.Contains(cid, "-") {
		cid = strings.Split(cid, "-")[0]
	}
	tx := NewTxTrpc()
	var d *DetailPageRsp
	for _, Ccode := range CountryCode {
		tx.QQLiveHead.LocationInfo.CountryCode = Ccode
		d, err = tx.GetDetailPage("", cid)
		if err != nil {
			if strings.Contains(err.Error(), "版权限制") {
				continue
			}
			log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
			code = 1
			return
		}
		break
	}

	for _, feed := range d.FeedList {
		switch *feed.Type {
		case "FeedDetailsInfo":
			var feedDetailsInfo FeedDetailsInfo
			err = proto.Unmarshal(feed.FeedData, &feedDetailsInfo)
			ret.SeriesTitle = *feedDetailsInfo.Title
			break
		case "FeedDetailsToolbar":
			var feedDetailsToolbar FeedDetailsToolbar
			err = proto.Unmarshal(feed.FeedData, &feedDetailsToolbar)
			break
		case "FeedPlayListHorizontal":
			var feedDetailsRecommend FeedPlayListHorizontal
			err = proto.Unmarshal(feed.FeedData, &feedDetailsRecommend)
			if err != nil {
				log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
				code = 1
				return
			}
			if feedDetailsRecommend.VideoList != nil && len(feedDetailsRecommend.VideoList) > 0 {
				ret.SeriesId = *feedDetailsRecommend.VideoList[0].Cid
				AddVideoList(ret, feedDetailsRecommend.VideoList)
			}
			break
		}
	}
	if len(ret.VideoList) == 0 {
		err = tx.GetVideoList(cid, 0, "", "")
		if err != nil {
			return nil, 2
		}
		AddVideoList(ret, tx.videoList)
	}
	return ret, 0
}
