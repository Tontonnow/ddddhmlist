package Litv

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ctx = context.WithValue(context.Background(), "requestId", "test")

func TestGetList(t *testing.T) {
	t.Run("剧集", func(t *testing.T) {
		url := "https://www.litv.tv/comic/watch/VOD00297197"
		d, err := GetMateInfo(ctx, url)
		assert.Equal(t, err, 0)
		assert.Equal(t, "64742", d.SeriesId)
		assert.Equal(t, "咒術迴戰 第二季", d.SeriesTitle)
		assert.Equal(t, "VOD00298361", d.VideoList[0].EpisodeId)
	})
	t.Run("电影", func(t *testing.T) {
		url := "https://www.litv.tv/movie/watch/VOD00321615"
		d, err := GetMateInfo(ctx, url)
		assert.Equal(t, err, 0)
		assert.Equal(t, "69114", d.SeriesId)
		assert.Equal(t, "嫌疑犯X", d.SeriesTitle)
		assert.Contains(t, d.Extra["assets"], "000000M001_5000K")

	})
}
