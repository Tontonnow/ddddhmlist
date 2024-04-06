package KKTV

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Tontonnow/ddddhmlist/config"
	"github.com/Tontonnow/ddddhmlist/server"
	"github.com/Tontonnow/ddddhmlist/utils/sesssion"
	"strings"
	"trpc.group/trpc-go/trpc-go/log"
)

var (
	web    = "kktv"
	client = sesssion.NewClient(config.Conf.WebConfig[web])
)

type Mate struct {
	Status struct {
		Type    string `json:"type"`
		Subtype string `json:"subtype"`
		Message string `json:"message"`
	} `json:"status"`
	Data struct {
		Id                            string        `json:"id"`
		Title                         string        `json:"title"`
		TitleType                     string        `json:"title_type"`
		Copyright                     string        `json:"copyright"`
		Cover                         string        `json:"cover"`
		Stills                        []string      `json:"stills"`
		ContentLabels                 []interface{} `json:"content_labels"`
		ContentLabelsForExpiredUser   []interface{} `json:"content_labels_for_expired_user"`
		ContentLabelsForFreetrialUser []interface{} `json:"content_labels_for_freetrial_user"`
		TitleAliases                  interface{}   `json:"title_aliases"`
		Available                     bool          `json:"available"`
		IsEnding                      bool          `json:"is_ending"`
		IsContainingAvod              bool          `json:"is_containing_avod"`
		ReverseDisplayOrder           bool          `json:"reverse_display_order"`
		IsValidated                   bool          `json:"is_validated"`
		FreeTrial                     bool          `json:"free_trial"`
		ChildLock                     bool          `json:"child_lock"`
		EndYear                       int           `json:"end_year"`
		ReleaseYear                   int           `json:"release_year"`
		TitleExtra                    struct {
			Id string `json:"id"`
		} `json:"title_extra"`
		Review struct {
			Content string `json:"content"`
		} `json:"review"`
		ReleaseInfo        string `json:"release_info"`
		Status             string `json:"status"`
		Summary            string `json:"summary"`
		LatestUpdateInfo   string `json:"latest_update_info"`
		TotalEpisodeCounts struct {
			Field1 int `json:"0600167301"`
		} `json:"total_episode_counts"`
		TotalSeriesCount int     `json:"total_series_count"`
		Rating           int     `json:"rating"`
		UserRatingCount  int     `json:"user_rating_count"`
		UserRating       float64 `json:"user_rating"`
		WikiOrig         string  `json:"wiki_orig"`
		WikiZh           string  `json:"wiki_zh"`
		Ost              struct {
			ArtistName string `json:"artist_name"`
			Image      string `json:"image"`
			Title      string `json:"title"`
			Url        string `json:"url"`
		} `json:"ost"`
		Country struct {
			Id             string `json:"id"`
			Title          string `json:"title"`
			CollectionType string `json:"collection_type"`
			CollectionName string `json:"collection_name"`
		} `json:"country"`
		ContentAgents []struct {
			Id             string `json:"id"`
			CollectionType string `json:"collection_type"`
			CollectionName string `json:"collection_name"`
		} `json:"content_agents"`
		ContentProviders interface{} `json:"content_providers"`
		Themes           []struct {
			Id             string `json:"id"`
			CollectionType string `json:"collection_type"`
			CollectionName string `json:"collection_name"`
		} `json:"themes"`
		Genres []struct {
			Id             string `json:"id"`
			Title          string `json:"title"`
			CollectionType string `json:"collection_type"`
			CollectionName string `json:"collection_name"`
		} `json:"genres"`
		Tags []struct {
			Id             string `json:"id"`
			CollectionType string `json:"collection_type"`
			CollectionName string `json:"collection_name"`
		} `json:"tags"`
		Directors []struct {
			Id             string `json:"id"`
			CollectionType string `json:"collection_type"`
			CollectionName string `json:"collection_name"`
		} `json:"directors"`
		Writers []struct {
			Id             string `json:"id"`
			CollectionType string `json:"collection_type"`
			CollectionName string `json:"collection_name"`
		} `json:"writers"`
		Producers []struct {
			Id             string `json:"id"`
			CollectionType string `json:"collection_type"`
			CollectionName string `json:"collection_name"`
		} `json:"producers"`
		Casts []struct {
			Id             string `json:"id"`
			CollectionType string `json:"collection_type"`
			CollectionName string `json:"collection_name"`
		} `json:"casts"`
		Series []struct {
			AvodEpisodeHint string `json:"avod_episode_hint"`
			Id              string `json:"id"`
			Title           string `json:"title"`
			Episodes        []struct {
				Id              string      `json:"id"`
				Title           string      `json:"title"`
				SeriesTitle     string      `json:"series_title"`
				SeriesId        string      `json:"series_id"`
				ContentAgent    string      `json:"content_agent"`
				ContentProvider string      `json:"content_provider"`
				Duration        float64     `json:"duration"`
				EndOffset       int         `json:"end_offset"`
				EndYear         int         `json:"end_year"`
				ReleaseYear     int         `json:"release_year"`
				StartYear       int         `json:"start_year"`
				Available       bool        `json:"available"`
				FreeTrial       bool        `json:"free_trial"`
				HasSubtitles    bool        `json:"has_subtitles"`
				IsEnding        bool        `json:"is_ending"`
				IsValidated     bool        `json:"is_validated"`
				OfflineMode     bool        `json:"offline_mode"`
				PlayZone        bool        `json:"play_zone"`
				IsAvod          bool        `json:"is_avod"`
				RoamingMode     bool        `json:"roaming_mode"`
				SecurePlayback  bool        `json:"secure_playback"`
				LicenseEnd      int         `json:"license_end"`
				LicenseStart    int         `json:"license_start"`
				PublishTime     int         `json:"publish_time"`
				Pub             int         `json:"pub"`
				Unpub           int         `json:"unpub"`
				Still           string      `json:"still"`
				Subtitles       interface{} `json:"subtitles"`
				Mezzanines      struct {
					Dash struct {
						Size  int64          `json:"size"`
						Sizes map[string]int `json:"sizes"`
						Uri   string         `json:"uri"`
					} `json:"dash"`
					Hls struct {
						Size  int64          `json:"size"`
						Sizes map[string]int `json:"sizes"`
						Uri   string         `json:"uri"`
					} `json:"hls"`
				} `json:"mezzanines"`
			} `json:"episodes"`
		} `json:"series"`
		UserActions struct {
		} `json:"user_actions"`
	} `json:"data"`
}

