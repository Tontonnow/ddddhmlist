package QQTV

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ctx = context.WithValue(context.Background(), "requestId", "test")

func TestGetPlayInfo(t *testing.T) {
	t.Run("ExtractCid", func(t *testing.T) {
		url := "https://v.qq.com/x/cover/mzc002000mc2t5k/m004878ddyw.html"
		cid := ExtractCid(url)
		assert.Equal(t, "mzc002000mc2t5k", cid)
	})
	t.Run("qqtv", func(t *testing.T) {
		url := "https://m.v.qq.com/x/m/play?vid=c0048s8n6p3&cid=mzc00200as5tv65&ptag=share_11_11&url_from=share&second_share=0&share_from=qqf&pgid=page_detail&mod_id=mod_toolbar_new"
		d, err := GetMateInfo(ctx, url)
		assert.Equal(t, err, 0)
		assert.Equal(t, "mzc00200as5tv65", d.SeriesId)
		assert.Equal(t, "哈哈哈哈哈 第4季", d.SeriesTitle)
		assert.Equal(t, "m0048l524x0", d.VideoList[0].EpisodeId)
	})
	t.Run("wetv", func(t *testing.T) {
		d, err := GetWeTvMateInfo(ctx, "https://wetv.vip/zh-cn/play/0cf9i81kpms8469-%E5%AD%A4%E7%94%B7%E5%AF%A1%E5%A5%B3/q0035x0szr1-%E5%AD%A4%E7%94%B7%E5%AF%A1%E5%A5%B3")
		assert.Equal(t, err, 0)
		assert.Equal(t, "0cf9i81kpms8469", d.SeriesId)
		assert.Equal(t, "孤男寡女", d.SeriesTitle)
	})
	return
}
