package YOUKUTV

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ctx = context.WithValue(context.Background(), "requestId", "test")

func TestGetMateInfo(t *testing.T) {
	t.Run("剧集", func(t *testing.T) {
		U := "https://v.youku.com/v_show/id_XMzk1NjM1MjAw.html?s=cc003400962411de83b1&spm=a2hje.13141534.1_3.d_1_1&scm=20140719.apircmd.239064.video_XMzk1NjM1MjAw"
		d, err := GetMateInfo(ctx, U)
		assert.Equal(t, err, 0)
		assert.Equal(t, "名侦探柯南", d.SeriesTitle)
		//assert.Equal(t, "XMzk1NjM1MjAw", d.VideoList[0].EpisodeId)
		assert.True(t, len(d.VideoList) > 1170) //可能和ip有关
	})
	t.Run("电影", func(t *testing.T) {
		U := "https://v.youku.com/v_show/id_XNjAxNzc3NjE3Ng==.html"
		d, err := GetMateInfo(ctx, U)
		assert.Equal(t, err, 0)
		assert.Equal(t, "奇迹·笨小孩 IMAX版", d.SeriesTitle)
	})
	t.Run("国际版", func(t *testing.T) {
		U := "https://www.youku.tv/v/v_show/id_XNjM4MTM3MDgwNA==.html"
		d, err := GetMateInfo(ctx, U)
		assert.Equal(t, err, 0)
		assert.Equal(t, "花间令 简体版 01", d.VideoList[0].EpisodeTitle)
	})
}
