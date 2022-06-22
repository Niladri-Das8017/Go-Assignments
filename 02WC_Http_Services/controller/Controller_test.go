package controller

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWC(t *testing.T) {

	//Creating Request
	text := strings.NewReader("My name is Niladri.")
	req, err := http.NewRequest("POST", "localhost:4040/wordcount", text)
	if err != nil {
		t.Fatal("Could not create request.\n", err)
	}

	//Recorder or ResponseWriter
	rec := httptest.NewRecorder()
	WordCount(rec, req)

	response := rec.Result()

	if response.StatusCode != http.StatusOK {

		t.Error("Expected Status Ok, got ", response.StatusCode)
	} else {
		t.Log(`"
		WordCount:
		Status Code- "`, response.StatusCode)
	}

	defer response.Body.Close()
	result, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal("Could not read response body")
	}
	t.Log("\nTest Result : ", string(result))
}
