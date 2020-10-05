package cmd

import (
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestHitEndPointReturns200(t *testing.T) {
	var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	assert.Equal(t, HitEndPoint(apiStub.URL), "Successfully hit httpbin.org")
}

func TestHitEndPointReturns500(t *testing.T) {
	var apiStub = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))

	assert.Equal(t, HitEndPoint(apiStub.URL), "httpbin returned an http status error: 500")
}