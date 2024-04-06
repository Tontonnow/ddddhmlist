package MytvSuper

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Tontonnow/ddddhmlist/config"
	"github.com/Tontonnow/ddddhmlist/server"
	"github.com/Tontonnow/ddddhmlist/utils/sesssion"
	"regexp"
	"sort"
	"strconv"
	"time"
	"trpc.group/trpc-go/trpc-go/log"
)

var (
	web               = "mytvsuper"
	GetEpisodeListUrl = "https://content-api.mytvsuper.com/v1/episode/list?programme_id=%s&start_episode_no=1&end_episode_no=%d&sort_desc=true&platform=web"
	client            = sesssion.NewClient(config.Conf.WebConfig[web])
	GetDetailsUrl     = "https://content-api.mytvsuper.com/v1/programme/details?programme_id=%s&platform=web"
)

type Details struct {
	Error         string    `json:"Error"`
	ProgrammeId   int       `json:"programme_id"`
	NameTc        string    `json:"name_tc"`
	NameEn        string    `json:"name_en"`
	Path          string    `json:"path"`
	PayStartTime  time.Time `json:"pay_start_time"`
	PayEndTime    time.Time `json:"pay_end_time"`
	FreeStartTime time.Time `json:"free_start_time"`
	FreeEndTime   time.Time `json:"free_end_time"`
	StartTime     time.Time `json:"start_time"`
	EndTime       time.Time `json:"end_time"`
	Enabled       bool      `json:"enabled"`
	ShortDescTc   string    `json:"short_desc_tc"`
	ShortDescEn   string    `json:"short_desc_en"`
	LongDescTc    string    `json:"long_desc_tc"`
	LongDescEn    string    `json:"long_desc_en"`
	Image         struct {
		PortraitSmall   string `json:"portrait_small"`
		PortraitMedium  string `json:"portrait_medium"`
		PortraitLarge   string `json:"portrait_large"`
		LandscapeSmall  string `json:"landscape_small"`
		LandscapeMedium string `json:"landscape_medium"`
		LandscapeLarge  string `json:"landscape_large"`
	} `json:"image"`
	VersionType     string `json:"version_type"`
	ParentalLock    bool   `json:"parental_lock"`
	NumOfEpisodes   int    `json:"num_of_episodes"`
	LatestEpisodeNo int    `json:"latest_episode_no"`
	Tags            []struct {
		TagId  int    `json:"tag_id"`
		Type   string `json:"type"`
		NameTc string `json:"name_tc"`
		NameEn string `json:"name_en"`
	} `json:"tags"`
	Artists []struct {
		RefType string `json:"ref_type"`
		RefId   int    `json:"ref_id"`
		NameTc  string `json:"name_tc"`
		NameEn  string `json:"name_en"`
	} `json:"artists"`
	DirectorsTc   []interface{} `json:"directors_tc"`
	DirectorsEn   []interface{} `json:"directors_en"`
	EpisodeGroups []struct {
		StartEpisodeNo int    `json:"start_episode_no"`
		EndEpisodeNo   int    `json:"end_episode_no"`
		GroupNameTc    string `json:"group_name_tc"`
		GroupNameEn    string `json:"group_name_en"`
	} `json:"episode_groups"`
	DisplayEpiTitleOnly bool     `json:"display_epi_title_only"`
	ShowOffshelf        bool     `json:"show_offshelf"`
	Adunit              string   `json:"adunit"`
	AdtargetGenre       []string `json:"adtarget_genre"`
	AdtargetCategory    []string `json:"adtarget_category"`
	AdtargetSubCategory []string `json:"adtarget_sub_category"`
	AdCustomParams      struct {
		Brand             string `json:"brand"`
		Channel           string `json:"channel"`
		Contentrelated    string `json:"contentrelated"`
		Countryoforigin   string `json:"countryoforigin"`
		Decade            string `json:"decade"`
		Familycode        string `json:"familycode"`
		Othercat          string `json:"othercat"`
		Prodyear          string `json:"prodyear"`
		Programmetemplate string `json:"programmetemplate"`
		Seasonal          string `json:"seasonal"`
		Shootinglocation  string `json:"shootinglocation"`
		Topical           string `json:"topical"`
		Versiontype       string `json:"versiontype"`
	} `json:"ad_custom_params"`
	SeoDescTc           string        `json:"seo_desc_tc"`
	SeoDescEn           string        `json:"seo_desc_en"`
	ShortClips          []interface{} `json:"short_clips"`
	CallAd              bool          `json:"call_ad"`
	RecommendOtherTitle bool          `json:"recommend_other_title"`
	GoldContent         bool          `json:"gold_content"`
	ProfileClass        []string      `json:"profile_class"`
	ModifiedAt          time.Time     `json:"modified_at"`
	LabellingGroup      []interface{} `json:"labelling_group"`
}
type Episode struct {
	Error string `json:"Error"`
	Items []struct {
		EpisodeId     int           `json:"episode_id"`
		VideoId       int           `json:"video_id"`
		NameTc        string        `json:"name_tc"`
		NameEn        string        `json:"name_en"`
		DescTc        string        `json:"desc_tc"`
		DescEn        string        `json:"desc_en"`
		EpisodeNo     int           `json:"episode_no"`
		AllowDownload interface{}   `json:"allow_download"`
		PayStartTime  time.Time     `json:"pay_start_time"`
		PayEndTime    time.Time     `json:"pay_end_time"`
		FreeStartTime interface{}   `json:"free_start_time"`
		FreeEndTime   interface{}   `json:"free_end_time"`
		GoldStartTime interface{}   `json:"gold_start_time"`
		GoldEndTime   interface{}   `json:"gold_end_time"`
		EarlyRelease  []interface{} `json:"early_release"`
		IsPreview     interface{}   `json:"is_preview"`
		Image         struct {
			Small  string `json:"small"`
			Medium string `json:"medium"`
			Large  string `json:"large"`
		} `json:"image"`
		Duration float64 `json:"duration"`
	} `json:"items"`
}

