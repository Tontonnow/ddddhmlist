package Hami

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Tontonnow/ddddhmlist/config"
	"github.com/Tontonnow/ddddhmlist/server"
	"github.com/Tontonnow/ddddhmlist/utils/sesssion"
	"regexp"
	"strconv"
	"trpc.group/trpc-go/trpc-go/log"
)

var (
	web            = "hami"
	getVodInfoById = "https://apl-hamivideo.cdn.hinet.net/HamiVideo/getVodInfoById.php?appVersion=&appOS=TV&productId="
	client         = sesssion.NewClient(config.Conf.WebConfig[web])
)

type VodInfo struct {
	ReturnCode    string `json:"returnCode"`
	ReturnMessage string `json:"returnMessage"`
	LastModified  string `json:"lastModified"`
	CacheTime     string `json:"cacheTime"`
	ProductInfo   struct {
		ProductId                     string        `json:"productId"`
		ProductName                   string        `json:"productName"`
		ProductNameEn                 string        `json:"productNameEn"`
		ProductType                   string        `json:"productType"`
		ContentId                     string        `json:"contentId"`
		ContentPk                     string        `json:"contentPk"`
		Description                   string        `json:"description"`
		ImageId                       string        `json:"imageId"`
		ListPrice                     string        `json:"listPrice"`
		StartTime                     string        `json:"startTime"`
		EndTime                       string        `json:"endTime"`
		UpdateOrder                   string        `json:"updateOrder"`
		FreeProduct                   string        `json:"freeProduct"`
		Category                      string        `json:"category"`
		Subcategory                   string        `json:"subcategory"`
		Length                        string        `json:"length"`
		SeriesDone                    string        `json:"seriesDone"`
		IsAdult                       string        `json:"isAdult"`
		AdditionalTag                 string        `json:"additionalTag"`
		ContentCount                  string        `json:"contentCount"`
		DisplaySeries                 string        `json:"displaySeries"`
		Score                         string        `json:"score"`
		ProgramInfo                   []interface{} `json:"programInfo"`
		PlayingTimes                  string        `json:"playingTimes"`
		ImdbRating                    string        `json:"imdbRating"`
		NoChildren                    string        `json:"noChildren"`
		IsHollywood                   string        `json:"isHollywood"`
		HollywoodVender               string        `json:"hollywoodVender"`
		TpPartner                     string        `json:"tpPartner"`
		VodType                       string        `json:"vodType"`
		LiveType                      string        `json:"liveType"`
		AllCategory                   []string      `json:"allCategory"`
		SeriesClientUpdate            string        `json:"seriesClientUpdate"`
		VodLiteInfo                   []interface{} `json:"vodLiteInfo"`
		AuthDevice                    string        `json:"authDevice"`
		TicketAllowed                 string        `json:"ticketAllowed"`
		DeviceControlDesc             string        `json:"deviceControlDesc"`
		Quality                       string        `json:"quality"`
		IsComingUp                    string        `json:"isComingUp"`
		ComingupStime                 string        `json:"comingupStime"`
		AuthRegion                    string        `json:"authRegion"`
		RelativeSalesProductGroupInfo []struct {
			ProductGroupType string `json:"productGroupType"`
			ProductGroupInfo []struct {
				ProductId   string `json:"productId"`
				ProductName string `json:"productName"`
				LowestPrice string `json:"lowestPrice"`
			} `json:"productGroupInfo"`
		} `json:"relativeSalesProductGroupInfo"`
		TermsVersionId  string `json:"termsVersionId"`
		ExtraTag        string `json:"extraTag"`
		ReleasePlatform string `json:"releasePlatform"`
		FeedInfo        struct {
			FeedProductId string `json:"feedProductId"`
			FeedType      string `json:"feedType"`
			DeepLink      struct {
				DesktopWebPlatform string `json:"DesktopWebPlatform"`
				MobileWebPlatform  string `json:"MobileWebPlatform"`
				IOSPlatform        string `json:"IOSPlatform"`
				AndroidPlatform    string `json:"AndroidPlatform"`
				AndroidTVPlatform  string `json:"AndroidTVPlatform"`
			} `json:"deepLink"`
			SeriesName    string `json:"seriesName"`
			EpisodeNumber string `json:"episodeNumber"`
			SeasonNumber  string `json:"seasonNumber"`
		} `json:"feedInfo"`
		IsSpeedup   string `json:"isSpeedup"`
		ContentInfo []struct {
			ContentId           string `json:"contentId"`
			ContentPk           string `json:"contentPk"`
			ContentTitle        string `json:"contentTitle"`
			ContentTitleEn      string `json:"contentTitleEn"`
			ContractStart       string `json:"contractStart"`
			ContractEnd         string `json:"contractEnd"`
			ContentType         string `json:"contentType"`
			TpPartner           string `json:"tpPartner"`
			DisplayOrder        string `json:"displayOrder"`
			AuthDevice          string `json:"authDevice"`
			IfFree              string `json:"ifFree"`
			ProductId           string `json:"productId"`
			RelativeProductInfo []struct {
				ProductId   string `json:"productId"`
				ProductName string `json:"productName"`
				LowestPrice string `json:"lowestPrice"`
			} `json:"relativeProductInfo"`
			RelativeProductGroupInfo []struct {
				ProductGroupType string `json:"productGroupType"`
				ProductGroupInfo []struct {
					ProductId   string `json:"productId"`
					ProductName string `json:"productName"`
					LowestPrice string `json:"lowestPrice"`
				} `json:"productGroupInfo"`
			} `json:"relativeProductGroupInfo"`
			Metadata struct {
				SeriesNumber     string        `json:"seriesNumber"`
				Description      string        `json:"description"`
				Language         string        `json:"language"`
				Audio            string        `json:"audio"`
				Rating           string        `json:"rating"`
				Length           string        `json:"length"`
				Actors           string        `json:"actors"`
				ActorsList       []string      `json:"actorsList"`
				Director         string        `json:"director"`
				DirectorList     []interface{} `json:"directorList"`
				Screenwriter     string        `json:"screenwriter"`
				ScreenwriterList []interface{} `json:"screenwriterList"`
				Subtitle         string        `json:"subtitle"`
				PublishYear      string        `json:"publishYear"`
				IsAdult          string        `json:"isAdult"`
				PosterURL        string        `json:"posterURL"`
				PosterImageId    string        `json:"posterImageId"`
				VenderId         string        `json:"venderId"`
				DrmProtect       string        `json:"drmProtect"`
				Category         string        `json:"category"`
				Subcategory      string        `json:"subcategory"`
				Country          string        `json:"country"`
				Comment          string        `json:"comment"`
				Quality          string        `json:"quality"`
				VenderName       string        `json:"venderName"`
				PreviewInfo      []struct {
					Quality string `json:"quality"`
					Format  string `json:"format"`
					Device  string `json:"device"`
					Bitrate string `json:"bitrate"`
					Drm     string `json:"drm"`
					Url     string `json:"url"`
				} `json:"previewInfo"`
				VodPosterInfo     []interface{} `json:"vodPosterInfo"`
				HistoryTimes      string        `json:"historyTimes"`
				ThirtydaysTimes   string        `json:"thirtydaysTimes"`
				WeekTimes         string        `json:"weekTimes"`
				ImdbRating        string        `json:"imdbRating"`
				DisplayVolume     string        `json:"displayVolume"`
				NoChildren        string        `json:"noChildren"`
				IsHollywood       string        `json:"isHollywood"`
				HollywoodVender   string        `json:"hollywoodVender"`
				VodType           string        `json:"vodType"`
				MultiviewDescUrl  string        `json:"multiviewDescUrl"`
				VodLabel          string        `json:"vodLabel"`
				TitleContainsDone string        `json:"titleContainsDone"`
				DeviceControlDesc string        `json:"deviceControlDesc"`
				IsSpeedup         string        `json:"isSpeedup"`
				TagImageId        string        `json:"tagImageId"`
			} `json:"metadata"`
		} `json:"contentInfo"`
		ProductInfo   []interface{} `json:"productInfo"`
		CategoryStyle string        `json:"categoryStyle"`
		RecommendVod  []struct {
			VodType         string `json:"vodType"`
			IsLust          string `json:"isLust"`
			NoChildren      string `json:"noChildren"`
			ActionType      string `json:"actionType"`
			SubType         string `json:"subType"`
			ProductType     string `json:"productType"`
			CategoryStyle   string `json:"categoryStyle"`
			ProductId       string `json:"productId"`
			ContentPk       string `json:"contentPk"`
			TsId            string `json:"tsId"`
			FunctionId      string `json:"functionId"`
			ImageId         string `json:"imageId"`
			Title           string `json:"title"`
			SubTitle        string `json:"subTitle"`
			DescText        string `json:"descText"`
			ProductName     string `json:"productName"`
			Description     string `json:"description"`
			Length          string `json:"length"`
			DisplaySeries   string `json:"displaySeries"`
			Percent         string `json:"percent"`
			ImdbRating      string `json:"imdbRating"`
			Score           string `json:"score"`
			StartTime       string `json:"startTime"`
			EndTime         string `json:"endTime"`
			IsPlay          string `json:"isPlay"`
			Link            string `json:"link"`
			TagImageId      string `json:"tagImageId"`
			TagImageIdRD    string `json:"tagImageId_RD"`
			AdditionalTag   string `json:"additionalTag"`
			ListPrice       string `json:"listPrice"`
			FreeProduct     string `json:"freeProduct"`
			Category        string `json:"category"`
			Subcategory     string `json:"subcategory"`
			SeriesDone      string `json:"seriesDone"`
			IsAdult         string `json:"isAdult"`
			ContentCount    string `json:"contentCount"`
			PlayingTimes    string `json:"playingTimes"`
			CompetitionText string `json:"competitionText"`
			CountryInfo     struct {
			} `json:"countryInfo"`
			ProgramInfo []interface{} `json:"programInfo"`
		} `json:"recommendVod"`
		RelativeVod            []interface{} `json:"relativeVod"`
		ScoreMessage           string        `json:"scoreMessage"`
		AuthDeviceDesc         string        `json:"authDeviceDesc"`
		RecommendVodFireBaseId string        `json:"recommendVodFireBaseId"`
		RelativeVodFireBaseId  string        `json:"relativeVodFireBaseId"`
		WarningInfo            struct {
			Title string `json:"title"`
			App   string `json:"app"`
			Web   string `json:"web"`
			Tv    string `json:"tv"`
		} `json:"warningInfo"`
		ModPaymentInfo string `json:"modPaymentInfo"`
	} `json:"productInfo"`
}

