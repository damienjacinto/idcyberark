package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"idcyberark/counter"
)

func TestRouter(t *testing.T) {
	c := counter.New(counter.MaxCounter)
	r := Router(c)
	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/id/jenkins")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code for /id{jenkins} is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusOK)
	}

	res, err = http.Post(ts.URL+"/id/jenkins", "text/plain", nil)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status code for /id/{jenkins} is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusMethodNotAllowed)
	}

	res, err = http.Get(ts.URL + "/not-exists")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Status code for /id/{jenkins} is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusNotFound)
	}
}