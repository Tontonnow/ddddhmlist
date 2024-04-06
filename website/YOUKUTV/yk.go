package YOUKUTV

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Tontonnow/ddddhmlist/config"
	"github.com/Tontonnow/ddddhmlist/server"
	"github.com/Tontonnow/ddddhmlist/utils/sesssion"
	"github.com/wangluozhe/requests/models"
	"github.com/wangluozhe/requests/url"
	ut "github.com/wangluozhe/requests/utils"
	"regexp"
	"strconv"
	"strings"
	"time"
	"trpc.group/trpc-go/trpc-go/log"
)

var (
	web       = "youku"
	GetShowId = "https://openapi.youku.com/v2/videos/show.json?client_id=53e6cc67237fc59a&package=com.huawei.hwvplayer.youku&ext=show&video_id="
	GetVideos = "https://openapi.youku.com/v2/shows/videos.json?show_videotype=%E6%AD%A3%E7%89%87&count=100&client_id=53e6cc67237fc59a&page=1&show_id="
	client    = sesssion.NewClient(config.Conf.WebConfig[web])
)

type MtopReq struct {
	Jsv      string `json:"jsv"`
	AppKey   string `json:"appKey"`
	Api      string `json:"api"`
	V        string `json:"v"`
	DataType string `json:"dataType"`
	Type     string `json:"type"`
	Data     string `json:"data"`
	Sign     string `json:"sign"`
}
type MtopRsp struct {
	Api string `json:"api"`
	//Data interface{} `json:"data"`
	Ret []string `json:"ret"`
	V   string   `json:"v"`
}
type FirstPageQ struct {
	ShowId     string `json:"showId"`
	PageSize   int    `json:"pageSize"`
	SystemInfo string `json:"systemInfo"`
	Property   string `json:"property"`
}
type Mtop struct {
	Token string
}
type Show struct {
	Error struct {
		Code        int    `json:"code"`
		Type        string `json:"type"`
		Description string `json:"description"`
	} `json:"error"`
	Id            string `json:"id"`
	Title         string `json:"title"`
	Link          string `json:"link"`
	Thumbnail     string `json:"thumbnail"`
	BigThumbnail  string `json:"bigThumbnail"`
	Duration      string `json:"duration"`
	Category      string `json:"category"`
	State         string `json:"state"`
	Created       string `json:"created"`
	Published     string `json:"published"`
	Description   string `json:"description"`
	Player        string `json:"player"`
	PublicType    string `json:"public_type"`
	CopyrightType string `json:"copyright_type"`
	User          struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"user"`
	Tags           string        `json:"tags"`
	ViewCount      int           `json:"view_count"`
	FavoriteCount  int           `json:"favorite_count"`
	CommentCount   int           `json:"comment_count"`
	UpCount        int           `json:"up_count"`
	DownCount      int           `json:"down_count"`
	OperationLimit []interface{} `json:"operation_limit"`
	DownloadStatus []string      `json:"download_status"`
	Streamtypes    []string      `json:"streamtypes"`
	IsPanorama     int           `json:"is_panorama"`
	Ischannel      int           `json:"ischannel"`
	Show           struct {
		Id          string      `json:"id"`
		Name        string      `json:"name"`
		Link        string      `json:"link"`
		Paid        int         `json:"paid"`
		PayType     interface{} `json:"pay_type"`
		Type        string      `json:"type"`
		Seq         string      `json:"seq"`
		Stage       string      `json:"stage"`
		Collecttime string      `json:"collecttime"`
	} `json:"show"`
	Source struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"source"`
	ReferenceCount int `json:"reference_count"`
}
type FirstPage struct {
	Api  string `json:"api"`
	Data struct {
		Addition struct {
			Data       []interface{} `json:"data"`
			HasNext    string        `json:"hasNext"`
			TotalCount string        `json:"totalCount"`
		} `json:"addition"`
		Artist struct {
			Data []struct {
				Job        string `json:"job"`
				Name       string `json:"name"`
				PersonId   string `json:"personId"`
				Scm        string `json:"scm"`
				ScmInfo    string `json:"scmInfo"`
				ThumbUrl   string `json:"thumbUrl"`
				ThumbUrlLg string `json:"thumbUrlLg"`
				PosterUrl  string `json:"posterUrl,omitempty"`
			} `json:"data"`
			HasNext    string `json:"hasNext"`
			PageNo     string `json:"pageNo"`
			PageSize   string `json:"pageSize"`
			TotalCount string `json:"totalCount"`
			TotalPage  string `json:"totalPage"`
		} `json:"artist"`
		Charge struct {
			AllowCoupon         string `json:"allowCoupon"`
			AllowDiscount       string `json:"allowDiscount"`
			BlurayOpenVipButton struct {
				PicUrl    string `json:"picUrl"`
				PicUrlBig string `json:"picUrlBig"`
				SubTitle  string `json:"subTitle"`
				Title     string `json:"title"`
			} `json:"blurayOpenVipButton"`
			ChargeType   string `json:"chargeType"`
			CurrentPrice string `json:"currentPrice"`
			DisCountText string `json:"disCountText"`
			Duration     string `json:"duration"`
			FreeLook     struct {
				FreeSequences []string `json:"freeSequences"`
				TagText       string   `json:"tagText"`
				TagType       string   `json:"tagType"`
			} `json:"freeLook"`
			FreeSequences              []string `json:"freeSequences"`
			GoldenUpgradeDiamondEnable string   `json:"goldenUpgradeDiamondEnable"`
			Golive                     string   `json:"golive"`
			HasPromoTicket             string   `json:"hasPromoTicket"`
			IsBelongMovies             string   `json:"isBelongMovies"`
			IsBelongTBO                string   `json:"isBelongTBO"`
			IsDiscount                 string   `json:"isDiscount"`
			IsPay                      string   `json:"isPay"`
			IsPurchased                string   `json:"isPurchased"`
			IsVip                      string   `json:"isVip"`
			IsYoukuVip                 string   `json:"isYoukuVip"`
			LoginPlaySetMap            struct {
				Field1 string `json:"0"`
				Field2 string `json:"7"`
				Field3 string `json:"9"`
				Field4 string `json:"10"`
			} `json:"loginPlaySetMap"`
			OnlyVip       string `json:"onlyVip"`
			OpenVipButton struct {
				PicUrl    string `json:"picUrl"`
				PicUrlBig string `json:"picUrlBig"`
				SubTitle  string `json:"subTitle"`
				Title     string `json:"title"`
			} `json:"openVipButton"`
			PackageId      string `json:"packageId"`
			PaidPlaySetMap struct {
				Field1 string `json:"0"`
				Field2 string `json:"7"`
				Field3 string `json:"9"`
				Field4 string `json:"10"`
			} `json:"paidPlaySetMap"`
			PeriodText     string   `json:"periodText"`
			PkgIds         []string `json:"pkgIds"`
			Price          string   `json:"price"`
			PromoTicketNum string   `json:"promoTicketNum"`
			TokenValid     string   `json:"tokenValid"`
			TvPayInfoResp  struct {
				Attributes      string `json:"attributes"`
				BuyDesc         string `json:"buyDesc"`
				IsPay           string `json:"isPay"`
				IsPurchased     string `json:"isPurchased"`
				PlayerBarDesc   string `json:"playerBarDesc"`
				PlayerBarList   string `json:"playerBarList"`
				PlayerRightList string `json:"playerRightList"`
			} `json:"tvPayInfoResp"`
			UsedCouponNo        string `json:"usedCouponNo"`
			UsedPromoTicketText string `json:"usedPromoTicketText"`
			VipDays             string `json:"vipDays"`
		} `json:"charge"`
		Paras struct {
			CommentsOpen string   `json:"commentsOpen"`
			DanmuOpen    string   `json:"danmuOpen"`
			ModuleSort   []string `json:"moduleSort"`
			Qrcode       struct {
				Link string `json:"link"`
				Show string `json:"show"`
				Tips string `json:"tips"`
			} `json:"qrcode"`
			SnapshotOpen string `json:"snapshotOpen"`
		} `json:"paras"`
		Recommend struct {
			Data       []interface{} `json:"data"`
			HasNext    string        `json:"hasNext"`
			TotalCount string        `json:"totalCount"`
		} `json:"recommend"`
		Show struct {
			Area                []string      `json:"area"`
			Audiolang           string        `json:"audiolang"`
			Belong              string        `json:"belong"`
			CopyrightPic        string        `json:"copyrightPic"`
			CopyrightStatusMap  string        `json:"copyrightStatusMap"`
			CopyrightText       string        `json:"copyrightText"`
			Director            []string      `json:"director"`
			EpisodeLast         string        `json:"episodeLast"`
			EpisodeTotal        string        `json:"episodeTotal"`
			Exclusive           string        `json:"exclusive"`
			ExtShowId           string        `json:"extShowId"`
			ExtType             string        `json:"extType"`
			From                string        `json:"from"`
			Genre               []string      `json:"genre"`
			HasStar             string        `json:"hasStar"`
			Hasvideotype        string        `json:"hasvideotype"`
			Heat                string        `json:"heat"`
			IsDynTotal          string        `json:"isDynTotal"`
			KidsAgeMax          string        `json:"kidsAgeMax"`
			KidsAgeMin          string        `json:"kidsAgeMin"`
			LastSequence        string        `json:"lastSequence"`
			LicensePic          string        `json:"licensePic"`
			Mark                string        `json:"mark"`
			Name                string        `json:"name"`
			OriVThumbUrl        string        `json:"oriVThumbUrl"`
			Paid                string        `json:"paid"`
			PayType             []string      `json:"payType"`
			PayTypeStr          string        `json:"payTypeStr"`
			Performer           []string      `json:"performer"`
			PlaySet             string        `json:"playSet"`
			Prevue              string        `json:"prevue"`
			ProgramId           string        `json:"programId"`
			ReleaseDate         string        `json:"releaseDate"`
			Seconds             string        `json:"seconds"`
			SeriesId            string        `json:"seriesId"`
			ShowCategory        string        `json:"showCategory"`
			ShowCategoryName    string        `json:"showCategoryName"`
			ShowDesc            string        `json:"showDesc"`
			ShowId              string        `json:"showId"`
			ShowLength          string        `json:"showLength"`
			ShowLongId          string        `json:"showLongId"`
			ShowName            string        `json:"showName"`
			ShowStrId           string        `json:"showStrId"`
			ShowSubtitle        string        `json:"showSubtitle"`
			ShowThumbUrl        string        `json:"showThumbUrl"`
			ShowTotalVv         string        `json:"showTotalVv"`
			ShowType            string        `json:"showType"`
			ShowVersions        []interface{} `json:"showVersions"`
			ShowVthumbUrl       string        `json:"showVthumbUrl"`
			Starring            []string      `json:"starring"`
			Subscribe           string        `json:"subscribe"`
			ThumbRgb            string        `json:"thumbRgb"`
			Tips                string        `json:"tips"`
			VideosMultiLanguage string        `json:"videosMultiLanguage"`
			ViewPoint           string        `json:"viewPoint"`
			ViewTag             string        `json:"viewTag"`
			VmacState           string        `json:"vmacState"`
			VodFullPrice        string        `json:"vodFullPrice"`
			VodTicket           string        `json:"vodTicket"`
			YoukuReleaseDate    string        `json:"youkuReleaseDate"`
			ZrealDownloadStatus string        `json:"zrealDownloadStatus"`
		} `json:"show"`
		Video struct {
			ZP struct {
				Data []struct {
					Bcp           string            `json:"bcp"`
					ExtType       string            `json:"extType"`
					ExtVideoStrId string            `json:"extVideoStrId"`
					FileSize      map[string]string `json:"fileSize"`
					GmtModified   string            `json:"gmtModified"`
					LibraryTag    []string          `json:"libraryTag"`
					NeteaseInfo   struct {
					} `json:"neteaseInfo"`
					Paid     string `json:"paid"`
					PlayInfo struct {
						ExtVideoStrId string `json:"extVideoStrId"`
						ExtShowId     string `json:"extShowId"`
						WebUrl        string `json:"webUrl"`
						ProgramId     string `json:"programId"`
					} `json:"playInfo"`
					PlaySet   string `json:"playSet"`
					ProgramId string `json:"programId"`
					PureDrm   string `json:"pureDrm"`
					RcTitle   string `json:"rcTitle"`
					RealMap   map[string]struct {
						With   string `json:"with"`
						Size   string `json:"size"`
						Height string `json:"height"`
					} `json:"realMap"`
					Seconds            string   `json:"seconds"`
					SeeTa              string   `json:"seeTa,omitempty"`
					Sequence           string   `json:"sequence"`
					ShowVideoInvisible string   `json:"showVideoInvisible"`
					ShowVideoSeq       string   `json:"showVideoSeq"`
					ShowVideoStage     string   `json:"showVideoStage"`
					Strategy           string   `json:"strategy"`
					Stream             []string `json:"stream"`
					SubStage           string   `json:"subStage"`
					Tail               string   `json:"tail,omitempty"`
					ThumbUrl           string   `json:"thumbUrl"`
					Title              string   `json:"title"`
					VideoType          string   `json:"videoType"`
					YoukuShowLongId    string   `json:"youkuShowLongId"`
					YoukuVideoLongId   string   `json:"youkuVideoLongId"`
					Head               string   `json:"head,omitempty"`
					Mark               string   `json:"mark,omitempty"`
					Mon                string   `json:"mon,omitempty"`
					MonVod             string   `json:"monVod,omitempty"`
					PayBag             string   `json:"payBag,omitempty"`
					Svip               string   `json:"svip,omitempty"`
					SvipAhead          string   `json:"svipAhead,omitempty"`
					SvipFree           string   `json:"svipFree,omitempty"`
					Tips               string   `json:"tips,omitempty"`
					Vod                string   `json:"vod,omitempty"`
				} `json:"data"`
				HasNext    string `json:"hasNext"`
				PageNo     string `json:"pageNo"`
				PageSize   string `json:"pageSize"`
				TotalCount string `json:"totalCount"`
				TotalPage  string `json:"totalPage"`
			} `json:"正片"`
			ZB struct {
				Data []struct {
					Bcp           string `json:"bcp"`
					ExtType       string `json:"extType"`
					ExtVideoStrId string `json:"extVideoStrId"`
					GmtModified   string `json:"gmtModified"`
					NeteaseInfo   struct {
					} `json:"neteaseInfo"`
					Paid     string `json:"paid"`
					PlayInfo struct {
						ExtVideoStrId string `json:"extVideoStrId"`
						ExtShowId     string `json:"extShowId"`
						WebUrl        string `json:"webUrl"`
						ProgramId     string `json:"programId"`
					} `json:"playInfo"`
					PlaySet            string        `json:"playSet"`
					ProgramId          string        `json:"programId"`
					PureDrm            string        `json:"pureDrm"`
					RcTitle            string        `json:"rcTitle"`
					RemainType         string        `json:"remainType"`
					Seconds            string        `json:"seconds"`
					Sequence           string        `json:"sequence"`
					ShowVideoInvisible string        `json:"showVideoInvisible"`
					ShowVideoSeq       string        `json:"showVideoSeq"`
					ShowVideoStage     string        `json:"showVideoStage"`
					Strategy           string        `json:"strategy"`
					Stream             []interface{} `json:"stream"`
					SubStage           string        `json:"subStage"`
					ThumbUrl           string        `json:"thumbUrl"`
					Title              string        `json:"title"`
					VTypeMark          string        `json:"vTypeMark"`
					VideoType          string        `json:"videoType"`
					YoukuShowLongId    string        `json:"youkuShowLongId"`
					YoukuVideoLongId   string        `json:"youkuVideoLongId"`
				} `json:"data"`
				HasNext    string `json:"hasNext"`
				PageNo     string `json:"pageNo"`
				PageSize   string `json:"pageSize"`
				TotalCount string `json:"totalCount"`
				TotalPage  string `json:"totalPage"`
			} `json:"周边"`
		} `json:"video"`
	} `json:"data"`
	Ret []string `json:"ret"`
	V   string   `json:"v"`
}
type Videos struct {
	Total  any     `json:"total"`
	Videos []Video `json:"videos"`
}
type Video struct {
	Id             string      `json:"id"`
	Stage          string      `json:"stage"`
	Seq            string      `json:"seq"`
	Title          string      `json:"title"`
	Link           string      `json:"link"`
	Streamtypes    []string    `json:"streamtypes"`
	Thumbnail      string      `json:"thumbnail"`
	ThumbnailV2    string      `json:"thumbnail_v2"`
	Bigthumbnail   string      `json:"bigthumbnail"`
	Duration       string      `json:"duration"`
	Category       string      `json:"category"`
	ViewCount      int         `json:"view_count"`
	FavoriteCount  int         `json:"favorite_count"`
	UpCount        int         `json:"up_count"`
	DownCount      int         `json:"down_count"`
	OperationLimit []string    `json:"operation_limit"`
	DownloadStatus []string    `json:"download_status"`
	Published      string      `json:"published"`
	State          string      `json:"state"`
	RcTitle        interface{} `json:"rc_title"`
	Audiolang      []struct {
		Vid                  string `json:"vid"`
		Langcode             string `json:"langcode"`
		OttCibnLicenseState  string `json:"ott_cibn_license_state"`
		ShowVideoInvisible   bool   `json:"show_video_invisible"`
		StrategyFlagsBitwise int    `json:"strategy_flags_bitwise"`
		Lang                 string `json:"lang"`
		OttWasuLicenseState  string `json:"ott_wasu_license_state"`
	} `json:"audiolang"`
	IsPanorama int `json:"is_panorama"`
	Paid       int `json:"paid"`
}

