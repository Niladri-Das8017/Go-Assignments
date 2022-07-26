package helper

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func FetchData(url string) (string, error) {

	//fetching all data from url
	response, err := http.Get(url)
	if err != nil {
		return "", errors.New("FetchData : Failed to fetch response from the url!")
	}

	//it will close te esponse after the required job is done
	defer response.Body.Close()
	//Read all the data as a slice of bytes
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", errors.New("FetchData : Failed to fetch content from response.body!")
	}

	//converting to string
	data := string(content)

	//Calling helper functions
	// imgBox := ImgFetch(strContent)
	// ImgDownload(imgBox)
	return data, nil

}
