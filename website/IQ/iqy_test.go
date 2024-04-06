package IQ

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

var ctx = context.WithValue(context.Background(), "requestId", "test")

func TestGetVid(t *testing.T) {
	t.Run("GetVid", func(t *testing.T) {
		v, err, c := GetVid("https://www.iqiyi.com/v_22dchdtcba0.html?r_area=recent_popular&r_source=1001&bkt=hp_bkt_02&e=3c7f1fb60e7e64bb39e37b55194f7711&stype=2&vfrm=pcw_home&vfrmblk=pcw_home_hot&vfrmrst=pcw_home_hot_image2")
		assert.Nil(t, err)
		assert.Equal(t, "22dchdtcba0", v)
		assert.False(t, c)
	})
	t.Run("GetVid", func(t *testing.T) {
		v, err, c := GetVid("https://www.iq.com/album/%E8%8A%B1%E6%BA%AA%E8%AE%B0-2023-1tzhhn9abvp?lang=zh_cn")
		assert.Nil(t, err)
		assert.Equal(t, "1tzhhn9abvp", v)
		assert.False(t, c)
	})
	t.Run("GetVid", func(t *testing.T) {
		v, err, c := GetVid("https://m.iqiyi.com/m/playShare?shareId=NTMzODkxNTA3MjE3MzQwMA%3D%3D&positiveId=NTMzODkxNTA3MjE3MzQwMA%3D%3D&type=0&rpage=sharepage_new")
		assert.Nil(t, err)
		assert.Equal(t, "5338915072173400", v)
		assert.True(t, c)
	})
}

func TestGetMateInfo(t *testing.T) {

	t.Run("电影", func(t *testing.T) {
		v, err := GetMateInfo(ctx, "https://www.iqiyi.com/v_2704o3ve8qw.html")
		assert.Equal(t, err, 0)
		assert.Equal(t, "8149704202414500", v.SeriesId)
		assert.Equal(t, "养蜂人", v.SeriesTitle)
		assert.False(t, v.IsSeries)
		fmt.Println(v.Extra)
	})
	t.Run("电视剧 动漫 纪录片 儿童 知识", func(t *testing.T) {
		v, err := GetMateInfo(ctx, "https://www.iqiyi.com/v_22dchdtcba0.html?r_area=recent_popular&r_source=1001&bkt=hp_bkt_02&e=3c7f1fb60e7e64bb39e37b55194f7711&stype=2&vfrm=pcw_home&vfrmblk=pcw_home_hot&vfrmrst=pcw_home_hot_image2")
		assert.Equal(t, err, 0)
		assert.Equal(t, "8826732176010901", v.SeriesId)
		assert.Equal(t, "乘风踏浪", v.SeriesTitle)
		assert.True(t, v.IsSeries)
	})
	t.Run("长篇", func(t *testing.T) {
		v, err := GetMateInfo(ctx, "https://www.iqiyi.com/v_19rrok4nt0.html")
		assert.Equal(t, err, 0)
		assert.Equal(t, "202861101", v.SeriesId)
		assert.Equal(t, "航海王", v.SeriesTitle)
		assert.True(t, v.IsSeries)
		for i := 0; i < 5; i++ {
			r := rand.Intn(len(v.VideoList))
			assert.Equal(t, v.VideoList[r].Episode, uint32(r+1))
		}
		fmt.Println(v.Extra)
	})
	t.Run("综艺", func(t *testing.T) {
		v, err := GetMateInfo(ctx, "https://www.iqiyi.com/v_19ru98a2w4.html?vfrm=pcw_playpage&vfrmblk=80521_listbox_positive&vfrmrst=0")
		assert.Equal(t, err, 0)
		assert.Equal(t, "246967201", v.SeriesId)
		assert.Equal(t, "坑王驾到第4季", v.SeriesTitle)
		assert.True(t, v.IsSeries)
		fmt.Println(v.Extra)
	})

}
