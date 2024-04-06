package IQ

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Tontonnow/ddddhmlist/config"
	"github.com/Tontonnow/ddddhmlist/server"
	j "github.com/Tontonnow/ddddhmlist/utils/jwt"
	"github.com/Tontonnow/ddddhmlist/utils/sesssion"
	"strconv"
	"time"
	"trpc.group/trpc-go/trpc-go/log"
)

var (
	web          = "iq"
	DeviceId     = "1008611"
	registerMode = "https://tv-api2.iq.com/api/registerMode?sdkintVer=28&apkVer=8.2.0&model=TV&deviceId=%s&uuid=20201127144624572xXkbWBuKY100031&macAddr=0&localLang=zh_cn"
	GetEpisode   = "https://tv-api2.iq.com/api/v2/episodeListWithPreview/%s?langCode=%s&num=60&gps=1&ua=TV&deviceId=&platform=&isVip=true&sid=&network=1&uid=&pos=%d&modeCode=intl&pspStatus=-1"
	header       = map[string]string{
		"User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36",
	}
	GetIqEpgInfo = "https://pcw-api.iq.com/api/epgInfo/%s?platformId=4&langCode=zh_cn&modeCode=intl&deviceId=ew&uid=&pspStatus="
	client       = sesssion.NewClient(config.Conf.WebConfig[web])
)

type EpgInfoIQ struct {
	Total      int    `json:"total"`
	Pos        int    `json:"pos"`
	HasMore    bool   `json:"hasMore"`
	ChnId      int    `json:"chnId"`
	EpgS       []Epg  `json:"epg"`
	Code       string `json:"code"`
	SourceCode int    `json:"sourceCode"`
	AlbumId    int    `json:"albumId"`
	AlbumName  string `json:"albumName"`
}
type Epg struct {
	QipuId             int64       `json:"qipuId"`
	QipuIdStr          string      `json:"qipuIdStr"`
	DefMultiImage      interface{} `json:"defMultiImage"`
	BackGroundPicColor interface{} `json:"backGroundPicColor"`
	ChnId              int         `json:"chnId"`
	Name               string      `json:"name"`
	ShortName          string      `json:"shortName"`
	AlbumPic           string      `json:"albumPic"`
	PosterPic          interface{} `json:"posterPic"`
	AlbumWebpPic       string      `json:"albumWebpPic"`
	Focus              string      `json:"focus"`
	Score              string      `json:"score"`
	Rating             string      `json:"rating"`
	VipInfo            struct {
	} `json:"vipInfo"`
	VipType             string        `json:"vipType"`
	IsExclusive         int           `json:"isExclusive"`
	Is3D                int           `json:"is3D"`
	Is1080              int           `json:"is1080"`
	IsDolby             int           `json:"isDolby"`
	PublishTime         string        `json:"publishTime"`
	InitIssueTime       string        `json:"initIssueTime"`
	Desc                string        `json:"desc"`
	Drm                 string        `json:"drm"`
	Hdr                 string        `json:"hdr"`
	SubTitle            string        `json:"subTitle"`
	ContentType         int           `json:"contentType"`
	IsSeries            int           `json:"isSeries"`
	Order               int           `json:"order"`
	AllowRegion         interface{}   `json:"allowRegion"`
	Categories          []int64       `json:"categories"`
	ShareAllowed        interface{}   `json:"shareAllowed"`
	PlayLocSuffix       string        `json:"playLocSuffix"`
	AlbumLocSuffix      interface{}   `json:"albumLocSuffix"`
	PlayHrefLangPile    interface{}   `json:"playHrefLangPile"`
	AlbumHrefLangPile   interface{}   `json:"albumHrefLangPile"`
	IsQiyiProduced      int           `json:"isQiyiProduced"`
	MatchedLang         int           `json:"matchedLang"`
	AlternativeTitles   []interface{} `json:"alternativeTitles"`
	AlbumPicColor       interface{}   `json:"albumPicColor"`
	IsFollowFeatured    bool          `json:"isFollowFeatured"`
	InfoControlStatus   interface{}   `json:"infoControlStatus"`
	VolunteerTranslates []interface{} `json:"volunteerTranslates"`
	FatherCollectionIds []interface{} `json:"fatherCollectionIds"`
	Season              int           `json:"season"`
	FocusImage          struct {
		FocusImagesWithLang []interface{} `json:"focusImagesWithLang"`
	} `json:"focusImage"`
	FirstPlayTimeLine     interface{} `json:"firstPlayTimeLine"`
	FirstPlayTimeOnlyDate interface{} `json:"firstPlayTimeOnlyDate"`
	PeopleInfosMap        struct {
	} `json:"peopleInfosMap"`
	AlbumId    int         `json:"albumId"`
	AlbumName  string      `json:"albumName"`
	Len        int         `json:"len"`
	Type4K     string      `json:"type4k"`
	Dolby      string      `json:"dolby"`
	SourceCode int         `json:"sourceCode"`
	Pic        string      `json:"pic"`
	PrePic     interface{} `json:"prePic"`
	DefaultVid string      `json:"defaultVid"`
	Screenshot struct {
		ImgUrl     string `json:"imgUrl"`
		WebImgUrl  string `json:"webImgUrl"`
		ImageSize  string `json:"imageSize"`
		Interval   int    `json:"interval"`
		MergeCount string `json:"mergeCount"`
	} `json:"screenshot"`
	FatherEpisodeId  int `json:"fatherEpisodeId"`
	MultiEpisodeInfo struct {
		MultiEpisodeCount int  `json:"multiEpisodeCount"`
		MultiEpisodeOrder int  `json:"multiEpisodeOrder"`
		MultiEpisode      bool `json:"multiEpisode"`
	} `json:"multiEpisodeInfo"`
	ExtraName            string      `json:"extraName"`
	AlbumIdStr           string      `json:"albumIdStr"`
	FatherEpisodeIdStr   string      `json:"fatherEpisodeIdStr"`
	FatherEpisodeIdOrder interface{} `json:"fatherEpisodeIdOrder"`
	AlbumDesc            interface{} `json:"albumDesc"`
	PCount               int         `json:"pCount"`
	PImgSize             interface{} `json:"pImgSize"`
	CImgSize             string      `json:"cImgSize"`
	CWebpImgSize         string      `json:"cWebpImgSize"`
}

