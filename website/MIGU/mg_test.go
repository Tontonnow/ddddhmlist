package MIGU

import (
	"context"
	"github.com/Tontonnow/ddddhmlist/server"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ctx = context.WithValue(context.Background(), "requestId", "test")

func TestGetVideoInfo(t *testing.T) {
	type args struct {
		ctx context.Context
		u   string
	}
	tests := []struct {
		name    string
		args    args
		want    *server.Data
		wantErr bool
	}{
		{
			name: "TestGetVideoInfo",
			args: args{
				ctx: ctx,
				u:   "https://m.miguvideo.com/m/detail/899519544?channelId=10010001005",
			},
			want: &server.Data{
				SeriesId:    "5105268712",
				SeriesTitle: "龙凤店传奇",
			},
			wantErr: false,
		},
		{
			name: "TestGetVideoInfo",
			args: args{
				ctx: ctx,
				u:   "https://www.miguvideo.com/p/detail/887504017",
			},
			want: &server.Data{
				SeriesId:    "5105231814",
				SeriesTitle: "以爱为营",
			},
			wantErr: false,
		},
		{
			name: "TestGetVideoInfo",
			args: args{
				ctx: ctx,
				u:   "https://m.miguvideo.com/m/detail/905514014?channelId=10010001005",
			},
			want: &server.Data{
				SeriesId:    "5105269347",
				SeriesTitle: "饥饿游戏：鸣鸟与蛇之歌",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, err := GetMateInfo(tt.args.ctx, tt.args.u)
			assert.Equal(t, err, 0)
			assert.Equal(t, tt.want.SeriesId, d.SeriesId)
			assert.Equal(t, tt.want.SeriesTitle, d.SeriesTitle)
		})
	}
}
