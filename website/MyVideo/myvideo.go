package MyVideo

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Tontonnow/ddddhmlist/config"
	"github.com/Tontonnow/ddddhmlist/server"
	"github.com/Tontonnow/ddddhmlist/utils/sesssion"
	"strings"
	"trpc.group/trpc-go/trpc-go/log"
)

var (
	web                 = "myvideo"
	client              = sesssion.NewClient(config.Conf.WebConfig[web])
	FindMyVideoContent6 = "https://mgw.myvideo.net.tw/twmsgw.api/FindMyVideoContent6.json?chl=android&chk=9e76ab&contentId=%s&isSeries=%s&devType=Handset"
	GetVideoListDataUrl = "https://myvideo.net.tw/ajax/ajaxGetVideoListData?twmContentId=%s"
)

type MyVideo struct {
	Name   string `json:"name"`
	Status struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"status"`
	Data Content6 `json:"data"`
}
type Content6 struct {
	Video2 struct {
		Id                  string `json:"id"`
		VideoTitle          string `json:"videoTitle"`
		Episode             string `json:"episode"`
		VolumnDesc          string `json:"volumnDesc"`
		VideoPic01          string `json:"videoPic01"`
		VideoPic02          string `json:"videoPic02"`
		VideoPic03          string `json:"videoPic03"`
		VideoPic04          string `json:"videoPic04"`
		VideoPic05          string `json:"videoPic05"`
		VideoPic06          string `json:"videoPic06"`
		VideoPic07          string `json:"videoPic07"`
		SeriesType          string `json:"seriesType"`
		SubTitle            string `json:"subTitle"`
		PublishYear         string `json:"publishYear"`
		MovieLength         string `json:"movieLength"`
		MovieLengthDesc     string `json:"movieLengthDesc"`
		IsHD                string `json:"isHD"`
		VideoQuality        string `json:"videoQuality"`
		InStoreService      string `json:"inStoreService"`
		MainCategory        string `json:"mainCategory"`
		VideoDesc           string `json:"videoDesc"`
		Language            string `json:"language"`
		PublishDate         string `json:"publishDate"`
		HasTrailer          string `json:"hasTrailer"`
		TrailerUrl          string `json:"trailerUrl"`
		TrailerUrl2         string `json:"trailerUrl2"`
		TrailerUrl3         string `json:"trailerUrl3"`
		Issuer              string `json:"issuer"`
		PrePurchaseStatus   string `json:"prePurchaseStatus"`
		IsSupportChromeCast string `json:"isSupportChromeCast"`
		IsDvdSync           string `json:"isDvdSync"`
		IsFree              string `json:"isFree"`
		IsDefaultVideo      string `json:"isDefaultVideo"`
		CanDownload         string `json:"canDownload"`
		YValue              string `json:"yValue"`
		ChromeCastPic       string `json:"chromeCastPic"`
		MainCategoryDesc    string `json:"mainCategoryDesc"`
		IsYoutube           string `json:"isYoutube"`
		YoutubeId           string `json:"youtubeId"`
		GradeCode           string `json:"gradeCode"`
		GradeCodeDesc       string `json:"gradeCodeDesc"`
		TrailerType         string `json:"trailerType"`
		IsFormal            string `json:"isFormal"`
		ActorChtList        struct {
			ActorChtName []string `json:"actorChtName"`
		} `json:"actorChtList"`
		DirectorChtList struct {
			DirectorChtName []string `json:"directorChtName"`
		} `json:"directorChtList"`
		SubCategoryList struct {
			Category []struct {
				CategoryCode string `json:"categoryCode"`
				CategoryName string `json:"categoryName"`
			} `json:"category"`
		} `json:"subCategoryList"`
		IsMultiLanguage  string `json:"isMultiLanguage"`
		SubtitleType     string `json:"subtitleType"`
		MainStill        string `json:"mainStill"`
		IsDolby          string `json:"isDolby"`
		IsAVOD           string `json:"isAVOD"`
		LikeCount        string `json:"likeCount"`
		DislikeCount     string `json:"dislikeCount"`
		DisplayStartDate string `json:"displayStartDate"`
		StillImageList   struct {
			StillImage []struct {
				StillUrl       string `json:"stillUrl"`
				StillUrl2      string `json:"stillUrl2"`
				StillUrl3      string `json:"stillUrl3"`
				IsMasterVision int    `json:"isMasterVision"`
			} `json:"stillImage"`
		} `json:"stillImageList"`
		CountryList struct {
			CountryName []string `json:"countryName"`
		} `json:"countryList"`
		OpeningInSec         int    `json:"openingInSec"`
		OpeningOutSec        int    `json:"openingOutSec"`
		EndingInSec          int    `json:"endingInSec"`
		EndingOutSec         int    `json:"endingOutSec"`
		KidMode              string `json:"kidMode"`
		ShowNewDirectorActor string `json:"showNewDirectorActor"`
		DirectorList         struct {
			Director []struct {
				PersonId    string `json:"personId"`
				ChtName     string `json:"chtName"`
				ClickAction string `json:"clickAction"`
			} `json:"director"`
		} `json:"directorList"`
		ActorList struct {
			Actor []struct {
				PersonId    string `json:"personId"`
				ChtName     string `json:"chtName"`
				ClickAction string `json:"clickAction"`
				EngName     string `json:"engName,omitempty"`
			} `json:"actor"`
		} `json:"actorList"`
		GoogleVideoType string `json:"googleVideoType"`
		Rating          struct {
			Type     interface{} `json:"type"`
			Title    interface{} `json:"title"`
			Score    interface{} `json:"score"`
			MaxScore interface{} `json:"maxScore"`
			Url      interface{} `json:"url"`
		} `json:"rating"`
		IsLive             string `json:"isLive"`
		Rank               int    `json:"rank"`
		RankDesc           string `json:"rankDesc"`
		OffShelfDate       string `json:"offShelfDate"`
		SeoUpdateTime      string `json:"seoUpdateTime"`
		IsDolbyVision      string `json:"isDolbyVision"`
		IsDolbyAtmos       string `json:"isDolbyAtmos"`
		SeoTitle           string `json:"seoTitle"`
		AuthorizeStartDate string `json:"authorizeStartDate"`
	} `json:"video2"`
	PurchaseInfoList struct {
		PurchaseInfo []struct {
			PurchaseType        string `json:"purchaseType"`
			PurchaseTitle       string `json:"purchaseTitle"`
			PurchaseDesc        string `json:"purchaseDesc"`
			PurchaseCaption     string `json:"purchaseCaption"`
			CanPurchase         string `json:"canPurchase"`
			PurchaseInfoStatus  string `json:"purchaseInfoStatus"`
			PurchaseActionTitle string `json:"purchaseActionTitle"`
		} `json:"purchaseInfo"`
	} `json:"purchaseInfoList"`
	DispWording  string `json:"dispWording"`
	DispWording2 string `json:"dispWording2"`
	SeriesInfo   struct {
		Id                string `json:"id"`
		TwmContentId      string `json:"twmContentId"`
		SeriesContentName string `json:"seriesContentName"`
		SeriesType        int    `json:"seriesType"`
		SeasonDispNumber  int    `json:"seasonDispNumber"`
	} `json:"seriesInfo"`
	ZeroPriceInfo struct {
		IsZeroPrice   string `json:"isZeroPrice"`
		ZeroPriceType int    `json:"zeroPriceType"`
	} `json:"zeroPriceInfo"`
	VideoPlayRight struct {
		PrePurchaseStatus string `json:"prePurchaseStatus"`
		HasPlayRight      string `json:"hasPlayRight"`
	} `json:"videoPlayRight"`
	ShowOutStreamAD string `json:"showOutStreamAD"`
	UpsellInfo      struct {
		IsShow string `json:"isShow"`
	} `json:"upsellInfo"`
	IsMultiScene string `json:"isMultiScene"`
}
type VideoListData struct {
	Data   string `json:"data"`
	Status string `json:"status"`
}
type Base64DDDD struct {
	BundleVideoList []struct {
		Id       string `json:"id"`
		DispName string `json:"dispName"`
	} `json:"bundleVideoList"`
	PackingId string `json:"packingId"`
}

func GetVideoListData(id string) (b Base64DDDD, err error) {
	mv := &VideoListData{}
	u := fmt.Sprintf(GetVideoListDataUrl, id)
	resp, err := client.Do("get", u, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp.Content, &mv)
	if err != nil {
		return
	}
	by, err := base64.StdEncoding.DecodeString(mv.Data)
	if err != nil {
		return
	}
	err = json.Unmarshal(by, &b)
	return
}
func GetVideoInfo(id string, isSeries string) (m Content6, err error) {
	mv := &MyVideo{}
	if isSeries != "0" {
		isSeries = "1"
	}
	u := fmt.Sprintf(FindMyVideoContent6, id, isSeries)
	resp, err := client.Do("get", u, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp.Content, &mv)
	m = mv.Data
	return
}
func GetMateInfo(ctx context.Context, sharerUrl string) (v *server.Data, code int) {
	requestId := ctx.Value("requestId").(string)
	log.Debugf("site: %s, start GetMateInfo", web)
	v = &server.Data{}
	us := strings.Split(sharerUrl, "/")
	isSeries := us[len(us)-2]
	contentId := us[len(us)-1]
	m, err := GetVideoInfo(contentId, isSeries)
	if err != nil {
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		code = 1
		return
	}
	v.SeriesTitle = m.Video2.VideoTitle
	v.SeriesId = m.Video2.Id
	v.IsVip = m.Video2.IsFree == "N"
	v.Extra = map[string]string{
		"videoQuality": m.Video2.VideoQuality,
	}
	v.IsSeries = isSeries != "0"
	if v.IsSeries {
		var BundleVideoList Base64DDDD
		BundleVideoList, err = GetVideoListData(m.SeriesInfo.TwmContentId)
		if err != nil {
			log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
			code = 1
			return
		}
		for _, b := range BundleVideoList.BundleVideoList {
			vv := &server.Video{
				EpisodeTitle: b.DispName,
				EpisodeId:    b.Id,
			}
			v.VideoList = append(v.VideoList, vv)
		}
	}
	return
}
