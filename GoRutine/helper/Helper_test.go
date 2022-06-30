package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var data string
var imgBox []string
var err error

func TestFetchData(t *testing.T) {
	data, err = FetchData("http://golang.org")
	assert.Nil(t, err)

}

func TestImgFetch(t *testing.T) {

	//Calling helper functions
	imgBox, err = ImgFetch(data)
	assert.Nil(t, err)

	expcted := "/images/go-logo-white.svg"
	if imgBox[0] == expcted {
		t.Log("Feching Image Sources Passed")
	} else {
		t.Error("Feching Image Sources FAILED!")
	}

}

func TestImgDownload(t *testing.T) {
	
	err = ImgDownload(imgBox)
	assert.Nil(t, err)

}
