package Bilibili

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMateInfo(t *testing.T) {
	var ctx = context.WithValue(context.Background(), "requestId", "test")
	t.Run("drama", func(t *testing.T) {
		url := "https://www.bilibili.com/bangumi/play/ep811437"
		d, err := GetMateInfo(ctx, url)
		assert.Equal(t, err, 0)
		assert.True(t, d.IsSeries)
		assert.Equal(t, "咸鱼哥 第二季", d.SeriesTitle)
		assert.Equal(t, "1447505461", d.VideoList[0].EpisodeId)

	})
	t.Run("movie", func(t *testing.T) {
		url := "https://www.bilibili.com/bangumi/play/ss45822"
		d, err := GetMateInfo(ctx, url)
		assert.Equal(t, err, 0)
		assert.False(t, d.IsSeries)
		assert.Equal(t, "金手指", d.SeriesTitle)
		assert.Equal(t, "1432240939", d.VideoList[0].EpisodeId)
		assert.Equal(t, "814488", d.VideoList[0].Extra["ep_id"])
	})
}
