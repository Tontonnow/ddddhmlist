package XiGUA

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Tontonnow/ddddhmlist/config"
	"github.com/Tontonnow/ddddhmlist/server"
	"github.com/Tontonnow/ddddhmlist/utils/sesssion"
	"github.com/wangluozhe/requests/url"
	"regexp"
	"trpc.group/trpc-go/trpc-go/log"
)

var (
	web    = "ixigua"
	api    = "https://ib.snssdk.com/vapp/lvideo/api/info/"
	upapi  = "https://ib.snssdk.com/video/app/article/full/0/0/%s/0/0/0/"
	client = sesssion.NewClient(config.Conf.WebConfig[web])
)

type VideoInfo struct {
	BaseResp struct {
		StatusCode    int    `json:"status_code"`
		StatusMessage string `json:"status_message"`
	} `json:"base_resp"`
	Album struct {
		AlbumId       string `json:"album_id"`
		Title         string `json:"title"`
		TotalEpisodes int    `json:"total_episodes"`
		LatestSeq     int    `json:"latest_seq"`
		Attribute     string `json:"attribute"`
		BottomLabel   string `json:"bottom_label"`
		CoverList     []struct {
			Url           string   `json:"url"`
			Uri           string   `json:"uri"`
			Width         string   `json:"width"`
			Height        string   `json:"height"`
			UrlList       []string `json:"url_list"`
			ImageStyle    int      `json:"image_style"`
			LargeUrlList  []string `json:"large_url_list"`
			MediumUrlList []string `json:"medium_url_list"`
			ThumbUrlList  []string `json:"thumb_url_list"`
		} `json:"cover_list"`
		AlbumTypeList []int    `json:"album_type_list"`
		SeqType       int      `json:"seq_type"`
		Year          string   `json:"year"`
		Intro         string   `json:"intro"`
		TagList       []string `json:"tag_list"`
		AreaList      []string `json:"area_list"`
		ActorList     []struct {
			CelebrityId   string `json:"celebrity_id"`
			Name          string `json:"name"`
			Summary       string `json:"summary"`
			RoleName      string `json:"role_name,omitempty"`
			Rank          int    `json:"rank"`
			CelebrityType int    `json:"celebrity_type"`
		} `json:"actor_list"`
		DirectorList []struct {
			CelebrityId   string `json:"celebrity_id"`
			Name          string `json:"name"`
			Summary       string `json:"summary,omitempty"`
			Rank          int    `json:"rank"`
			CelebrityType int    `json:"celebrity_type"`
		} `json:"director_list"`
		ShareUrl      string `json:"share_url"`
		LogPb         string `json:"log_pb"`
		GroupSource   int    `json:"group_source"`
		SubTitle      string `json:"sub_title"`
		ReleaseStatus int    `json:"release_status"`
		SearchTagList []struct {
			Name      string `json:"name"`
			SearchKey string `json:"search_key"`
		} `json:"search_tag_list"`
		CelebrityList []struct {
			CelebrityId   string `json:"celebrity_id"`
			Name          string `json:"name"`
			Summary       string `json:"summary,omitempty"`
			RoleName      string `json:"role_name,omitempty"`
			Rank          int    `json:"rank"`
			CelebrityType int    `json:"celebrity_type"`
		} `json:"celebrity_list"`
		UserInfo struct {
			SecUserId string `json:"sec_user_id"`
		} `json:"user_info"`
		Duration string `json:"duration"`
		Label    struct {
			Text      string `json:"text"`
			TextColor string `json:"text_color"`
			BgColor   string `json:"bg_color"`
			LabelType int    `json:"label_type"`
		} `json:"label"`
		TipList []struct {
			TipId      string `json:"tip_id"`
			Type       int    `json:"type"`
			Style      int    `json:"style"`
			ButtonList []struct {
				ContentInfo struct {
					Content string `json:"content"`
				} `json:"content_info"`
				TagName string `json:"tag_name"`
			} `json:"button_list"`
		} `json:"tip_list"`
		ActorInfoList []struct {
			ActorName string `json:"actor_name"`
		} `json:"actor_info_list"`
		TagInfoList []struct {
			TagName string `json:"tag_name"`
		} `json:"tag_info_list"`
		AlbumGroupId         string `json:"album_group_id"`
		TvIncentiveAdControl string `json:"tv_incentive_ad_control"`
	} `json:"album"`
	BlockList []struct {
		Name       string `json:"name"`
		Title      string `json:"title"`
		Type       int    `json:"type"`
		ActionList []struct {
			Name     string `json:"name"`
			Text     string `json:"text"`
			Position int    `json:"position"`
		} `json:"action_list"`
		LogPb string `json:"log_pb"`
		Intro struct {
			Title         string   `json:"title"`
			TotalEpisodes int      `json:"total_episodes"`
			LatestSeq     int      `json:"latest_seq"`
			AlbumTypeList []int    `json:"album_type_list"`
			SeqType       int      `json:"seq_type"`
			Year          string   `json:"year"`
			Intro         string   `json:"intro"`
			TagList       []string `json:"tag_list"`
			AreaList      []string `json:"area_list"`
			ReleaseStatus int      `json:"release_status"`
			Label         struct {
				Text      string `json:"text"`
				TextColor string `json:"text_color"`
				BgColor   string `json:"bg_color"`
				LabelType int    `json:"label_type"`
			} `json:"label"`
		} `json:"intro,omitempty"`
		Id    string `json:"id,omitempty"`
		Style int    `json:"style,omitempty"`
		Cells []struct {
			CellType int    `json:"cell_type"`
			Offset   string `json:"offset"`
			Episode  struct {
				EpisodeId string `json:"episode_id"`
				AlbumId   string `json:"album_id"`
				Rank      int    `json:"rank"`
				SeqOld    string `json:"seq_old"`
				Title     string `json:"title"`
				Name      string `json:"name"`
				VideoInfo struct {
					Width                string  `json:"width"`
					Height               string  `json:"height"`
					Vid                  string  `json:"vid"`
					AuthToken            string  `json:"auth_token,omitempty"`
					Duration             float64 `json:"duration"`
					EncodedVideoInfoList []struct {
						Definition string  `json:"Definition"`
						Size       string  `json:"Size"`
						Duration   float64 `json:"Duration"`
						Height     string  `json:"Height"`
						Width      string  `json:"Width"`
					} `json:"encoded_video_info_list"`
					BusinessToken  string `json:"business_token,omitempty"`
					PlayAuthToken  string `json:"play_auth_token,omitempty"`
					VideoModelJson string `json:"video_model_json,omitempty"`
				} `json:"video_info"`
				CoverList []struct {
					Url           string   `json:"url"`
					Uri           string   `json:"uri"`
					Width         string   `json:"width"`
					Height        string   `json:"height"`
					UrlList       []string `json:"url_list"`
					ImageStyle    int      `json:"image_style"`
					LargeUrlList  []string `json:"large_url_list"`
					MediumUrlList []string `json:"medium_url_list"`
					ThumbUrlList  []string `json:"thumb_url_list"`
				} `json:"cover_list"`
				ShareUrl           string `json:"share_url"`
				DiggCount          string `json:"digg_count"`
				GroupSource        int    `json:"group_source"`
				LogPb              string `json:"log_pb"`
				Attribute          string `json:"attribute"`
				EpisodeType        int    `json:"episode_type"`
				BottomLabel        string `json:"bottom_label"`
				InteractionControl string `json:"interaction_control"`
				Seq                int    `json:"seq"`
				SeqType            int    `json:"seq_type"`
				TipList            []struct {
					TipId      string `json:"tip_id"`
					Type       int    `json:"type"`
					Style      int    `json:"style"`
					ButtonList []struct {
						ContentInfo struct {
							Content string `json:"content"`
							Schema  string `json:"schema,omitempty"`
						} `json:"content_info"`
						TagName string `json:"tag_name"`
					} `json:"button_list,omitempty"`
					Description struct {
						Content string `json:"content"`
					} `json:"description,omitempty"`
				} `json:"tip_list"`
				VipPlayControl int `json:"vip_play_control,omitempty"`
				UserInfo       struct {
					SecUserId string `json:"sec_user_id"`
				} `json:"user_info"`
				ImmersionInfo struct {
				} `json:"immersion_info"`
				TvIncentiveAdControl string   `json:"tv_incentive_ad_control,omitempty"`
				TvBreakTitle         []string `json:"tv_break_title"`
				Extra                string   `json:"extra"`
				Label                struct {
					Text      string `json:"text"`
					TextColor string `json:"text_color"`
					BgColor   string `json:"bg_color"`
					LabelType int    `json:"label_type"`
				} `json:"label,omitempty"`
				TvDisableRecovery bool `json:"tv_disable_recovery,omitempty"`
			} `json:"episode,omitempty"`
			CellSize int `json:"cell_size,omitempty"`
			Album    struct {
				AlbumId       string `json:"album_id"`
				Title         string `json:"title"`
				TotalEpisodes int    `json:"total_episodes"`
				LatestSeq     int    `json:"latest_seq"`
				BottomLabel   string `json:"bottom_label"`
				CoverList     []struct {
					Url           string   `json:"url"`
					Uri           string   `json:"uri"`
					Width         string   `json:"width"`
					Height        string   `json:"height"`
					UrlList       []string `json:"url_list"`
					ImageStyle    int      `json:"image_style"`
					LargeUrlList  []string `json:"large_url_list"`
					MediumUrlList []string `json:"medium_url_list"`
					ThumbUrlList  []string `json:"thumb_url_list"`
				} `json:"cover_list"`
				AlbumTypeList []int    `json:"album_type_list"`
				SeqType       int      `json:"seq_type"`
				Year          string   `json:"year"`
				Intro         string   `json:"intro"`
				TagList       []string `json:"tag_list"`
				AreaList      []string `json:"area_list"`
				ShareUrl      string   `json:"share_url"`
				LogPb         string   `json:"log_pb"`
				GroupSource   int      `json:"group_source"`
				SubTitle      string   `json:"sub_title"`
				ReleaseStatus int      `json:"release_status"`
				SearchTagList []struct {
					Name      string `json:"name"`
					SearchKey string `json:"search_key"`
				} `json:"search_tag_list"`
				UserInfo struct {
					SecUserId string `json:"sec_user_id"`
				} `json:"user_info"`
				Duration string `json:"duration"`
				Label    struct {
					Text      string `json:"text"`
					TextColor string `json:"text_color"`
					BgColor   string `json:"bg_color"`
					LabelType int    `json:"label_type,omitempty"`
				} `json:"label"`
				ActorInfoList []struct {
					ActorName string `json:"actor_name"`
				} `json:"actor_info_list"`
				TagInfoList []struct {
					TagName string `json:"tag_name"`
				} `json:"tag_info_list"`
				AlbumGroupId  string `json:"album_group_id"`
				ImmersionInfo struct {
				} `json:"immersion_info"`
				Attribute string `json:"attribute,omitempty"`
				TipList   []struct {
					TipId      string `json:"tip_id"`
					Type       int    `json:"type"`
					Style      int    `json:"style"`
					ButtonList []struct {
						ContentInfo struct {
							Content string `json:"content"`
						} `json:"content_info"`
						TagName string `json:"tag_name"`
					} `json:"button_list"`
				} `json:"tip_list,omitempty"`
				TvIncentiveAdControl string `json:"tv_incentive_ad_control,omitempty"`
				SubscribeStatus      int    `json:"SubscribeStatus,omitempty"`
				SubscribeOnlineTime  string `json:"SubscribeOnlineTime,omitempty"`
				InteractionControl   string `json:"interaction_control,omitempty"`
			} `json:"album,omitempty"`
		} `json:"cells,omitempty"`
	} `json:"block_list"`
	Episode struct {
		EpisodeId string `json:"episode_id"`
		AlbumId   string `json:"album_id"`
		Rank      int    `json:"rank"`
		SeqOld    string `json:"seq_old"`
		Title     string `json:"title"`
		Name      string `json:"name"`
		VideoInfo struct {
			Width                string  `json:"width"`
			Height               string  `json:"height"`
			Vid                  string  `json:"vid"`
			AuthToken            string  `json:"auth_token"`
			Duration             float64 `json:"duration"`
			EncodedVideoInfoList []struct {
				Definition string  `json:"Definition"`
				Size       string  `json:"Size"`
				Duration   float64 `json:"Duration"`
				Height     string  `json:"Height"`
				Width      string  `json:"Width"`
			} `json:"encoded_video_info_list"`
			BusinessToken  string `json:"business_token"`
			PlayAuthToken  string `json:"play_auth_token"`
			VideoModelJson string `json:"video_model_json"`
		} `json:"video_info"`
		CoverList []struct {
			Url           string   `json:"url"`
			Uri           string   `json:"uri"`
			Width         string   `json:"width"`
			Height        string   `json:"height"`
			UrlList       []string `json:"url_list"`
			ImageStyle    int      `json:"image_style"`
			LargeUrlList  []string `json:"large_url_list"`
			MediumUrlList []string `json:"medium_url_list"`
			ThumbUrlList  []string `json:"thumb_url_list"`
		} `json:"cover_list"`
		ShareUrl           string `json:"share_url"`
		DiggCount          string `json:"digg_count"`
		GroupSource        int    `json:"group_source"`
		LogPb              string `json:"log_pb"`
		Attribute          string `json:"attribute"`
		EpisodeType        int    `json:"episode_type"`
		BottomLabel        string `json:"bottom_label"`
		InteractionControl string `json:"interaction_control"`
		Seq                int    `json:"seq"`
		SeqType            int    `json:"seq_type"`
		TipList            []struct {
			TipId      string `json:"tip_id"`
			Type       int    `json:"type"`
			Style      int    `json:"style"`
			ButtonList []struct {
				ContentInfo struct {
					Content string `json:"content"`
				} `json:"content_info"`
				TagName string `json:"tag_name"`
			} `json:"button_list"`
		} `json:"tip_list"`
		VipPlayControl int `json:"vip_play_control"`
		UserInfo       struct {
			SecUserId string `json:"sec_user_id"`
		} `json:"user_info"`
		TvIncentiveAdControl string   `json:"tv_incentive_ad_control"`
		TvBreakTitle         []string `json:"tv_break_title"`
		Extra                string   `json:"extra"`
	} `json:"episode"`
	CommentPrompt string `json:"comment_prompt"`
	InfoControl   struct {
	} `json:"info_control"`
	PlayerOperationSet struct {
		Loading struct {
			EffectiveRange int `json:"effective_range"`
		} `json:"loading"`
		ProgressBar struct {
			EffectiveRange int `json:"effective_range"`
		} `json:"progress_bar"`
	} `json:"player_operation_set"`
}
type ArticleInfo struct {
	Data struct {
		Delete                int         `json:"delete"`
		DeleteReason          string      `json:"delete_reason"`
		UserId                int         `json:"user_id"`
		Abstract              string      `json:"abstract"`
		ActivityLabel         interface{} `json:"activity_label"`
		AggrType              int         `json:"aggr_type"`
		AnchorInfo            interface{} `json:"anchor_info"`
		ArticleSubType        int         `json:"article_sub_type"`
		ArticleType           int         `json:"article_type"`
		ArticleUrl            string      `json:"article_url"`
		BanAudioComment       bool        `json:"ban_audio_comment"`
		BanAudioCommentReason string      `json:"ban_audio_comment_reason"`
		BanComment            int         `json:"ban_comment"`
		BanDanmaku            int         `json:"ban_danmaku"`
		BanDanmakuReason      string      `json:"ban_danmaku_reason"`
		BanDanmakuSend        int         `json:"ban_danmaku_send"`
		BanDownload           int         `json:"ban_download"`
		BanDownloadReason     string      `json:"ban_download_reason"`
		BanImmersive          int         `json:"ban_immersive"`
		BanShare              int         `json:"ban_share"`
		BuryCount             int         `json:"bury_count"`
		CanCommentLevel       int         `json:"can_comment_level"`
		Categories            []string    `json:"categories"`
		CellType              int         `json:"cell_type"`
		CommentCount          int         `json:"comment_count"`
		Composition           int         `json:"composition"`
		Content               string      `json:"content"`
		CoverMainColor        string      `json:"cover_main_color"`
		DanmakuCount          int         `json:"danmaku_count"`
		DefaultDanmaku        int         `json:"default_danmaku"`
		DetailSchema          string      `json:"detail_schema"`
		DiggCount             int         `json:"digg_count"`
		DisplayUrl            string      `json:"display_url"`
		DxUpgradedVideo       bool        `json:"dx_upgraded_video"`
		ExtensionsAdRawData   interface{} `json:"extensions_ad_raw_data"`
		FeedRevisitSatisfied  int         `json:"feed_revisit_satisfied"`
		FilterWords           []struct {
			Id         string `json:"id"`
			IsSelected bool   `json:"is_selected"`
			Name       string `json:"name"`
		} `json:"filter_words"`
		FirstFrameImage struct {
			Height  int    `json:"height"`
			Uri     string `json:"uri"`
			Url     string `json:"url"`
			UrlList []struct {
				Url string `json:"url"`
			} `json:"url_list"`
			Width int `json:"width"`
		} `json:"first_frame_image"`
		Gid             string `json:"gid"`
		GroupFlags      int    `json:"group_flags"`
		GroupId         int64  `json:"group_id"`
		GroupSource     int    `json:"group_source"`
		HasVideo        bool   `json:"has_video"`
		HideType        int    `json:"hide_type"`
		HistoryDuration int    `json:"history_duration"`
		IesItemExtra    struct {
			AwemeUid     int         `json:"aweme_uid"`
			GroupIdList0 interface{} `json:"group_id_list_0"`
			GroupIdList1 interface{} `json:"group_id_list_1"`
			SimIdList0   interface{} `json:"sim_id_list_0"`
			SimIdList1   interface{} `json:"sim_id_list_1"`
		} `json:"ies_item_extra"`
		ImpressionCount int    `json:"impression_count"`
		InsertAds       string `json:"insert_ads"`
		IpInfo          struct {
			Address string `json:"address"`
		} `json:"ip_info"`
		IsKeyVideo     bool  `json:"is_key_video"`
		IsOriginal     bool  `json:"is_original"`
		IsSubscribe    bool  `json:"is_subscribe"`
		ItemId         int64 `json:"item_id"`
		LargeImageList []struct {
			Height  int    `json:"height"`
			Uri     string `json:"uri"`
			Url     string `json:"url"`
			UrlList []struct {
				Url string `json:"url"`
			} `json:"url_list"`
			Width int `json:"width"`
		} `json:"large_image_list"`
		LogPb struct {
			AuthorId        string `json:"author_id"`
			AwemeAuthorId   string `json:"aweme_author_id"`
			AwemeItemId     string `json:"aweme_item_id"`
			GroupId         string `json:"group_id"`
			GroupSource     string `json:"group_source"`
			GroupType       string `json:"group_type"`
			ImprId          string `json:"impr_id"`
			IsFollowing     string `json:"is_following"`
			IsFromAwemeSort string `json:"is_from_aweme_sort"`
			IsRevisit       string `json:"is_revisit"`
			XgAuthorId      string `json:"xg_author_id"`
			XgItemId        string `json:"xg_item_id"`
		} `json:"log_pb"`
		LvideoIpAggregationBriefInfos interface{} `json:"lvideo_ip_aggregation_brief_infos"`
		MediaInfo                     struct {
			AvatarUrl    string      `json:"avatar_url"`
			MediaId      int64       `json:"media_id"`
			Name         string      `json:"name"`
			Subcribed    int         `json:"subcribed"`
			Subscribed   int         `json:"subscribed"`
			UserId       interface{} `json:"user_id"`
			UserVerified bool        `json:"user_verified"`
		} `json:"media_info"`
		MediaName   string `json:"media_name"`
		MiddleImage struct {
			Height  int    `json:"height"`
			Uri     string `json:"uri"`
			Url     string `json:"url"`
			UrlList []struct {
				Url string `json:"url"`
			} `json:"url_list"`
			Width int `json:"width"`
		} `json:"middle_image"`
		NearId              int64  `json:"near_id"`
		NearId2             int64  `json:"near_id2"`
		NearId3             int    `json:"near_id3"`
		PlayAuthToken       string `json:"play_auth_token"`
		PlayBizToken        string `json:"play_biz_token"`
		PreadParams         string `json:"pread_params"`
		PublishTime         int    `json:"publish_time"`
		Rank                int    `json:"rank"`
		RepinCount          int    `json:"repin_count"`
		ShareCount          int    `json:"share_count"`
		ShareToken          string `json:"share_token"`
		ShareUrl            string `json:"share_url"`
		ShowPortrait        bool   `json:"show_portrait"`
		ShowPortraitArticle bool   `json:"show_portrait_article"`
		Source              string `json:"source"`
		SuitableListening   bool   `json:"suitable_listening"`
		SuperDiggControl    struct {
			AnimeBodyUrl string `json:"anime_body_url"`
			AnimeHeadUrl string `json:"anime_head_url"`
			AnimeKey     string `json:"anime_key"`
			AnimeLokiId  string `json:"anime_loki_id"`
			Audio        struct {
				AudioName string `json:"audio_name"`
				AudioType string `json:"audio_type"`
				Uri       string `json:"uri"`
				Url       string `json:"url"`
			} `json:"audio"`
		} `json:"super_digg_control"`
		Tag      string `json:"tag"`
		Title    string `json:"title"`
		UserInfo struct {
			AuthorDesc        string      `json:"author_desc"`
			AvatarUrl         string      `json:"avatar_url"`
			Description       string      `json:"description"`
			Follow            bool        `json:"follow"`
			FollowerCount     int         `json:"follower_count"`
			FollowersCountStr string      `json:"followers_count_str"`
			IsBlocked         bool        `json:"is_blocked"`
			IsBlocking        bool        `json:"is_blocking"`
			IsDiscipulus      bool        `json:"is_discipulus"`
			IsFollower        bool        `json:"is_follower"`
			IsLiving          bool        `json:"is_living"`
			IsTvTopConsume    bool        `json:"is_tv_top_consume"`
			Name              string      `json:"name"`
			SecUserId         interface{} `json:"sec_user_id"`
			UserAuthInfo      string      `json:"user_auth_info"`
			UserId            int64       `json:"user_id"`
			UserVerified      bool        `json:"user_verified"`
			VerifiedContent   string      `json:"verified_content"`
			VideoTotalCount   int         `json:"video_total_count"`
		} `json:"user_info"`
		UserVerified    int    `json:"user_verified"`
		VerifiedContent string `json:"verified_content"`
		VerifyReason    string `json:"verify_reason"`
		VerifyStatus    int    `json:"verify_status"`
		VideoCategories []struct {
			Level1  string `json:"level1"`
			Level2  string `json:"level2"`
			Level3  string `json:"level3"`
			Source  string `json:"source"`
			Version int    `json:"version"`
		} `json:"video_categories"`
		VideoDetailInfo struct {
			DetailVideoLargeImage struct {
				Height  int    `json:"height"`
				Uri     string `json:"uri"`
				Url     string `json:"url"`
				UrlList []struct {
					Url string `json:"url"`
				} `json:"url_list"`
				Width int `json:"width"`
			} `json:"detail_video_large_image"`
			DirectPlay          int    `json:"direct_play"`
			GroupFlags          int    `json:"group_flags"`
			LastPlayDuration    int    `json:"last_play_duration"`
			ShowPgcSubscribe    int    `json:"show_pgc_subscribe"`
			UseLastDuration     bool   `json:"use_last_duration"`
			VideoId             string `json:"video_id"`
			VideoPreloadingFlag int    `json:"video_preloading_flag"`
			VideoType           int    `json:"video_type"`
			VideoWatchCount     int    `json:"video_watch_count"`
		} `json:"video_detail_info"`
		VideoDuration          int     `json:"video_duration"`
		VideoExclusive         bool    `json:"video_exclusive"`
		VideoId                string  `json:"video_id"`
		VideoLikeCount         int     `json:"video_like_count"`
		VideoPlayInfo          string  `json:"video_play_info"`
		VideoProportion        float64 `json:"video_proportion"`
		VideoProportionArticle float64 `json:"video_proportion_article"`
		VideoUserLike          int     `json:"video_user_like"`
		XgVideoRichText        struct {
		} `json:"xg_video_rich_text"`
		XiRelated bool `json:"xi_related"`
	} `json:"data"`
	Message string `json:"message"`
}

