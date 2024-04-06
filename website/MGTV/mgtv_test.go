package MGTV

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ctx = context.WithValue(context.Background(), "requestId", "test")

func TestGetVideoInfo(t *testing.T) {
	t.Log("TestGetVideoInfo")
	t.Run("TestGetVideoInfo", func(t *testing.T) {
		t.Log("Test for movie")
		u := "https://www.mgtv.com/b/598961/20728335.html?fpa=1663&fpos=&lastp=ch_movie"
		_, err := GetMateInfo(ctx, u)
		assert.Equal(t, err, 0)
	})
	t.Run("TestGetVideoInfo", func(t *testing.T) {
		t.Log("Test for TV series")
		u := "https://www.mgtv.com/b/338497/8398205.html"
		_, err := GetMateInfo(ctx, u)
		assert.Equal(t, err, 0)
	})

}
