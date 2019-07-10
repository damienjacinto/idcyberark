package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"idcyberark/counter"
	"github.com/gorilla/mux"
)

func TestIdcyberark(t *testing.T) {
	request, _ := http.NewRequest("GET", "/id/jenkins", nil)
	c := counter.New()
	w := httptest.NewRecorder()
	id := idcyberark(c)
	id(w, request)
	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if have, want := resp.StatusCode, http.StatusOK; have != want {
		t.Errorf("Status code is wrong. Have: %d, want: %d.", have, want)
	}	

	if have, want := string(body), "1"; have != want {
		t.Errorf("Result call /id/{jenkins} failed. Have: %s, want: %s.", have, want)
	}
	
}