package IQ

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Tontonnow/ddddhmlist/server"
	"github.com/bitly/go-simplejson"
	"github.com/wangluozhe/requests/url"
	"sort"
	"strconv"
	"strings"
	"trpc.group/trpc-go/trpc-go/log"
)

var (
	decodeID       = "https://pcw-api.iq.com/api/decode/%s?platformId=3&modeCode=intl&langCode=intl"
	GetEpgInfo     = "https://itv.ptqy.gitv.tv/api/epgInfo/%s"
	timelist       = "2000,2001,2002,2003,2004,2005,2006,2007,2008,2009,2010,2011,2012,2013,2014,2015,2016,2017,2018,2019,2020,2021,2022,2023,2024,2025,2026"
	GetVarietyInfo = "https://pcw-api.iqiyi.com/album/source/svlistinfo?cid=6&sourceid=%s&timelist=%s"
	GetAvlistinfo  = "https://pcw-api.iqiyi.com/albums/album/avlistinfo?aid=%s&page=%s&size=200"
)

type EpgInfo struct {
	ChnId            int         `json:"chnId"`
	ChnName          string      `json:"chnName"`
	QipuId           int64       `json:"qipuId"`
	AlbumId          int         `json:"albumId"`
	AlbumName        string      `json:"albumName"`
	Name             string      `json:"name"`
	Focus            string      `json:"focus"`
	AlbumPic         string      `json:"albumPic"`
	PosterPic        string      `json:"posterPic"`
	Score            string      `json:"score"`
	Drm              string      `json:"drm"`
	Hdr              string      `json:"hdr"`
	ShortName        string      `json:"shortName"`
	IsExclusive      int         `json:"isExclusive"`
	Is3D             int         `json:"is3D"`
	SourceCode       int         `json:"sourceCode"`
	Order            int         `json:"order"`
	SuperId          int         `json:"superId"`
	InitIssueTime    string      `json:"initIssueTime"`
	Len              int         `json:"len"`
	VipType          string      `json:"vipType"`
	PublishTime      string      `json:"publishTime"`
	IsSeries         int         `json:"isSeries"`
	ContentType      int         `json:"contentType"`
	ContentTypeV2    int         `json:"contentTypeV2"`
	Type4K           string      `json:"type4k"`
	Dolby            string      `json:"dolby"`
	BusinessTypes    string      `json:"businessTypes"`
	PositiveId       int         `json:"positiveId"`
	InteractType     int         `json:"interactType"`
	Dance            string      `json:"dance"`
	VipCt            string      `json:"vipCt"`
	UpUid            string      `json:"upUid"`
	AlbumPic2        string      `json:"albumPic2"`
	Cormrk           string      `json:"cormrk"`
	CanSub           int         `json:"canSub"`
	PosiPay          int         `json:"posiPay"`
	Ctt              string      `json:"ctt"`
	IeType           string      `json:"ieType"`
	AlbumChnId       int         `json:"albumChnId"`
	PHeat            int         `json:"pHeat"`
	PAlbum           interface{} `json:"pAlbum"`
	ParentPosterPic  string      `json:"parentPosterPic"`
	EtV2             int         `json:"etV2"`
	SuTime           string      `json:"suTime"`
	ShortNameV2      string      `json:"shortNameV2"`
	Copyrmrk         string      `json:"copyrmrk"`
	LimitedFree      int         `json:"limitedFree"`
	FreeEndTime      int         `json:"freeEndTime"`
	SubTitle         string      `json:"subTitle"`
	VQ               string      `json:"vQ"`
	AQ               string      `json:"aQ"`
	PCount           int         `json:"pCount"`
	Desc             string      `json:"desc"`
	Tag              string      `json:"tag"`
	TagV2            string      `json:"tagV2"`
	Color            string      `json:"color"`
	ITime            string      `json:"iTime"`
	FstFrmCov        string      `json:"fstFrmCov"`
	HRecType         string      `json:"hRecType"`
	HRecSentence     string      `json:"hRecSentence"`
	Hot              int         `json:"hot"`
	HotSwitch        int         `json:"hotSwitch"`
	ControlStatus    int         `json:"controlStatus"`
	Rating           string      `json:"rating"`
	MultiEpisodeInfo struct {
		MultiEpisodeCount int  `json:"multiEpisodeCount"`
		MultiEpisodeOrder int  `json:"multiEpisodeOrder"`
		MultiEpisode      bool `json:"multiEpisode"`
	} `json:"multiEpisodeInfo"`
	Pic                string `json:"pic"`
	CImgSize           string `json:"cImgSize"`
	Is1080             int    `json:"is1080"`
	Season             int    `json:"season"`
	InfoControlStatus  int    `json:"infoControlStatus"`
	FatherEpisodeId    int    `json:"fatherEpisodeId"`
	FatherEpisodeIdStr string `json:"fatherEpisodeIdStr"`
	IsQiyiProduced     int    `json:"isQiyiProduced"`
	AlbumIdStr         string `json:"albumIdStr"`
	ContentRating      struct {
		Display bool   `json:"display"`
		Rating  string `json:"rating"`
		Warning string `json:"warning"`
	} `json:"contentRating"`
	IsFollowFeatured    bool          `json:"isFollowFeatured"`
	VolunteerTranslates []interface{} `json:"volunteerTranslates"`
	PlayLocSuffix       string        `json:"playLocSuffix"`
	AllowRegion         []string      `json:"allowRegion"`
	Screenshot          struct {
		ImgUrl     string `json:"imgUrl"`
		WebImgUrl  string `json:"webImgUrl"`
		Interval   int    `json:"interval"`
		ImageSize  string `json:"imageSize"`
		MergeCount string `json:"mergeCount"`
	} `json:"screenshot"`
	CWebpImgSize      string   `json:"cWebpImgSize"`
	IsDolby           int      `json:"isDolby"`
	ExtraName         string   `json:"extraName"`
	QipuIdStr         string   `json:"qipuIdStr"`
	AlternativeTitles []string `json:"alternativeTitles"`
	AlbumWebpPic      string   `json:"albumWebpPic"`
	MatchedLang       int      `json:"matchedLang"`
	DefaultVid        string   `json:"defaultVid"`
	AlbumDesc         string   `json:"albumDesc"`
}
type AvList struct {
	AlbumId           string        `json:"albumId"`
	EpisodeList       []AvEpisode   `json:"epsodelist"`
	BeforeEpisodeList []interface{} `json:"beforeEpisodeList"`
	AfterEpisodeList  []interface{} `json:"afterEpisodeList"`
	PreEpisodeList    []interface{} `json:"preEpisodeList"`
	StarEpisodeList   []interface{} `json:"starEpisodeList"`
	Updateprevuelist  []interface{} `json:"updateprevuelist"`
	Vipprevuelist     []interface{} `json:"vipprevuelist"`
	PrePrevueList     []interface{} `json:"prePrevueList"`
	StarPrevueList    []interface{} `json:"starPrevueList"`
	Size              int           `json:"size"`
	Page              int           `json:"page"`
	Total             int           `json:"total"`
	Part              int           `json:"part"`
	LatestOrder       int           `json:"latestOrder"`
	VideoCount        int           `json:"videoCount"`
	HasMore           bool          `json:"hasMore"`
}
type AvEpisode struct {
	TvId         int    `json:"tvId"`
	Description  string `json:"description"`
	Subtitle     string `json:"subtitle"`
	Vid          string `json:"vid"`
	Name         string `json:"name"`
	PlayUrl      string `json:"playUrl"`
	IssueTime    int64  `json:"issueTime"`
	PublishTime  int64  `json:"publishTime"`
	ContentType  int    `json:"contentType"`
	PayMark      int    `json:"payMark"`
	PayMarkUrl   string `json:"payMarkUrl"`
	ImageUrl     string `json:"imageUrl"`
	Duration     string `json:"duration"`
	Period       string `json:"period"`
	Exclusive    bool   `json:"exclusive"`
	Order        int    `json:"order"`
	Effective    bool   `json:"effective"`
	QiyiProduced bool   `json:"qiyiProduced"`
	Focus        string `json:"focus"`
	ShortTitle   string `json:"shortTitle"`
	People       struct {
	} `json:"people"`
	InteractionType      int      `json:"interactionType"`
	IsEnabledInteraction int      `json:"isEnabledInteraction"`
	ImageSize            []string `json:"imageSize"`
	ImageProductionType  []string `json:"imageProductionType"`
	OrderName            string   `json:"orderName"`
}
type Rsp struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}
type Iqy struct {
	ModeCode  string          `json:"modeCode,omitempty"`
	LangCode  string          `json:"langCode,omitempty"`
	Proxy     string          `json:"proxy,omitempty"`
	VideoList []*server.Video `json:"videoList,omitempty"`
	code      int
}