func (i *Iqy) RefreshToken() (err error) {
	sprintf := fmt.Sprintf(registerMode, DeviceId)
	data, err := i.doRequest(sprintf, "token")
	if err != nil {
		return
	}
	token, err := data.String()
	if err != nil {
		return
	}
	jjj, _ := j.ParseJwtWithClaims(token)
	exp, err := jjj.GetExpirationTime()
	if err != nil {
		return
	}
	if exp.Sub(time.Now()) < time.Hour {
		err = fmt.Errorf("IQ RefreshToken Error: %v", "Token Expired")
		return
	}
	header["Authorization"] = "Bearer " + token
	for k, v := range header {
		client.Session.Headers.Set(k, v)
	}
	return
}
func (i *Iqy) GetEpisodeList(aid string) (r []*server.Video, err error) {
	u := fmt.Sprintf(GetEpisode, aid, "zh_cn", 0)
	data, err := i.doRequest(u, "")
	if err != nil {
		return
	}
	b, err := data.MarshalJSON()
	if err != nil {
		return

	}
	vList := EpgInfoIQ{}
	err = json.Unmarshal(b, &vList)
	if err != nil {
		return
	}
	var v []Epg
	v = append(v, vList.EpgS...)
	if vList.HasMore {
		//每页最多60
		var ch = make(chan []Epg, vList.Total/60+1)
		for n := 60; n < vList.Total; n += 60 {
			go func(n int) {
				u := fmt.Sprintf(GetEpisode, aid, "zh_cn", n)
				data, err := i.doRequest(u)
				if err != nil {
					ch <- nil
				}
				b, err := data.MarshalJSON()
				if err != nil {
					ch <- nil
				}
				list := EpgInfoIQ{}
				err = json.Unmarshal(b, &list)
				if err != nil {
					ch <- nil
				}
				ch <- list.EpgS
			}(n)
		}
		for n := 60; n < vList.Total; n += 60 {
			v = append(v, <-ch...)
		}
	}
	for _, e := range v {
		r = append(r, &server.Video{
			EpisodeId:    strconv.FormatInt(e.QipuId, 10),
			EpisodeTitle: e.Name,
			IsVip:        e.VipInfo != struct{}{},
			Episode:      uint32(e.Order),
			Extra: map[string]string{
				"subtitle": e.SubTitle,
			},
		})

	}
	return
}
func NewIq() *Iqy {
	i := &Iqy{}
	err := i.RefreshToken()
	if err != nil {
		return nil
	}
	return i
}
func (i *Iqy) GetIqMateInfo(sharerUrl string) (r *server.Data, err error) {
	r = &server.Data{}
	epgInfo, err := i.GetBaseInfo(r, sharerUrl, true)
	if err != nil {
		i.code = 1
		return
	}
	if epgInfo.ChnId == 1 {
		r.SeriesTitle = epgInfo.Name
		r.SeriesId = strconv.FormatInt(epgInfo.QipuId, 10)
		return
	} else {
		r.SeriesTitle = epgInfo.AlbumName
		r.SeriesId = strconv.Itoa(epgInfo.AlbumId)
		r.VideoList, err = i.GetEpisodeList(r.SeriesId)
		if err != nil {
			i.code = 1
			return
		}

	}
	return
}
func GetIqMateInfo(ctx context.Context, sharerUrl string) (r *server.Data, code int) {
	var (
		err error
	)
	r = &server.Data{}
	requestId := ctx.Value("requestId").(string)
	log.Debugf("site: %s, start GetMateInfo", web)
	i := NewIq()
	r, err = i.GetIqMateInfo(sharerUrl)
	if err != nil {
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
	}
	code = i.code
	return
}
