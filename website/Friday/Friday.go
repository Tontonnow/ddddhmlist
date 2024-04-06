package Friday

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Tontonnow/ddddhmlist/config"
	"github.com/Tontonnow/ddddhmlist/server"
	"github.com/Tontonnow/ddddhmlist/utils/sesssion"
	"github.com/wangluozhe/requests/url"
	"regexp"
	"strconv"
	"trpc.group/trpc-go/trpc-go/log"
)

var (
	web     = "friday"
	headers = map[string]string{
		"User-Agent":   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.114 Safari/537.36",
		"Platform":     "Android Phone",
		"Version":      "905",
		"Model":        "Pixel",
		"UDID":         config.Conf.AndroidId,
		"Device-Token": "",
	}
	ContentType = map[string]int{
		"movie":  1,
		"drama":  2,
		"anime":  3,
		"tvshow": 4,
		"show":   4,
	}
	re         = regexp.MustCompile(`video.friday.tw/(drama|anime|movie|show|tvshow)/detail/(\d+)`)
	getEpisode = "http://vbmspxy.video.friday.tw/apiv2/content/get?contentId=%d&contentType=%d&eventPageId=01&length=%d&offset=%d&recommendId=-1"
	getToken   = "http://vbmspxy.video.friday.tw/apiv2/token/getv2"
	Headers    = url.ParseHeaders(headers)
	client     = sesssion.NewClient(config.Conf.WebConfig[web])
)

type Episode struct {
	ContentId              int    `json:"contentId"`
	ContentType            int    `json:"contentType"`
	StreamingId            int    `json:"streamingId"`
	StreamingType          int    `json:"streamingType"`
	EpisodeName            string `json:"episodeName"`
	Sort                   string `json:"sort"`
	ChineseName            string `json:"chineseName"`
	EnglishName            string `json:"englishName"`
	Duration               string `json:"duration"`
	StillImageUrl          string `json:"stillImageUrl"`
	SeparationName         string `json:"separationName"`
	SeparationIntroduction string `json:"separationIntroduction"`
	IsPlay                 int    `json:"isPlay"`
	PaymentTagList         []int  `json:"paymentTagList"`
	EnableAd               bool   `json:"enableAd"`
}

type Response struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		AccessToken string `json:"accessToken"`
		Scope       string `json:"scope"`
		ExpiresIn   int    `json:"expiresIn"`
		IdToken     string `json:"idToken"`
		LoginType   int    `json:"LoginType"`
		Content     struct {
			EpisodeList       []Episode `json:"episodeList"`
			ContentId         int       `json:"contentId"`
			ContentType       int       `json:"contentType"`
			Rating            int       `json:"rating"`
			ChineseName       string    `json:"chineseName"`
			EnglishName       string    `json:"englishName"`
			Introduction      string    `json:"introduction"`
			Year              string    `json:"year"`
			Duration          string    `json:"duration"`
			FridayScore       float64   `json:"fridayScore"`
			ScoreCount        int       `json:"scoreCount"`
			ImageUrl          string    `json:"imageUrl"`
			StillImageUrl     string    `json:"stillImageUrl"`
			PropertyTagList   []int     `json:"propertyTagList"`
			BroadcastTime     string    `json:"broadcastTime"`
			Want              int       `json:"want"`
			Area              string    `json:"area"`
			StreamingId       int       `json:"streamingId"`
			StreamingType     int       `json:"streamingType"`
			CanPlay           bool      `json:"canPlay"`
			PaymentTagList    []int     `json:"paymentTagList"`
			IsEmpty           bool      `json:"isEmpty"`
			MirrorEnable      bool      `json:"mirrorEnable"`
			EffectiveDateTime string    `json:"effectiveDateTime"`
			ExpireDate        string    `json:"expireDate"`
			EnableAirPlay     bool      `json:"enableAirPlay"`
			EnableChromecast  bool      `json:"enableChromecast"`
			TwOnlineTime      string    `json:"twOnlineTime"`
			TwBoxOffice       string    `json:"twBoxOffice"`
			GlobalBoxOffice   string    `json:"globalBoxOffice"`
			CategoryList      []struct {
				CategoryId int    `json:"categoryId"`
				Name       string `json:"name"`
			} `json:"categoryList"`
			ContentTagList []struct {
				ContentTag int    `json:"contentTag"`
				Name       string `json:"name"`
			} `json:"contentTagList"`
			ArtistList []struct {
				ChineseName string `json:"chineseName"`
				EnglishName string `json:"englishName,omitempty"`
				ArtistType  int    `json:"artistType"`
			} `json:"artistList"`
			ShortIntroduction   string `json:"shortIntroduction"`
			AwardsList          []int  `json:"awardsList"`
			FinalPlayList       []int  `json:"finalPlayList"`
			CanBook             bool   `json:"canBook"`
			IsBook              bool   `json:"isBook"`
			SrcRecommendId      string `json:"srcRecommendId"`
			RecommendId         string `json:"recommendId"`
			SrcEventPageId      string `json:"srcEventPageId"`
			IsMultiView         bool   `json:"isMultiView"`
			StreamingListScene  []int  `json:"streamingListScene"`
			StreamingListPerson []int  `json:"streamingListPerson"`
			EnableAutoPlay      bool   `json:"enableAutoPlay"`
			EnableAd            bool   `json:"enableAd"`
		}
	} `json:"data"`
}

