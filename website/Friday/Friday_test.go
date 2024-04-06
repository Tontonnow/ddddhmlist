package Friday

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetEpq(t *testing.T) {
	var ctx = context.WithValue(context.Background(), "requestId", "test")
	err := RefreshToken()
	assert.Nil(t, err)
	t.Run("drama", func(t *testing.T) {
		url := "https://video.friday.tw/drama/detail/3356"
		d, err := GetMateInfo(ctx, url)
		assert.Equal(t, err, 0)
		assert.True(t, d.IsSeries)
		assert.Equal(t, "青春時代", d.SeriesTitle)
		assert.Equal(t, "103581", d.VideoList[0].Extra["ContentId"])
		assert.Equal(t, "2", d.VideoList[0].Extra["ContentType"])
		assert.Equal(t, "114913", d.VideoList[0].Extra["StreamingId"])
		assert.Equal(t, "2", d.VideoList[0].Extra["StreamingType"])

	})
	t.Run("movie", func(t *testing.T) {
		url := "https://video.friday.tw/movie/detail/102915/%E6%B0%B4%E9%AC%BC"
		d, err := GetMateInfo(ctx, url)
		assert.Equal(t, err, 0)
		assert.False(t, d.IsSeries)
		assert.Equal(t, "水鬼", d.SeriesTitle)
		assert.Equal(t, "102915", d.Extra["ContentId"])
		assert.Equal(t, "1", d.Extra["ContentType"])
		assert.Equal(t, "114169", d.Extra["StreamingId"])
		assert.Equal(t, "2", d.Extra["StreamingType"])
	})
}
