// © 2019 Imhotep Software LLC. All rights reserved.
package web

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gopherland/ntnx_labs/deploy/picker/internal/dictionary"
	"github.com/stretchr/testify/assert"
)

func TestLoadHandler(t *testing.T) {
	var (
		rr   = httptest.NewRecorder()
		r, _ = http.NewRequest("GET", "http://example.com/api/v1/load", nil)
	)

	q := url.Values{"name": []string{"musicians"}}
	r.URL.RawQuery = q.Encode()

	h := New("../assets")
	h.LoadHandler(rr, r)
	assert.Equal(t, http.StatusOK, rr.Code)

	var resp Response
	err := json.NewDecoder(rr.Body).Decode(&resp)
	assert.Nil(t, err)
	assert.Equal(t, 5, resp.WordCount)
	assert.Equal(t, "assets", resp.AssetDir)
	assert.Equal(t, "musicians", resp.Dictionary)
}

func TestPickerHandler(t *testing.T) {
	var (
		rr   = httptest.NewRecorder()
		r, _ = http.NewRequest("GET", "http://example.com/api/v1/picker", nil)
	)

	var err error
	h := New("assets")
	h.words, err = dictionary.Load("../assets", "musicians")
	assert.Nil(t, err)

	h.PickHandler(rr, r)
	assert.Equal(t, http.StatusOK, rr.Code)

	var resp struct {
		Word string
	}
	err = json.NewDecoder(rr.Body).Decode(&resp)
	assert.Nil(t, err)
	assert.False(t, resp.Word == "")
}
