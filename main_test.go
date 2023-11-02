package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func makeRequest(method, path, body string) *httptest.ResponseRecorder {
	router := getRouter()
	req, _ := http.NewRequest(method, path, strings.NewReader(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func TestGetAlbums(t *testing.T) {
	w := makeRequest("GET", "/albums", "")
	if w.Code != http.StatusOK {
		t.Fatalf("Wrong response code.")
		t.Fail()
	}
}

func TestGetAlbumByID(t *testing.T) {
	w := makeRequest("GET", "/albums/1", "")
	if w.Code != http.StatusOK {
		t.Fatal("Wrong response code")
		t.Fail()
	}
}

func TestIDNotFound(t *testing.T) {
	w := makeRequest("GET", "/albums/-1", "")
	if w.Code != http.StatusNotFound {
		t.Fatal("Must be 404")
		t.Fail()
	}
}

func TestPostAndGetAlbum(t *testing.T) {
	albumPost := album{ID: "5", Title: "TestTitle", Artist: "TestArtist"}
	jsn, _ := json.Marshal(albumPost)
	makeRequest("POST", "/albums/5", string(jsn))
	w := makeRequest("GET", "/albums/5", "")
	var albumGet album
	// TODO: find way to convert *buffer.Body to []byte 
	body, _ := ioutil.ReadAll(w.Body)
	json.Unmarshal(body, &albumGet)
	if albumGet.isEqual(albumPost) {
		t.Fatal("Objects not equal")
		t.Fail()
	}

}