func getProgrammeId(u string) string {
	re, _ := regexp.Compile(`_(\d+)/`)
	programmeId := re.FindStringSubmatch(u)
	if len(programmeId) == 0 {
		return ""
	}
	return programmeId[1]
}
func EpisodeList(programmeId string, end int) (e Episode, err error) {
	url := fmt.Sprintf(GetEpisodeListUrl, programmeId, end)
	rsp, err := client.Do("get", url, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(rsp.Content, &e)
	if e.Error != "" {
		err = fmt.Errorf(e.Error)
		return
	}
	return
}
func getDetails(programmeId string) (d Details, err error) {
	url := fmt.Sprintf(GetDetailsUrl, programmeId)
	rsp, _ := client.Do("get", url, nil)
	err = json.Unmarshal(rsp.Content, &d)
	if err != nil {
		return
	}
	if d.Error != "" {
		err = fmt.Errorf(d.Error)
		return
	}
	return

}
func GetMateInfo(ctx context.Context, sharerUrl string) (r *server.Data, code int) {
	requestId := ctx.Value("requestId").(string)
	log.Debugf("site: %s, start GetMateInfo", web)
	r = &server.Data{}
	r = &server.Data{}
	v := getProgrammeId(sharerUrl)
	if v == "" {
		err := fmt.Errorf("invalid URL")
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		code = 2
		return

	}
	d, err := getDetails(v)
	if err != nil {
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		code = 1
		return
	}
	r.SeriesTitle = d.NameTc
	r.SeriesId = strconv.Itoa(d.ProgrammeId)
	r.Extra = map[string]string{
		"name_en": d.NameEn,
	}
	e, err := EpisodeList(v, d.LatestEpisodeNo)
	if err != nil {
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		code = 1
		return
	}
	for _, v := range e.Items {
		r.VideoList = append(r.VideoList, &server.Video{
			EpisodeTitle: v.NameTc,
			EpisodeId:    strconv.Itoa(v.EpisodeId),
			Episode:      uint32(v.EpisodeNo),
			Extra: map[string]string{
				"name_en":  v.NameEn,
				"video_id": strconv.Itoa(v.VideoId),
			},
		})
	}
	sort.Sort(r)
	return
}
