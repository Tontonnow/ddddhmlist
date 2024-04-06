package MytvSuper

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	u   = "https://www.mytvsuper.com/tc/programme/tokyoloveactstory0001_139100/%E6%9D%B1%E4%BA%AC%E6%84%9B%E6%83%85%E5%8B%95%E4%BD%9C%E6%95%85%E4%BA%8B-(4K%E7%89%88)"
	ctx = context.WithValue(context.Background(), "requestId", "test")
)

func TestGetProgrammeId(t *testing.T) {
	got := getProgrammeId(u)
	assert.Equal(t, "139100", got)
}

func TestGetMateInfo(t *testing.T) {
	r, err := GetMateInfo(ctx, u)
	assert.Equal(t, err, 0)
	assert.Equal(t, "東京愛情動作故事 (4K版)", r.SeriesTitle)
	assert.Equal(t, "748903", r.VideoList[0].Extra["video_id"])
}
