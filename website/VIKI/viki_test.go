package VIKI

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ctx = context.WithValue(context.Background(), "requestId", "test")

func TestGetId(t *testing.T) {

	t.Run("Test GetId", func(t *testing.T) {
		u := "https://www.viki.com/movies/38677c-bo-knows-love"
		movies, id := GetId(u)
		assert.Equal(t, "movies", movies)
		assert.Equal(t, "38677c", id)
	})
	t.Run("Test GetId", func(t *testing.T) {
		u := "https://www.viki.com/tv/39828c-decline"
		movies, id := GetId(u)
		assert.Equal(t, "tv", movies)
		assert.Equal(t, "39828c", id)
	})
}
func TestGetMateInfo(t *testing.T) {
	t.Run("Test movies", func(t *testing.T) {
		u := "https://www.viki.com/movies/38677c-bo-knows-love"
		movies, err := GetMateInfo(ctx, u)
		assert.Equal(t, err, 0)
		assert.Equal(t, "我叫梁山伯", movies.SeriesTitle)
	})
	t.Run("Test tv", func(t *testing.T) {
		u := "https://www.viki.com/tv/39828c-decline"
		movies, err := GetMateInfo(ctx, u)
		assert.Equal(t, err, 0)
		assert.Equal(t, "江湖少年诀", movies.SeriesTitle)
		assert.Equal(t, "1246978v", movies.VideoList[0].EpisodeId)
	})
}
