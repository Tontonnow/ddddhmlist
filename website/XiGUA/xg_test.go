package XiGUA

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

var ctx = context.WithValue(context.Background(), "requestId", "test")

func TestGetMateInfo(t *testing.T) {
	r, err := GetMateInfo(ctx, "https://www.ixigua.com/7117099761831772686?logTag=84d96f5b647aa7e1fb48")
	assert.Equal(t, err, 0)
	assert.Equal(t, "7348752092212757019", r.SeriesId)
	assert.Equal(t, "v02043g10000cnu23e3c77u49t96guj0", r.VideoList[0].Extra["vid"])
}
