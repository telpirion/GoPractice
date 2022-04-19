package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	rq := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	handler(w, rq)

	rs := w.Result()
	defer rs.Body.Close()

	data, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Error(err)
	}
	want := "URL"
	s := string(data)
	if !strings.Contains(s, want) {
		t.Errorf("wanted %s; got %s\n", want, s)
	}
}
