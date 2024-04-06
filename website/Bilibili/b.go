package Bilibili

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Tontonnow/ddddhmlist/config"
	"github.com/Tontonnow/ddddhmlist/server"
	"github.com/Tontonnow/ddddhmlist/utils/sesssion"
	"regexp"
	"trpc.group/trpc-go/trpc-go/log"
)

var (
	web    = "bilibili"
	api    = "https://api.bilibili.com/pgc/view/web/season?ep_id=%s&season_id=%s" //https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/docs/bangumi/info.md
	client = sesssion.NewClient(config.Conf.WebConfig[web])
	re     = regexp.MustCompile(`/(ss|ep)(\d+)`)
)

type Season struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  struct {
		Activity struct {
			HeadBgUrl string `json:"head_bg_url"`
			Id        int    `json:"id"`
			Title     string `json:"title"`
		} `json:"activity"`
		Actors string `json:"actors"`
		Alias  string `json:"alias"`
		Areas  []struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"areas"`
		BkgCover              string `json:"bkg_cover"`
		Cover                 string `json:"cover"`
		DeliveryFragmentVideo bool   `json:"delivery_fragment_video"`
		EnableVt              bool   `json:"enable_vt"`
		Episodes              []struct {
			Aid       int    `json:"aid"`
			Badge     string `json:"badge"`
			BadgeInfo struct {
				BgColor      string `json:"bg_color"`
				BgColorNight string `json:"bg_color_night"`
				Text         string `json:"text"`
			} `json:"badge_info"`
			BadgeType int    `json:"badge_type"`
			Bvid      string `json:"bvid"`
			Cid       int    `json:"cid"`
			Cover     string `json:"cover"`
			Dimension struct {
				Height int `json:"height"`
				Rotate int `json:"rotate"`
				Width  int `json:"width"`
			} `json:"dimension"`
			Duration    int    `json:"duration"`
			EnableVt    bool   `json:"enable_vt"`
			EpId        int    `json:"ep_id"`
			From        string `json:"from"`
			Id          int    `json:"id"`
			IsViewHide  bool   `json:"is_view_hide"`
			Link        string `json:"link"`
			LongTitle   string `json:"long_title"`
			PubTime     int    `json:"pub_time"`
			Pv          int    `json:"pv"`
			ReleaseDate string `json:"release_date"`
			Rights      struct {
				AllowDemand   int `json:"allow_demand"`
				AllowDm       int `json:"allow_dm"`
				AllowDownload int `json:"allow_download"`
				AreaLimit     int `json:"area_limit"`
			} `json:"rights"`
			ShareCopy          string `json:"share_copy"`
			ShareUrl           string `json:"share_url"`
			ShortLink          string `json:"short_link"`
			ShowDrmLoginDialog bool   `json:"showDrmLoginDialog"`
			Status             int    `json:"status"`
			Subtitle           string `json:"subtitle"`
			Title              string `json:"title"`
			Vid                string `json:"vid"`
		} `json:"episodes"`
		Evaluate string `json:"evaluate"`
		Freya    struct {
			BubbleDesc    string `json:"bubble_desc"`
			BubbleShowCnt int    `json:"bubble_show_cnt"`
			IconShow      int    `json:"icon_show"`
		} `json:"freya"`
		HideEpVvVtDm int `json:"hide_ep_vv_vt_dm"`
		IconFont     struct {
			Name string `json:"name"`
			Text string `json:"text"`
		} `json:"icon_font"`
		JpTitle string `json:"jp_title"`
		Link    string `json:"link"`
		MediaId int    `json:"media_id"`
		Mode    int    `json:"mode"`
		NewEp   struct {
			Desc  string `json:"desc"`
			Id    int    `json:"id"`
			IsNew int    `json:"is_new"`
			Title string `json:"title"`
		} `json:"new_ep"`
		Payment struct {
			Discount int `json:"discount"`
			PayType  struct {
				AllowDiscount    int `json:"allow_discount"`
				AllowPack        int `json:"allow_pack"`
				AllowTicket      int `json:"allow_ticket"`
				AllowTimeLimit   int `json:"allow_time_limit"`
				AllowVipDiscount int `json:"allow_vip_discount"`
				ForbidBb         int `json:"forbid_bb"`
			} `json:"pay_type"`
			Price             string `json:"price"`
			Promotion         string `json:"promotion"`
			Tip               string `json:"tip"`
			ViewStartTime     int    `json:"view_start_time"`
			VipDiscount       int    `json:"vip_discount"`
			VipFirstPromotion string `json:"vip_first_promotion"`
			VipPrice          string `json:"vip_price"`
			VipPromotion      string `json:"vip_promotion"`
		} `json:"payment"`
		Positive struct {
			Id    int    `json:"id"`
			Title string `json:"title"`
		} `json:"positive"`
		Publish struct {
			IsFinish      int    `json:"is_finish"`
			IsStarted     int    `json:"is_started"`
			PubTime       string `json:"pub_time"`
			PubTimeShow   string `json:"pub_time_show"`
			UnknowPubDate int    `json:"unknow_pub_date"`
			Weekday       int    `json:"weekday"`
		} `json:"publish"`
		Rating struct {
			Count int     `json:"count"`
			Score float64 `json:"score"`
		} `json:"rating"`
		Record string `json:"record"`
		Rights struct {
			AllowBp         int    `json:"allow_bp"`
			AllowBpRank     int    `json:"allow_bp_rank"`
			AllowDownload   int    `json:"allow_download"`
			AllowReview     int    `json:"allow_review"`
			AreaLimit       int    `json:"area_limit"`
			BanAreaShow     int    `json:"ban_area_show"`
			CanWatch        int    `json:"can_watch"`
			Copyright       string `json:"copyright"`
			ForbidPre       int    `json:"forbid_pre"`
			FreyaWhite      int    `json:"freya_white"`
			IsCoverShow     int    `json:"is_cover_show"`
			IsPreview       int    `json:"is_preview"`
			OnlyVipDownload int    `json:"only_vip_download"`
			Resource        string `json:"resource"`
			WatchPlatform   int    `json:"watch_platform"`
		} `json:"rights"`
		SeasonId    int           `json:"season_id"`
		SeasonTitle string        `json:"season_title"`
		Seasons     []interface{} `json:"seasons"`
		Section     []struct {
			Attr       int           `json:"attr"`
			EpisodeId  int           `json:"episode_id"`
			EpisodeIds []interface{} `json:"episode_ids,omitempty"`
			Episodes   []struct {
				Aid       int    `json:"aid"`
				Badge     string `json:"badge"`
				BadgeInfo struct {
					BgColor      string `json:"bg_color"`
					BgColorNight string `json:"bg_color_night"`
					Text         string `json:"text"`
				} `json:"badge_info"`
				BadgeType int    `json:"badge_type,omitempty"`
				Bvid      string `json:"bvid,omitempty"`
				Cid       int    `json:"cid"`
				Cover     string `json:"cover"`
				Dimension struct {
					Height int `json:"height"`
					Rotate int `json:"rotate"`
					Width  int `json:"width"`
				} `json:"dimension,omitempty"`
				Duration int    `json:"duration,omitempty"`
				EnableVt bool   `json:"enable_vt"`
				EpId     int    `json:"ep_id"`
				From     string `json:"from,omitempty"`
				IconFont struct {
					Name string `json:"name"`
					Text string `json:"text"`
				} `json:"icon_font"`
				Id          int    `json:"id"`
				IsViewHide  bool   `json:"is_view_hide"`
				Link        string `json:"link"`
				LongTitle   string `json:"long_title,omitempty"`
				PubTime     int    `json:"pub_time"`
				Pv          int    `json:"pv"`
				ReleaseDate string `json:"release_date,omitempty"`
				Rights      struct {
					AllowDemand   int `json:"allow_demand"`
					AllowDm       int `json:"allow_dm"`
					AllowDownload int `json:"allow_download"`
					AreaLimit     int `json:"area_limit"`
				} `json:"rights,omitempty"`
				ShareCopy          string `json:"share_copy,omitempty"`
				ShareUrl           string `json:"share_url,omitempty"`
				ShortLink          string `json:"short_link,omitempty"`
				ShowDrmLoginDialog bool   `json:"showDrmLoginDialog"`
				Stat               struct {
					Coin     int `json:"coin"`
					Danmakus int `json:"danmakus"`
					Likes    int `json:"likes"`
					Play     int `json:"play"`
					Reply    int `json:"reply"`
					Vt       int `json:"vt"`
				} `json:"stat"`
				StatForUnity struct {
					Coin    int `json:"coin"`
					Danmaku struct {
						Icon     string `json:"icon"`
						PureText string `json:"pure_text"`
						Text     string `json:"text"`
						Value    int    `json:"value"`
					} `json:"danmaku"`
					Likes int `json:"likes"`
					Reply int `json:"reply"`
					Vt    struct {
						Icon     string `json:"icon"`
						PureText string `json:"pure_text"`
						Text     string `json:"text"`
						Value    int    `json:"value"`
					} `json:"vt"`
				} `json:"stat_for_unity,omitempty"`
				Status      int    `json:"status"`
				Subtitle    string `json:"subtitle,omitempty"`
				Title       string `json:"title"`
				Vid         string `json:"vid,omitempty"`
				ArchiveAttr int    `json:"archive_attr,omitempty"`
				LinkType    string `json:"link_type,omitempty"`
				Report      struct {
					Aid         string `json:"aid"`
					EpTitle     string `json:"ep_title"`
					Position    string `json:"position"`
					SeasonId    string `json:"season_id"`
					SeasonType  string `json:"season_type"`
					SectionId   string `json:"section_id"`
					SectionType string `json:"section_type"`
				} `json:"report,omitempty"`
			} `json:"episodes"`
			Id     int    `json:"id"`
			Title  string `json:"title"`
			Type   int    `json:"type"`
			Type2  int    `json:"type2"`
			Report struct {
				SeasonId    string `json:"season_id"`
				SeasonType  string `json:"season_type"`
				SecTitle    string `json:"sec_title"`
				SectionId   string `json:"section_id"`
				SectionType string `json:"section_type"`
			} `json:"report,omitempty"`
		} `json:"section"`
		Series struct {
			DisplayType int    `json:"display_type"`
			SeriesId    int    `json:"series_id"`
			SeriesTitle string `json:"series_title"`
		} `json:"series"`
		ShareCopy     string `json:"share_copy"`
		ShareSubTitle string `json:"share_sub_title"`
		ShareUrl      string `json:"share_url"`
		Show          struct {
			WideScreen int `json:"wide_screen"`
		} `json:"show"`
		ShowSeasonType int    `json:"show_season_type"`
		SquareCover    string `json:"square_cover"`
		Staff          string `json:"staff"`
		Stat           struct {
			Coins      int    `json:"coins"`
			Danmakus   int    `json:"danmakus"`
			Favorite   int    `json:"favorite"`
			Favorites  int    `json:"favorites"`
			FollowText string `json:"follow_text"`
			Likes      int    `json:"likes"`
			Reply      int    `json:"reply"`
			Share      int    `json:"share"`
			Views      int    `json:"views"`
			Vt         int    `json:"vt"`
		} `json:"stat"`
		Status   int      `json:"status"`
		Styles   []string `json:"styles"`
		Subtitle string   `json:"subtitle"`
		Title    string   `json:"title"`
		Total    int      `json:"total"`
		Type     int      `json:"type"`
		UpInfo   struct {
			Avatar             string `json:"avatar"`
			AvatarSubscriptUrl string `json:"avatar_subscript_url"`
			Follower           int    `json:"follower"`
			IsFollow           int    `json:"is_follow"`
			Mid                int    `json:"mid"`
			NicknameColor      string `json:"nickname_color"`
			Pendant            struct {
				Image string `json:"image"`
				Name  string `json:"name"`
				Pid   int    `json:"pid"`
			} `json:"pendant"`
			ThemeType  int    `json:"theme_type"`
			Uname      string `json:"uname"`
			VerifyType int    `json:"verify_type"`
			VipLabel   struct {
				BgColor     string `json:"bg_color"`
				BgStyle     int    `json:"bg_style"`
				BorderColor string `json:"border_color"`
				Text        string `json:"text"`
				TextColor   string `json:"text_color"`
			} `json:"vip_label"`
			VipStatus int `json:"vip_status"`
			VipType   int `json:"vip_type"`
		} `json:"up_info"`
		UserStatus struct {
			AreaLimit    int `json:"area_limit"`
			BanAreaShow  int `json:"ban_area_show"`
			Follow       int `json:"follow"`
			FollowStatus int `json:"follow_status"`
			Login        int `json:"login"`
			Pay          int `json:"pay"`
			PayPackPaid  int `json:"pay_pack_paid"`
			Sponsor      int `json:"sponsor"`
		} `json:"user_status"`
	} `json:"result"`
}

func GetSeason(url string) (s Season, err error) {
	rsp, err := client.Do("GET", url, nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(rsp.Content, &s)
	if err != nil {
		return
	}
	if s.Code != 0 {
		err = fmt.Errorf(s.Message)
		return
	}
	return
}
func GetMateInfo(ctx context.Context, sharerUrl string) (r *server.Data, code int) {
	var (
		err      error
		epId     = "0"
		seasonId = "0"
	)
	r = &server.Data{}
	requestId := ctx.Value("requestId").(string)
	log.Debugf("site: %s, start GetMateInfo", web)
	m := re.FindStringSubmatch(sharerUrl)
	if len(m) != 3 {
		err = fmt.Errorf("URL Error: %s", sharerUrl)
		log.Errorf("site: %s, requestId: %s, error: %v", web, requestId, err)
		code = 2
		return
	}
	if m[1] == "ss" {
		seasonId = m[2]
	} else {
		epId = m[2]
	}
	url := fmt.Sprintf(api, epId, seasonId)
	s, err := GetSeason(url)
	if err != nil {
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, url)
		code = 1
		return
	}
	r.SeriesTitle = s.Result.SeasonTitle
	r.SeriesId = fmt.Sprintf("%d", s.Result.SeasonId)
	r.IsSeries = s.Result.Type != 2
	/*
		for _, v := range s.Result.Section {//预告之类的
			s.Result.Episodes = append(s.Result.Episodes, v.Episodes...)
		}
	*/
	for _, v := range s.Result.Episodes { //正片
		r.VideoList = append(r.VideoList, &server.Video{
			EpisodeTitle: v.Title,
			EpisodeId:    fmt.Sprintf("%d", v.Cid),
			Extra: map[string]string{
				"Title:": v.ShareCopy,
				"ep_id":  fmt.Sprintf("%d", v.EpId),
				"badge":  v.Badge,
			},
		})
	}
	return

}
