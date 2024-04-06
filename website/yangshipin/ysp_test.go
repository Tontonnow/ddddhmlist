package yangshipin

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ctx = context.WithValue(context.Background(), "requestId", "test")

func TestGetEpq(t *testing.T) {
	pid := "600002264"
	var DataList []*CnYangshipinOmstvCommonProtoProgramModel_Program
	DataList = GetEpg(pid)
	assert.Equal(t, "18229996", DataList[0].ProgramId) //不准确，仅供参考
}

func TestGetEpgId(t *testing.T) {
	Data := GetEpgId()
	assert.NotNil(t, Data)
}

func TestGetList(t *testing.T) {
	u := "https://www.yangshipin.cn/#/video/home?vid=m000071dlup&cid=lbl89jz9be9npgx"
	list, err := GetMateInfo(ctx, u)
	assert.Equal(t, err, 0)
	assert.Equal(t, "lbl89jz9be9npgx", list.SeriesId)
	assert.Equal(t, "千秋诗颂（4K）", list.SeriesTitle)
	assert.Equal(t, "m000071dlup", list.VideoList[0].EpisodeId)
}
