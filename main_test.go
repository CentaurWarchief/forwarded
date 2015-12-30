package forwarded_test

import (
	"net/http"
	"testing"

	"github.com/CentaurWarchief/forwarded"
	"github.com/stretchr/testify/assert"
)

func TestWas(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)

	assert.False(t, forwarded.Was(req))

	req.Header.Set("Forwarded", "")
	assert.False(t, forwarded.Was(req))

	req.Header.Set("Forwarded", "for=192.0.2.60; proto=http; by=203.0.113.43")
	assert.True(t, forwarded.Was(req))
}

func TestParseEmptyForwardedHeader(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Forwarded", "")

	assert.Nil(t, forwarded.Parse(req))
}

func TestParseForwardedHeader(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	req.Header.Set("Forwarded", "for=192.0.2.60; proto=http; by=203.0.113.43")

	f := forwarded.Parse(req)

	assert.NotNil(t, f)
	assert.Equal(t, "192.0.2.60", f.For)
	assert.Equal(t, "http", f.Proto)
	assert.Equal(t, "203.0.113.43", f.By)
	assert.Equal(t, f.Host, "")
}
