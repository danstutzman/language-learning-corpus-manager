package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouter(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/hello")
	assert.Nil(t, err)
	assert.Equal(t, resp.StatusCode, http.StatusOK)

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello world!", string(b))
}