func GetVodInfoById(productId string) (vodInfo VodInfo, err error) {
	resp, err := client.Do("get", getVodInfoById+productId, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp.Content, &vodInfo)
	if err != nil {
		return
	}
	return
}
func GetProductId(url string) string {
	re := regexp.MustCompile(`/(\d+).do`)
	productId := re.FindStringSubmatch(url)
	if len(productId) <= 1 {
		return ""
	}
	return productId[1]
}
func GetMateInfo(ctx context.Context, sharerUrl string) (r *server.Data, code int) {
	var (
		err error
	)
	r = &server.Data{}
	requestId := ctx.Value("requestId").(string)
	log.Debugf("site: %s, start GetMateInfo", web)
	productId := GetProductId(sharerUrl)
	if productId == "" {
		err = fmt.Errorf("GET productId Error")
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		code = 2
		return
	}
	vodInfo, err := GetVodInfoById(productId)
	if err != nil {
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		code = 1
		return
	}
	if vodInfo.ReturnCode != "1" {
		err = fmt.Errorf(vodInfo.ReturnMessage)
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		code = 1
		return
	}
	r.SeriesTitle = vodInfo.ProductInfo.ProductName
	r.SeriesId = vodInfo.ProductInfo.ProductId
	r.Extra = map[string]string{
		"contentPk": vodInfo.ProductInfo.ContentPk,
		"quality":   vodInfo.ProductInfo.Quality,
	}
	for _, vod := range vodInfo.ProductInfo.ContentInfo {
		e, _ := strconv.Atoi(vod.Metadata.SeriesNumber)
		r.VideoList = append(r.VideoList, &server.Video{
			EpisodeTitle: vod.ContentTitle,
			EpisodeId:    vod.ContentId,
			IsVip:        vod.IfFree == "0",
			Episode:      uint32(e),
			Extra: map[string]string{
				"quality":   vod.Metadata.Quality,
				"contentPk": vod.ContentPk,
			},
		})
	}
	return

}
