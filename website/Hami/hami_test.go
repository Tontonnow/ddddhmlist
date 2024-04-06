package Hami

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ctx = context.WithValue(context.Background(), "requestId", "test")

func TestGetMateInfo(t *testing.T) {
	t.Run("GetMateInfo", func(t *testing.T) {
		v, err := GetMateInfo(ctx, "https://hamivideo.hinet.net/hamivideo/product/152279.do?cs=2")
		assert.Equal(t, err, 0)
		assert.Equal(t, "152279", v.SeriesId)
		assert.Equal(t, "火神的眼淚", v.SeriesTitle)
		assert.Equal(t, "4k", v.VideoList[0].Extra["quality"])

	})
}