func CreateIqy(req ...url.Request) *Iqy {
	i := &Iqy{}
	//i.Session = requests.NewSession()
	return i
}
func (i *Iqy) doRequest(u string, ks ...string) (data *simplejson.Json, err error) {
	r, err := client.Do("get", u, nil)
	if err != nil {
		return
	}
	if r.StatusCode != 200 {
		err = r.RaiseForStatus()
		if err != nil {
			return
		}
	}
	rsp, err := r.SimpleJson()
	if err != nil {
		return
	}
	codep := rsp.Get("code")
	var code string
	switch codep.Interface().(type) {
	case string:
		code, err = codep.String()
		if err != nil {
			return
		}
	case json.Number:
		cd, er := codep.Int()
		if er != nil {
			return
		}
		code = strconv.Itoa(cd)
	default:
		code = "0"
	}

	if code != "0" && code != "A00000" {
		msg, _ := rsp.Get("msg").String()
		err = fmt.Errorf("code:%s,msg:%s", code, msg)
		return
	}
	if len(ks) > 0 {
		key := ks[0]
		if key == "" {
			data = rsp
		} else {
			data = rsp.Get(key)
		}
	} else {
		data = rsp.Get("data")
	}
	return

}
func (i *Iqy) DecodeID(id string) (Rid string, err error) {
	u := fmt.Sprintf(decodeID, id)
	data, err := i.doRequest(u)
	if err != nil {
		return
	}
	t, err := data.Int()
	if err != nil {
		return
	}
	Rid = strconv.Itoa(t)
	return
}
func (i *Iqy) GetEpgInfo(u string) (r EpgInfo, err error) {
	data, err := i.doRequest(u)
	if err != nil {
		return
	}
	b, err := data.MarshalJSON()
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &r)
	return
}
func (i *Iqy) GetListinfo(aid string) (r []*server.Video, err error) {
	u := fmt.Sprintf(GetAvlistinfo, aid, "1")
	data, err := i.doRequest(u)
	if err != nil {
		return
	}
	b, err := data.MarshalJSON()
	if err != nil {
		return

	}
	vList := AvList{}
	err = json.Unmarshal(b, &vList)
	if err != nil {
		return
	}
	var v []AvEpisode
	v = append(v, vList.EpisodeList...)
	if vList.Total > 200 {
		var ch = make(chan []AvEpisode, vList.Total/200+2)
		for n := 2; n < vList.Total/200+2; n++ {
			go func(n int) {
				var (
					da *simplejson.Json
					by []byte
				)

				u := fmt.Sprintf(GetAvlistinfo, aid, strconv.Itoa(n))
				da, err = i.doRequest(u)
				if err != nil {
					ch <- nil
				}
				by, err = da.MarshalJSON()
				if err != nil {
					ch <- nil
				}
				vList := AvList{}
				err = json.Unmarshal(by, &vList)
				if err != nil {
					ch <- nil
				}
				ch <- vList.EpisodeList
			}(n)
		}
	}
	for _, e := range v {
		r = append(r, &server.Video{
			EpisodeId:    strconv.Itoa(e.TvId),
			EpisodeTitle: e.Name,
			IsVip:        e.PayMark != 0,
			Episode:      uint32(e.Order),
			Extra: map[string]string{
				"subtitle": e.Subtitle,
				"vid":      e.Vid,
			},
		})

	}
	return
}
func (i *Iqy) GetVarietyInfo(id string) (r []*server.Video, err error) {
	u := fmt.Sprintf(GetVarietyInfo, id, timelist)
	data, err := i.doRequest(u)
	if err != nil {
		return
	}
	var vList map[string][]AvEpisode
	b, err := data.MarshalJSON()
	if err != nil {
		return
	}
	err = json.Unmarshal(b, &vList)
	if err != nil {
		return

	}
	for _, e := range vList {
		for _, v := range e {
			r = append(r, &server.Video{
				EpisodeId:    strconv.Itoa(int(v.TvId)),
				EpisodeTitle: v.Name,
				IsVip:        v.PayMark != 0,
				Episode:      uint32(v.Order),
				Extra: map[string]string{
					"subtitle": v.Subtitle,
					"vid":      v.Vid,
				},
			})

		}

	}
	return
}
func (i *Iqy) GetBaseInfo(r *server.Data, sharerUrl string, iq bool) (epgInfo EpgInfo, err error) {
	vid, err, c := GetVid(sharerUrl)
	if err != nil {
		return
	}
	if !c {
		vid, err = i.DecodeID(vid)
		if err != nil {
			return
		}
	}
	var u string
	if iq {
		u = fmt.Sprintf(GetIqEpgInfo, vid)
	} else {
		u = fmt.Sprintf(GetEpgInfo, vid)
	}
	epgInfo, err = i.GetEpgInfo(u)
	if err != nil {
		return
	}
	r.IsSeries = epgInfo.IsSeries == 1
	r.Extra = map[string]string{
		"type4k":  epgInfo.Type4K,
		"dolby":   epgInfo.Dolby,
		"hdr":     epgInfo.Hdr,
		"isDolby": strconv.Itoa(epgInfo.IsDolby),
	}
	if iq {
		r.Extra["AllowRegion"] = strings.Join(epgInfo.AllowRegion, ",")
	}
	return
}
func (i *Iqy) GetMateInfo(sharerUrl string) (r *server.Data, err error) {
	r = &server.Data{}
	epgInfo, err := i.GetBaseInfo(r, sharerUrl, false)
	if err != nil {
		return
	}
	if epgInfo.ChnId == 1 {
		r.SeriesTitle = epgInfo.Name
		r.SeriesId = strconv.FormatInt(epgInfo.QipuId, 10)
		return
	} else {
		r.SeriesTitle = epgInfo.AlbumName
		r.SeriesId = strconv.Itoa(epgInfo.AlbumId)
		if r.SeriesTitle == "" {
			r.SeriesTitle = epgInfo.Name
			r.SeriesId = strconv.Itoa(epgInfo.SuperId)
		}
		// 电视剧、纪录片、动漫、知识、儿童
		if epgInfo.ChnId == 2 || epgInfo.ChnId == 3 || epgInfo.ChnId == 4 || epgInfo.ChnId == 12 || epgInfo.ChnId == 15 {
			r.VideoList, err = i.GetListinfo(r.SeriesId)
			if err != nil {
				return
			}
		} else if epgInfo.ChnId == 6 { // 综艺
			r.VideoList, err = i.GetVarietyInfo(r.SeriesId)
			if err != nil {
				return
			}
		}

		sort.Sort(r)
	}
	return
}
func GetVid(sharerUrl string) (vid string, err error, c bool) {
	ud, err := url.Parse(sharerUrl)
	if err != nil {
		return
	}

	if vid = ud.Params.Get("shareId"); vid != "" {
		vid = strings.ReplaceAll(vid, "%3D", "=")
		if decodedId, err := base64.StdEncoding.DecodeString(vid); err == nil {
			return string(decodedId), nil, true
		}
		return "", err, false
	}

	vid = strings.Split(ud.Path, "/")[len(strings.Split(ud.Path, "/"))-1]
	vid, _, _ = strings.Cut(vid, ".")
	vid = strings.Split(vid, "-")[len(strings.Split(vid, "-"))-1]
	vid = strings.Split(vid, "_")[len(strings.Split(vid, "_"))-1]
	return vid, nil, false
}
func GetMateInfo(ctx context.Context, sharerUrl string) (r *server.Data, code int) {
	var (
		err error
	)
	r = &server.Data{}
	requestId := ctx.Value("requestId").(string)
	log.Debugf("site: %s, start GetMateInfo", web)
	i := NewIq()
	r, err = i.GetMateInfo(sharerUrl)
	if err != nil {
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
	}
	code = i.code
	return
}