func SendRequest(method, u string) (rsp Response, err error) {
	Req := url.NewRequest()
	Req.Headers = Headers
	r, err := client.Do(method, u, Req)
	if err != nil {
		return
	}
	err = json.Unmarshal(r.Content, &rsp)
	if err != nil {
		return
	}
	if rsp.Status != 200 {
		err = fmt.Errorf(rsp.Message)
		return
	}
	return
}
func RefreshToken() (err error) {
	rsp, err := SendRequest("POST", getToken)
	if err != nil {
		return
	}
	Headers.Add("Authorization", "Bearer "+rsp.Data.AccessToken)
	log.Debug("Friday RefreshToken Success: ", rsp.Data.AccessToken)
	return
}
func GetMateInfo(ctx context.Context, sharerUrl string) (r *server.Data, code int) {
	var (
		err                                    error
		contentId, contentType, offset, length int
	)
	r = &server.Data{}
	requestId := ctx.Value("requestId").(string)
	log.Debugf("site: %s, start GetMateInfo", web)
	offset = 0
	length = 500
	m := re.FindStringSubmatch(sharerUrl)
	if len(m) != 3 {
		err = fmt.Errorf("URL Error: %s", sharerUrl)
		log.Errorf("site: %s, requestId: %s, error: %v", web, requestId, err)
		code = 2
		return

	}
	contentId, _ = strconv.Atoi(m[2])
	contentType = ContentType[m[1]]
	var EpisodeList []Episode
	for {
		sharerUrl = fmt.Sprintf(getEpisode, contentId, contentType, length, offset)
		d, err := SendRequest("POST", sharerUrl)
		if err != nil {
			log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
			code = 1
			return
		}
		content := d.Data.Content
		r.SeriesTitle = content.ChineseName
		r.SeriesId = strconv.Itoa(content.ContentId)
		if content.EpisodeList != nil {
			EpisodeList = append(EpisodeList, content.EpisodeList...)
			if len(content.EpisodeList) == length {
				offset += length
				continue
			}
			r.IsSeries = true
			for _, v := range EpisodeList {
				episode, _ := strconv.ParseUint(v.Sort, 10, 32)
				r.VideoList = append(r.VideoList, &server.Video{
					EpisodeId:    strconv.Itoa(v.ContentId),
					Episode:      uint32(episode),
					EpisodeTitle: v.ChineseName + " " + v.SeparationName,
					Extra: map[string]string{
						"ContentId":     strconv.Itoa(v.ContentId),
						"ContentType":   strconv.Itoa(v.ContentType),
						"StreamingId":   strconv.Itoa(v.StreamingId),
						"StreamingType": strconv.Itoa(v.StreamingType),
						"EnglishName":   v.EnglishName,
					},
				})

			}
		} else {
			r.IsSeries = false
			r.Extra = map[string]string{
				"ContentId":     strconv.Itoa(content.ContentId),
				"ContentType":   strconv.Itoa(content.ContentType),
				"StreamingId":   strconv.Itoa(content.StreamingId),
				"StreamingType": strconv.Itoa(content.StreamingType),
				"EnglishName":   content.EnglishName,
			}
		}
		break
	}
	return
}
