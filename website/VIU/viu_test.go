package Viu

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"runtime/debug"
	"testing"
)

var ctx = context.WithValue(context.Background(), "requestId", "test")

func TestGetViuSeries(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			stack := string(debug.Stack())
			fmt.Println(err)
			fmt.Println(stack)

		}
	}()
	r, err := GetMateInfo(ctx, "https://www.viu.com/ott/hk/zh/vod/2336233/%E5%85%A8%E7%9F%A5%E5%B9%B2%E9%A0%90%E8%A6%96%E8%A7%92-2024")
	assert.Equal(t, err, 0)
	assert.Equal(t, "72465", r.SeriesId)
}
