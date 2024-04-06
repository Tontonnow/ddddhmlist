package MyVideo

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ctx = context.WithValue(context.Background(), "requestId", "test")

func TestGetMateInfo(t *testing.T) {

	t.Run("电影", func(t *testing.T) {
		u := "https://www.myvideo.net.tw/details/0/385248"
		v, err := GetMateInfo(ctx, u)
		assert.Equal(t, err, 0)
		assert.False(t, v.IsSeries)
		assert.Equal(t, "385248", v.SeriesId)
		assert.Equal(t, "虛擬獵殺令", v.SeriesTitle)
	})
	t.Run("剧集", func(t *testing.T) {
		u := "https://www.myvideo.net.tw/details/3/26950"
		v, err := GetMateInfo(ctx, u)
		assert.Equal(t, err, 0)
		assert.True(t, v.IsSeries)
		assert.Equal(t, "26950", v.SeriesId)
		assert.Equal(t, "王者天下 第五季", v.SeriesTitle)
		assert.Equal(t, "392737", v.VideoList[0].EpisodeId)

	})

}
