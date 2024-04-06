package LETV

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ctx = context.WithValue(context.Background(), "requestId", "test")

func TestGetList(t *testing.T) {
	t.Run("剧集", func(t *testing.T) {
		url := "https://www.le.com/ptv/vplay/77532622.html"
		d, err := GetMateInfo(ctx, url)
		assert.Equal(t, err, 0)
		assert.Equal(t, "10070229", d.SeriesId)
		assert.Equal(t, "川流", d.SeriesTitle)
	})
	t.Run("电影", func(t *testing.T) {
		url := "https://www.le.com/ptv/vplay/77515961.html"
		d, err := GetMateInfo(ctx, url)
		assert.Equal(t, err, 0)
		assert.Equal(t, "10069543", d.SeriesId)
		assert.Equal(t, "速度与激战", d.SeriesTitle)

	})
}