func ExtractIDFromURL(u string) (id string) {
	if u[:4] != "http" {
		return u
	}
	/*
		re := regexp.MustCompile(`https://kktv.me/titles/(\d+)`)
		matches := re.FindStringSubmatch(u)
		if len(matches) < 2 {
			return ""
		}
		return matches[1]
	*/
	us := strings.Split(u, "/")
	return us[len(us)-1][:8]

}

func GetMateInfo(ctx context.Context, sharerUrl string) (ret *server.Data, code int) {
	var (
		err error
	)
	requestId := ctx.Value("requestId").(string)
	log.Debugf("site: %s, start GetMateInfo", web)
	ret = &server.Data{}
	id := ExtractIDFromURL(sharerUrl)
	if id == "" {
		err = fmt.Errorf("Invalid URL")
		code = 2
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		return
	}
	url := fmt.Sprintf("https://api.kktv.me/v3/titles/%s", id)
	r, err := client.Do("get", url, nil)
	if err != nil {
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		code = 1
		return
	}
	var t Mate
	err = json.Unmarshal(r.Content, &t)
	if err != nil {
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		code = 1
		return
	}
	if t.Status.Type != "OK" {
		err = fmt.Errorf(t.Status.Message)
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		code = 1
		return
	}
	ret.SeriesTitle = t.Data.Title
	ret.SeriesId = t.Data.Id
	if t.Data.ContentLabelsForExpiredUser == nil {
		ret.IsVip = false
	}
	for _, v := range t.Data.Series {
		for _, e := range v.Episodes {
			ret.VideoList = append(ret.VideoList, &server.Video{
				EpisodeTitle: v.Title + " " + e.Title,
				EpisodeId:    e.Id,
				Extra: map[string]string{
					"dash": e.Mezzanines.Dash.Uri,
					"hls":  e.Mezzanines.Hls.Uri,
				},
				IsVip:   !e.PlayZone,
				Episode: uint32(int(e.Id[len(e.Id)-1]) - 48),
			})
		}
	}
	return

}
