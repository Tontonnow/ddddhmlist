package Viu

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/Tontonnow/ddddhmlist/config"
	"github.com/Tontonnow/ddddhmlist/server"
	"github.com/Tontonnow/ddddhmlist/utils/sesssion"
	"github.com/wangluozhe/requests/url"
	"regexp"
	"strconv"
	"trpc.group/trpc-go/trpc-go/log"
)

var (
	web            = "viu"
	client         = sesssion.NewClient(config.Conf.WebConfig[web])
	languageFlagId = []int{1, 2, 3}
)

type VIU struct {
	languageFlagId int
	country        string
}
type ProductListResp struct {
	Server struct {
		Time int `json:"time"`
		Area struct {
			AreaId   int `json:"area_id"`
			Language []struct {
				LanguageFlagId string `json:"language_flag_id"`
				Label          string `json:"label"`
				Mark           string `json:"mark"`
				IsDefault      string `json:"is_default"`
			} `json:"language"`
			Country struct {
				Code string `json:"code"`
				Id   string `json:"id"`
			} `json:"country"`
			Vuclip bool `json:"vuclip"`
		} `json:"area"`
	} `json:"server"`
	Status struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"status"`
	Data struct {
		ProductList []ProductList `json:"product_list"`
	} `json:"data"`
}
type ProductList struct {
	ProductId                    string `json:"product_id"`
	Number                       string `json:"number"`
	Synopsis                     string `json:"synopsis"`
	ScheduleStartTime            string `json:"schedule_start_time"`
	ScheduleEndTime              string `json:"schedule_end_time"`
	CoverImageUrl                string `json:"cover_image_url"`
	SeriesCategoryName           string `json:"series_category_name"`
	IsParentalLockLimited        string `json:"is_parental_lock_limited"`
	IsParentalLockCompulsory     string `json:"is_parental_lock_compulsory"`
	Description                  string `json:"description"`
	AllowDownload                string `json:"allow_download"`
	OfflineTime                  string `json:"offline_time"`
	FreeTime                     int    `json:"free_time"`
	PremiumTime                  int    `json:"premium_time"`
	IsFreePremiumTime            int    `json:"is_free_premium_time"`
	UserLevel                    int    `json:"user_level"`
	PosterLogoUrl                string `json:"poster_logo_url"`
	SourceFlag                   string `json:"source_flag"`
	AllowTv                      string `json:"allow_tv"`
	IsMovie                      int    `json:"is_movie"`
	AllowTelstb                  string `json:"allow_telstb"`
	ReleasedProductTotal         int    `json:"released_product_total"`
	CoverLandscapeImageUrl       string `json:"cover_landscape_image_url"`
	CoverPortraitImageUrl        string `json:"cover_portrait_image_url"`
	SeriesCoverLandscapeImageUrl string `json:"series_cover_landscape_image_url"`
	SeriesCoverPortraitImageUrl  string `json:"series_cover_portrait_image_url"`
	TimeDuration                 string `json:"time_duration"`
	CcsProductId                 string `json:"ccs_product_id"`
	SeriesId                     string `json:"series_id"`
}
type VodDetail struct {
	/*
			Server struct {
			Time int `json:"time"`
			Area struct {
				AreaId   int `json:"area_id"`
				Language []struct {
					LanguageFlagId string `json:"language_flag_id"`
					Label          string `json:"label"`
					Mark           string `json:"mark"`
					IsDefault      string `json:"is_default"`
				} `json:"language"`
				Country struct {
					Code string `json:"code"`
					Id   string `json:"id"`
				} `json:"country"`
				Vuclip bool `json:"vuclip"`
			} `json:"area"`
		} `json:"server"`
	*/
	Data struct {
		CurrentProduct struct {
			SeriesId           string `json:"series_id"`
			ProductId          string `json:"product_id"`
			Number             string `json:"number"`
			Synopsis           string `json:"synopsis"`
			Description        string `json:"description"`
			ScheduleStartTime  string `json:"schedule_start_time"`
			ScheduleEndTime    string `json:"schedule_end_time"`
			SkipIntroStartTime int    `json:"skip_intro_start_time"`
			SkipIntroEndTime   int    `json:"skip_intro_end_time"`
			FreeTime           int    `json:"free_time"`
			PremiumTime        int    `json:"premium_time"`
			IsFreePremiumTime  int    `json:"is_free_premium_time"`
			DurationStart      int    `json:"duration_start"`
			CoverImageUrl      string `json:"cover_image_url"`
			CcsProductId       string `json:"ccs_product_id"`
			AllowDownload      string `json:"allow_download"`
			ShareUrl           string `json:"share_url"`
			Subtitle           []struct {
				IsDefault                 int    `json:"is_default"`
				Name                      string `json:"name"`
				Url                       string `json:"url"`
				SubtitleUrl               string `json:"subtitle_url"`
				ProductSubtitleId         string `json:"product_subtitle_id"`
				ProductSubtitleLanguageId string `json:"product_subtitle_language_id"`
				SecondSubtitleUrl         string `json:"second_subtitle_url"`
				SecondSubtitlePosition    int    `json:"second_subtitle_position"`
				Code                      string `json:"code"`
			} `json:"subtitle"`
			Focus []interface{} `json:"focus"`
			Ad    []struct {
				Position int `json:"position"`
				CodeList []struct {
					AdStuff     string      `json:"ad_stuff"`
					AdTrack     interface{} `json:"ad_track"`
					AdIsUserPay string      `json:"ad_is_user_pay"`
					ImaForce    string      `json:"ima_force"`
				} `json:"code_list"`
				StartTime string `json:"start_time"`
			} `json:"ad"`
			IsMovie                      int           `json:"is_movie"`
			IsParentalLockLimited        string        `json:"is_parental_lock_limited"`
			IsParentalLockCompulsory     string        `json:"is_parental_lock_compulsory"`
			AllowPlayBigScreen           string        `json:"allow_play_big_screen"`
			PlayBigScreenStartTime       string        `json:"play_big_screen_start_time"`
			PlayBigScreenEndTime         string        `json:"play_big_screen_end_time"`
			TimeDuration                 string        `json:"time_duration"`
			CampaignName                 string        `json:"campaign_name"`
			UserLevel                    int           `json:"user_level"`
			SeoTitle                     string        `json:"seo_title"`
			SeoDescription               string        `json:"seo_description"`
			HasContentWindow             string        `json:"has_content_window"`
			Classification               interface{}   `json:"classification"`
			ClassificationUrl            interface{}   `json:"classification_url"`
			ContentAdvisory              string        `json:"content_advisory"`
			CensorshipAdsMp4Url          string        `json:"censorship_ads_mp4_url"`
			ClientLogo                   interface{}   `json:"client_logo"`
			Keyword                      string        `json:"keyword"`
			OfflineTime                  string        `json:"offline_time"`
			PosterLogoUrl                string        `json:"poster_logo_url"`
			SourceFlag                   string        `json:"source_flag"`
			AllowTv                      string        `json:"allow_tv"`
			AllowTelstb                  string        `json:"allow_telstb"`
			ReleasedProductTotal         int           `json:"released_product_total"`
			CoverLandscapeImageUrl       string        `json:"cover_landscape_image_url"`
			CoverPortraitImageUrl        string        `json:"cover_portrait_image_url"`
			SeriesCoverLandscapeImageUrl string        `json:"series_cover_landscape_image_url"`
			SeriesCoverPortraitImageUrl  string        `json:"series_cover_portrait_image_url"`
			AllowChromecastPlayBigScreen int           `json:"allow_chromecast_play_big_screen"`
			AllowAirplayPlayBigScreen    int           `json:"allow_airplay_play_big_screen"`
			Overlay                      interface{}   `json:"overlay"`
			ProductTag                   []interface{} `json:"product_tag"`
			LShapeAd                     interface{}   `json:"l_shape_ad"`
		} `json:"current_product"`
		Series Series `json:"series"`
	} `json:"data"`
	Error struct {
		Message string `json:"message"`
		Code    string `json:"code"`
	} `json:"error"`
}

type Series struct {
	ProductList                  []ProductList `json:"product_list"`
	SeriesId                     string        `json:"series_id"`
	Name                         string        `json:"name"`
	ContentAdvisory              string        `json:"content_advisory"`
	AllowTelstb                  string        `json:"allow_telstb"`
	Description                  string        `json:"description"`
	ProductTotal                 string        `json:"product_total"`
	CategoryName                 string        `json:"category_name"`
	CategoryId                   string        `json:"category_id"`
	OttCate                      string        `json:"ott_cate"`
	ReleaseTime                  string        `json:"release_time"`
	ScheduleStartTime            string        `json:"schedule_start_time"`
	ScheduleEndTime              string        `json:"schedule_end_time"`
	CoverImageUrl                string        `json:"cover_image_url"`
	CoverLandscapeImageUrl       string        `json:"cover_landscape_image_url"`
	CoverPortraitImageUrl        string        `json:"cover_portrait_image_url"`
	UpdateCycleDescription       string        `json:"update_cycle_description"`
	CpLogoUrl                    string        `json:"cp_logo_url"`
	CpName                       string        `json:"cp_name"`
	SeriesLanguage               string        `json:"series_language"`
	AllowChromecastPlayBigScreen string        `json:"allow_chromecast_play_big_screen"`
	AllowAirplayPlayBigScreen    string        `json:"allow_airplay_play_big_screen"`
	IsWatermark                  string        `json:"is_watermark"`
	WatermarkPosition            int           `json:"watermark_position"`
	WatermarkUrl                 string        `json:"watermark_url"`
	AllowTv                      string        `json:"allow_tv"`
	SeoTitle                     string        `json:"seo_title"`
	SeoDescription               string        `json:"seo_description"`
	ReleaseOfYear                string        `json:"release_of_year"`
	Actor                        []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"actor"`
	SeriesTag []struct {
		Id   int    `json:"id"`
		Type string `json:"type"`
		Tags []struct {
			Id    string `json:"id,omitempty"`
			TagId string `json:"tag_id"`
			Name  string `json:"name"`
			Type  string `json:"type,omitempty"`
		} `json:"tags"`
	} `json:"series_tag"`
}
type StreamResp struct {
	Status struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"status"`
	Server struct {
		Time int `json:"time"`
	} `json:"server"`
	Data struct {
		Stream struct {
			Url struct {
				S240P  string `json:"s240p"`
				S480P  string `json:"s480p"`
				S720P  string `json:"s720p"`
				S1080P string `json:"s1080p"`
			} `json:"url"`
			Duration      string `json:"duration"`
			Cdn           string `json:"cdn"`
			Region        string `json:"region"`
			IspName       string `json:"ispName"`
			Ratio         string `json:"ratio"`
			DurationStart int    `json:"duration_start"`
		} `json:"stream"`
	} `json:"data"`
}

func (v *VIU) GetSeries(url string) (series Series, err error) {
	productId, country := v.ParseUrl(url)
	v.country = country
	var se any
	for _, languageFlagId := range languageFlagId {
		v.languageFlagId = languageFlagId
		se, err = v.GetViuSeries(productId, true)
		if err != nil {
			return
		}
		series = se.(Series)
		if series.SeriesId != "" {
			break
		}
	}
	if series.SeriesId == "" {
		err = errors.New("series not found")
	}
	return

}
func (v *VIU) GetViuSeries(productId string, isSeries bool) (any, error) {
	req := url.NewRequest()
	u := "https://api-gateway-global.viu.com/api/mobile"
	params := map[string]string{
		"r":                   "/vod/detail",
		"product_id":          productId,
		"platform_flag_label": "pad",
		"language_flag_id":    strconv.Itoa(v.languageFlagId),
		"ut":                  "0",
		"area_id":             "1",
		"os_flag_id":          "2",
		"countryCode":         "HK",
	}
	client.Session.Proxies = config.Conf.WebConfig[web].Proxy
	if v.country == "sg" {
		params["language_flag_id"] = strconv.Itoa(v.languageFlagId)
		params["area_id"] = "2"
		params["countryCode"] = "SG"
		client.Session.Proxies = config.Conf.Proxy["sg"]
	}
	req.Params = url.ParseParams(params)
	resp, err := client.Do("get", u, req)
	if err != nil {
		return nil, err
	}
	var vodDetail VodDetail
	err = json.Unmarshal(resp.Content, &vodDetail)
	if err != nil {
		return nil, err
	}
	if vodDetail.Error.Message != "" {
		return nil, errors.New(vodDetail.Error.Message)
	}
	if isSeries {
		series := vodDetail.Data.Series
		seriesID := series.SeriesId
		params["r"] = "/vod/product-list"
		params["series_id"] = seriesID
		params["size"] = "1000"
		req.Params = url.ParseParams(params)
		resp, err = client.Do("get", u, req)
		if err != nil {
			return nil, err
		}
		var productListResp ProductListResp
		err = json.Unmarshal(resp.Content, &productListResp)
		if err != nil {
			return nil, err
		}
		series.ProductList = productListResp.Data.ProductList
		return series, nil
	}
	return vodDetail.Data.CurrentProduct, nil
}
func (v *VIU) ParseUrl(url string) (productId, country string) {
	re := regexp.MustCompile(`/(\d+)`)
	productId = re.FindStringSubmatch(url)[1]
	re = regexp.MustCompile(`/ott/(\w+)/`)
	country = re.FindStringSubmatch(url)[1]
	if country == "" {
		country = "hk"
	}
	return productId, country
}
func GetMateInfo(ctx context.Context, sharerUrl string) (r *server.Data, code int) {
	var err error
	requestId := ctx.Value("requestId").(string)
	log.Debugf("site: %s, start GetMateInfo", web)
	r = &server.Data{}
	v := &VIU{}
	s, err := v.GetSeries(sharerUrl)
	if err != nil {
		code = 1
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		if err.Error() == "Device is out of region" {
			code = 3
		}
		return
	}
	r.SeriesTitle = s.Name
	r.SeriesId = s.SeriesId
	for _, v := range s.ProductList {
		r.VideoList = append(r.VideoList, &server.Video{
			EpisodeId:    v.ProductId,
			EpisodeTitle: v.Number + "_" + v.Synopsis,
		})
	}
	return
}
