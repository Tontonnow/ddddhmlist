package IQ

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewIq(t *testing.T) {
	i := NewIq()
	err := i.RefreshToken()
	assert.Nil(t, err)

}
func TestGetIqMateInfo(t *testing.T) {

	t.Run("电影", func(t *testing.T) {
		u := "https://www.iq.com/album/%E5%BA%9F%E5%93%81%E9%A3%9E%E8%BD%A6-2024-1bjmxliiwxc?lang=zh_cn"
		v, err := GetIqMateInfo(ctx, u)
		assert.Equal(t, err, 0)
		assert.Equal(t, "4876065378052700", v.SeriesId)
		assert.Equal(t, "废品飞车", v.SeriesTitle)
	})
	t.Run("剧集", func(t *testing.T) {
		u := "https://www.iq.com/play/%E7%B9%81%E6%98%9F%E4%B9%8B%E5%9F%8E-%E7%AC%AC1%E9%9B%86-1b7fo0drgkc?lang=zh_cn"
		v, err := GetIqMateInfo(ctx, u)
		assert.Equal(t, err, 0)
		assert.Equal(t, "3548480898144901", v.SeriesId)
		assert.Equal(t, "繁星之城", v.SeriesTitle)
	})
}