// GetArticleInfo 非点播视频
func GetArticleInfo(Id string) (a ArticleInfo, err error) {
	resp, err := client.Do("get", fmt.Sprintf(upapi, Id), nil)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp.Content, &a)
	if err != nil {
		return
	}
	switch a.Data.Delete {
	case 0:
		return
	case 1:
		err = fmt.Errorf("StatusCode: %d, StatusMessage: %s", a.Data.Delete, a.Data.Content)
		return
	case 2:
		err = fmt.Errorf("StatusCode: %d, StatusMessage: %s", a.Data.Delete, a.Data.Content)
		return
	default:
		err = fmt.Errorf("StatusCode: %d, StatusMessage: 未知错误", a.Data.Delete)
		return
	}
}

// GetVideoInfo 点播视频
func GetVideoInfo(albumId string) (v VideoInfo, err error) {
	params := map[string]string{
		"query_type": "0", //
		"album_id":   albumId,
		"format":     "json",
	}
	req := url.NewRequest()
	req.Params = url.ParseParams(params)
	resp, err := client.Do("get", api, req)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp.Content, &v)
	if err != nil {
		return
	}
	if v.BaseResp.StatusCode != 0 {
		err = fmt.Errorf("StatusCode: %d, StatusMessage: %s", v.BaseResp.StatusCode, v.BaseResp.StatusMessage)
		return
	}
	return
}
func GetMateInfo(ctx context.Context, sharerUrl string) (r *server.Data, code int) {
	var err error
	requestId := ctx.Value("requestId").(string)
	r = &server.Data{}
	log.Debugf("site: %s, start GetMateInfo", web)
	re := regexp.MustCompile(`/(\d+)`)
	find := re.FindStringSubmatch(sharerUrl)
	if len(find) == 0 {
		err = fmt.Errorf("invalid url")
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		code = 2
		return
	}
	albumId := find[1]
	v, err := GetVideoInfo(albumId)
	if err != nil {
		log.Errorf("site: %s, requestId: %s, error: %v, url: %s", web, requestId, err, sharerUrl)
		code = 1
		return
	}
	r.SeriesTitle = v.Album.Title
	r.SeriesId = v.Album.AlbumId
	r.IsSeries = v.Album.AlbumTypeList[0] != 1
	for _, v := range v.BlockList {
		if v.Style == 3 {
			for _, v := range v.Cells {
				r.VideoList = append(r.VideoList, &server.Video{
					EpisodeId:    v.Episode.EpisodeId,
					Episode:      uint32(v.Episode.Seq),
					EpisodeTitle: v.Episode.Title,
					IsVip:        v.Episode.VipPlayControl != 1,
					Extra: map[string]string{
						"vid": v.Episode.VideoInfo.Vid,
					},
				})
			}
		}
	}
	return
}
