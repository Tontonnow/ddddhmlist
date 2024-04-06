package VIKI

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Tontonnow/ddddhmlist/config"
	"github.com/Tontonnow/ddddhmlist/server"
	"github.com/Tontonnow/ddddhmlist/utils/sesssion"
	"regexp"
	"time"
	"trpc.group/trpc-go/trpc-go/log"
)

var (
	web            = "viki"
	client         = sesssion.NewClient(config.Conf.WebConfig[web])
	GetVideosUrl   = "https://api.viki.io/v4/videos/%s.json?app=100004a&token="
	GetEpisodesUrl = "https://api.viki.io/v4/containers/%s/episodes.json?token=&direction=asc&with_upcoming=true&sort=number&blocked=true&page=1&per_page=100&app=100000a"
)

type Videos struct {
	Error         string `json:"error"`
	Vcode         int    `json:"vcode"`
	Id            string `json:"id"`
	Type          string `json:"type"`
	Subtype       string `json:"subtype"`
	ContentOwners []struct {
		Id string `json:"id"`
	} `json:"content_owners"`
	Origin struct {
		Country  string `json:"country"`
		Language string `json:"language"`
	} `json:"origin"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Subscriptions struct {
		Count int `json:"count"`
	} `json:"subscriptions"`
	Team struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"team"`
	Titles struct {
		Fr string `json:"fr"`
		Ro string `json:"ro"`
		Zt string `json:"zt"`
		Zh string `json:"zh"`
		En string `json:"en"`
		Cs string `json:"cs"`
		It string `json:"it"`
		Es string `json:"es"`
		Pt string `json:"pt"`
		Ko string `json:"ko"`
	} `json:"titles"`
	Images struct {
		Poster struct {
			Url    string `json:"url"`
			Source string `json:"source"`
		} `json:"poster"`
		AtvCover struct {
			Url    string `json:"url"`
			Source string `json:"source"`
		} `json:"atv_cover"`
		LatestVideo struct {
			Poster struct {
				Url    string `json:"url"`
				Source string `json:"source"`
			} `json:"poster"`
		} `json:"latest_video"`
	} `json:"images"`
	Descriptions struct {
		Fr string `json:"fr"`
		Cs string `json:"cs"`
		Zt string `json:"zt"`
		Zh string `json:"zh"`
		It string `json:"it"`
		Es string `json:"es"`
		En string `json:"en"`
		Pt string `json:"pt"`
		Ko string `json:"ko"`
	} `json:"descriptions"`
	BlurbsGeneral struct {
		En string `json:"en"`
		Es string `json:"es"`
		Fr string `json:"fr"`
		Ko string `json:"ko"`
		Pt string `json:"pt"`
		Zh string `json:"zh"`
		Zt string `json:"zt"`
	} `json:"blurbs_general"`
	Genres              []string `json:"genres"`
	SubtitleCompletions struct {
		Cs int `json:"cs"`
		De int `json:"de"`
		En int `json:"en"`
		Es int `json:"es"`
		Fr int `json:"fr"`
		Hu int `json:"hu"`
		Id int `json:"id"`
		It int `json:"it"`
		Mk int `json:"mk"`
		Nl int `json:"nl"`
		Pl int `json:"pl"`
		Pt int `json:"pt"`
		Ro int `json:"ro"`
		Sv int `json:"sv"`
		Tr int `json:"tr"`
		Zt int `json:"zt"`
	} `json:"subtitle_completions"`
	Managers []struct {
		Id       string `json:"id"`
		Username string `json:"username"`
		Images   struct {
			Avatar struct {
				Url interface{} `json:"url"`
			} `json:"avatar"`
		} `json:"images"`
		Url struct {
			Web string `json:"web"`
			Api string `json:"api"`
		} `json:"url"`
	} `json:"managers"`
	Url struct {
		Web string `json:"web"`
		Api string `json:"api"`
		Fb  string `json:"fb"`
	} `json:"url"`
	Flags struct {
		Licensed          bool   `json:"licensed"`
		Hosted            bool   `json:"hosted"`
		State             string `json:"state"`
		Adult             bool   `json:"adult"`
		PrivateWatchParty bool   `json:"private_watch_party"`
		OnAir             bool   `json:"on_air"`
		Exclusive         bool   `json:"exclusive"`
		Original          bool   `json:"original"`
	} `json:"flags"`
	ReviewStats struct {
		AverageRating float64 `json:"average_rating"`
		Count         int     `json:"count"`
	} `json:"review_stats"`
	Movies struct {
		Count int `json:"count"`
		Url   struct {
			Api string `json:"api"`
		} `json:"url"`
	} `json:"movies"`
	Trailers struct {
		Count int `json:"count"`
		Url   struct {
			Api string `json:"api"`
		} `json:"url"`
	} `json:"trailers"`
	TitlesAka struct {
		Zt []string `json:"zt"`
		Zh []string `json:"zh"`
		En []string `json:"en"`
	} `json:"titles_aka"`
	TitlesPhonetic struct {
		En string `json:"en"`
	} `json:"titles_phonetic"`
	Distributors []struct {
		Type string      `json:"type"`
		Name string      `json:"name"`
		From string      `json:"from"`
		To   interface{} `json:"to"`
	} `json:"distributors"`
	WatchNow struct {
		Id   string `json:"id"`
		Type string `json:"type"`
		Url  struct {
			Web string `json:"web"`
			Api string `json:"api"`
		} `json:"url"`
	} `json:"watch_now"`
	Blocked  bool `json:"blocked"`
	Blocking struct {
		Geo      bool `json:"geo"`
		Paywall  bool `json:"paywall"`
		Upcoming bool `json:"upcoming"`
	} `json:"blocking"`
	Paywallable struct {
		Svod bool `json:"svod"`
		Tvod bool `json:"tvod"`
	} `json:"paywallable"`
	Download struct {
		ContentEnabled bool `json:"content_enabled"`
		Enabled        bool `json:"enabled"`
	} `json:"download"`
	Rating string `json:"rating"`
}
type Episodes struct {
	Error    string `json:"error"`
	Vcode    int    `json:"vcode"`
	More     bool   `json:"more"`
	Response []struct {
		Id            string `json:"id"`
		ContentOwners []struct {
			Id string `json:"id"`
		} `json:"content_owners"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Type      string    `json:"type"`
		Duration  int       `json:"duration"`
		Number    int       `json:"number"`
		RootId    string    `json:"root_id"`
		Origin    struct {
			Language string `json:"language"`
		} `json:"origin"`
		Titles struct {
		} `json:"titles"`
		TitlesPhonetic struct {
		} `json:"titles_phonetic"`
		TitlesAka struct {
		} `json:"titles_aka"`
		KalturaId    interface{} `json:"kaltura_id"`
		Descriptions struct {
		} `json:"descriptions"`
		BlurbsGeneral struct {
		} `json:"blurbs_general"`
		SubtitleCompletions struct {
			Ar int `json:"ar,omitempty"`
			Da int `json:"da"`
			De int `json:"de"`
			El int `json:"el"`
			En int `json:"en"`
			Es int `json:"es"`
			Fr int `json:"fr"`
			Hu int `json:"hu"`
			It int `json:"it"`
			Pl int `json:"pl"`
			Pt int `json:"pt"`
			Zh int `json:"zh,omitempty"`
		} `json:"subtitle_completions"`
		Container struct {
			Id      string `json:"id"`
			Type    string `json:"type"`
			Subtype string `json:"subtype"`
			Titles  struct {
				En string `json:"en"`
				Ja string `json:"ja"`
				It string `json:"it"`
				De string `json:"de"`
				Ko string `json:"ko"`
				Pt string `json:"pt"`
				Zt string `json:"zt"`
				Zh string `json:"zh"`
				Fr string `json:"fr"`
				Es string `json:"es"`
			} `json:"titles"`
			TeamName string   `json:"team_name"`
			Genres   []string `json:"genres"`
			Origin   struct {
				Country  string `json:"country"`
				Language string `json:"language"`
			} `json:"origin"`
			Managers []struct {
				Id       string `json:"id"`
				Username string `json:"username"`
				Images   struct {
					Avatar struct {
						Url interface{} `json:"url"`
					} `json:"avatar"`
				} `json:"images"`
				Url struct {
					Web string `json:"web"`
					Api string `json:"api"`
				} `json:"url"`
			} `json:"managers"`
			Images struct {
				Poster struct {
					Url    string `json:"url"`
					Source string `json:"source"`
				} `json:"poster"`
			} `json:"images"`
			Url struct {
				Web string `json:"web"`
				Api string `json:"api"`
			} `json:"url"`
			ReviewStats struct {
				AverageRating float64 `json:"average_rating"`
				Count         int     `json:"count"`
			} `json:"review_stats"`
			PlannedEpisodes int `json:"planned_episodes"`
		} `json:"container"`
		Hardsubs         []interface{} `json:"hardsubs"`
		HardsubLanguages []interface{} `json:"hardsub_languages"`
		Source           string        `json:"source"`
		Images           struct {
			Poster struct {
				Url    string `json:"url"`
				Source string `json:"source"`
			} `json:"poster"`
		} `json:"images"`
		Likes struct {
			Count int `json:"count"`
		} `json:"likes"`
		Flags struct {
			Licensed   bool   `json:"licensed"`
			Hosted     bool   `json:"hosted"`
			OnAir      bool   `json:"on_air"`
			Embeddable bool   `json:"embeddable"`
			State      string `json:"state"`
			Adult      bool   `json:"adult"`
			Hd         bool   `json:"hd"`
			HasStream  bool   `json:"has_stream"`
			Exclusive  bool   `json:"exclusive"`
			Original   bool   `json:"original"`
		} `json:"flags"`
		Url struct {
			Api string `json:"api"`
			Fb  string `json:"fb"`
			Web string `json:"web"`
		} `json:"url"`
		Embed struct {
			Iframe struct {
				Url string `json:"url"`
			} `json:"iframe"`
		} `json:"embed"`
		Parts []struct {
			Id   string `json:"id"`
			Part int    `json:"part"`
			Url  string `json:"url"`
		} `json:"parts"`
		StreamCreatedAt    time.Time   `json:"stream_created_at"`
		KcpStreamCreatedAt interface{} `json:"kcp_stream_created_at"`
		CreditsMarker      int         `json:"credits_marker"`
		PartIndex          int         `json:"part_index"`
		Author             string      `json:"author"`
		AuthorUrl          string      `json:"author_url"`
		VikiAirTime        int         `json:"viki_air_time"`
		Blocked            bool        `json:"blocked"`
		Blocking           struct {
			Geo      bool `json:"geo"`
			Paywall  bool `json:"paywall"`
			Upcoming bool `json:"upcoming"`
		} `json:"blocking"`
		Paywallable struct {
			Svod bool `json:"svod"`
			Tvod bool `json:"tvod"`
		} `json:"paywallable"`
		Download struct {
			ContentEnabled bool `json:"content_enabled"`
			Enabled        bool `json:"enabled"`
		} `json:"download"`
		Rating string `json:"rating"`
	} `json:"response"`
}

func GetId(url string) (movies, id string) {
	re := regexp.MustCompile(`/(\w+)/(\w+)-`)
	matches := re.FindStringSubmatch(url)
	if len(matches) < 3 {
		fmt.Println("No match found")
		return
	}
	movies = matches[1]
	id = matches[2]
	return
}
func GetVideos(id string) (r Videos, err error) {
	url := fmt.Sprintf(GetVideosUrl, id)
	resp, err := client.Do("get", url, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp.Content, &r)
	if r.Error != "" {
		err = fmt.Errorf(r.Error)
	}
	return
}
func GetEpisodes(id string) (r Episodes, err error) {
	url := fmt.Sprintf(GetEpisodesUrl, id)
	resp, err := client.Do("get", url, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp.Content, &r)
	if r.Error != "" {
		err = fmt.Errorf(r.Error)
	}
	return
}
func GetMateInfo(ctx context.Context, sharerUrl string) (r *server.Data, code int) {
	var err error
	requestId := ctx.Value("requestId").(string)
	log.Debugf("site: %s, start GetMateInfo", web)
	r = &server.Data{}
	movie, id := GetId(sharerUrl)
	if movie == "" || id == "" {
		err = fmt.Errorf("invalid URL")
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		code = 2
		return
	}
	videos, err := GetVideos(id)
	if err != nil {
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		code = 1
		return
	}
	r.SeriesTitle = videos.Titles.Zh
	r.SeriesId = videos.Id
	b, _ := json.Marshal(videos.Titles)
	r.Extra = map[string]string{
		"origin": fmt.Sprintf("%s|%s", videos.Origin.Country, videos.Origin.Language),
		"titles": string(b),
	}
	r.IsSeries = false
	if movie != "movies" {
		r.IsSeries = true
		var episodes Episodes
		episodes, err = GetEpisodes(id)
		if err != nil {
			log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
			code = 1
			return
		}
		for _, episode := range episodes.Response {
			r.VideoList = append(r.VideoList, &server.Video{
				EpisodeTitle: episode.Container.Titles.Zh,
				EpisodeId:    episode.Id,
				Episode:      uint32(episode.Number),
			})
		}

	}
	return
}
