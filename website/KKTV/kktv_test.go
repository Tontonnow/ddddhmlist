package KKTV

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ctx = context.WithValue(context.Background(), "requestId", "test")

func TestGetEpq(t *testing.T) {
	url := "https://kktv.me/titles/06001673"
	d, err := GetMateInfo(ctx, url)
	assert.Equal(t, err, 0)
	assert.NotNil(t, d)
	assert.Equal(t, "葬送的芙莉蓮", d.SeriesTitle)
	assert.Equal(t, "06001673", d.SeriesId)
	assert.Equal(t, "06001673010001", d.VideoList[0].EpisodeId)
	assert.Contains(t, d.VideoList[0].Extra["dash"], "https://theater.kktv.com.tw/73/06001673010001")
}