func NewMtop() (t *Mtop, err error) {
	t = &Mtop{
		//Session: requests.NewSession(),
	}
	err = t.ReFreshToken()
	return

}
func NewMtopReq(t, api, d, s string) map[string]string {
	return map[string]string{
		"jsv":      "0",
		"v":        "1.0",
		"dataType": "json",
		"type":     "json",
		"appKey":   "24679788",
		"api":      api,
		"data":     d,
		"sign":     s,
		"t":        t,
	}
}
func (m *Mtop) ReFreshToken() (err error) {
	rsp, err := client.Do("get", "https://acs.youku.com/h5/mtop.wenyu.video.show.detail.firstpage/1.0/?jsv=0&appKey=24679788", nil)
	if err != nil {
		return
	}
	ck := rsp.Cookies
	for _, v := range ck {
		if v.Name == "_m_h5_tk" {
			m.Token = strings.Split(v.Value, "_")[0]
		}
	}
	return
}
func (m *Mtop) MtopH5Request(api, data string) (r []byte, err error) {
	req := url.NewRequest()
	var rsp *models.Response
	var mp MtopRsp
	for {
		t := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
		p := "jsv=0&appKey=24679788&api=%s&v=1.0&dataType=json&type=json&data=%s&sign=%s&t=%s"
		s := m.Token + "&" + t + "&" + "24679788" + "&" + data
		sign := ut.MD5(s)
		pa := fmt.Sprintf(p, api, data, sign, t)
		rsp, err = client.Do("get", "https://acs.youku.com/h5/mtop.wenyu.video.show.detail.firstpage/1.0/?"+pa, req)
		if err != nil {
			return
		}

		if rsp.StatusCode != 200 {
			err = fmt.Errorf("status code %d", rsp.StatusCode)
			break
		}
		err = json.Unmarshal(rsp.Content, &mp)
		if err != nil {
			return
		}
		Ret := mp.Ret[0]
		if !strings.Contains(Ret, "SUCCESS") {
			if !strings.Contains(Ret, "TOKEN") {
				err = fmt.Errorf(Ret)
				break
			}
			continue
		}
		r = rsp.Content
		break
	}
	return
}
func (m *Mtop) GetList(showId string) (r []*server.Video, err error) {
	d := "{\"showId\":\"%s\",\"pageSize\":600,\"systemInfo\":\"{}\",\"property\":\"{}\"}"
	sd := fmt.Sprintf(d, showId)
	rsp, err := m.MtopH5Request("mtop.wenyu.video.show.detail.firstpage", sd)
	if err != nil {
		return
	}
	var f *FirstPage
	err = json.Unmarshal(rsp, &f)
	if err != nil {
		return
	}
	for _, v := range f.Data.Video.ZP.Data {
		e, _ := strconv.Atoi(v.Sequence)
		r = append(r, &server.Video{
			EpisodeTitle: v.Title,
			EpisodeId:    v.ExtVideoStrId,
			Episode:      uint32(e),
			Extra:        v.FileSize,
		})

	}
	return
}
func (m *Mtop) GetVid(vid string) string {
	re := regexp.MustCompile(`id_(.*?)\.`)
	vids := re.FindStringSubmatch(vid)
	if len(vids) < 2 {
		return ""
	}
	return vids[1]
}
func (m *Mtop) GetShowId(vid string) (show Show, err error) {
	rsp, err := client.Do("get", GetShowId+vid, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(rsp.Content, &show)
	if err != nil {
		return
	}
	if show.Error.Code != 0 {
		err = fmt.Errorf(show.Error.Description)
	}
	return
}
func (m *Mtop) GetVideos(showId string) (r []*server.Video, err error) {
	r = []*server.Video{}
	page := 1
	var vs []Video
	var rsp *models.Response
	for {
		end := false
		rsp, err = client.Do("get", GetVideos+showId+"&page="+strconv.Itoa(page), nil)
		if err != nil {
			return
		}
		var v Videos
		err = json.Unmarshal(rsp.Content, &v)
		if err != nil {
			return
		}
		vs = append(vs, v.Videos...)
		switch v.Total.(type) {
		case int:
			if v.Total.(int) <= len(vs) {
				end = true
			}
		case string:
			Total, _ := strconv.Atoi(v.Total.(string))
			if Total <= len(vs) {
				end = true
			}
		default:
			end = true
		}
		if end {
			break
		}
		page++
	}
	for _, v := range vs {
		e, _ := strconv.Atoi(v.Seq)
		r = append(r, &server.Video{
			EpisodeTitle: v.Title,
			EpisodeId:    v.Id,
			Episode:      uint32(e),
			Extra: map[string]string{
				"stream_types": strings.Join(v.Streamtypes, ","),
			},
		})
	}

	return
}
func GetMateInfo(ctx context.Context, sharerUrl string) (r *server.Data, code int) {
	var err error
	requestId := ctx.Value("requestId").(string)
	log.Debugf("site: %s, start GetMateInfo", web)
	r = &server.Data{}
	m, err := NewMtop()
	if err != nil {
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		code = 1
		return
	}
	vid := m.GetVid(sharerUrl)
	if vid == "" {
		err = fmt.Errorf("vid is empty")
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		code = 2
		return
	}
	s, err := m.GetShowId(vid)
	if err != nil {
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		code = 1
		return
	}
	r.SeriesTitle = s.Show.Name
	r.SeriesId = s.Show.Id
	r.VideoList, err = m.GetList(r.SeriesId)
	if err != nil { //国外服务器建议直接使用下面的方法
		r.VideoList, err = m.GetVideos(r.SeriesId)
		if err != nil {
			log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
			code = 1
		}
	}
	return
}
