package main

import (
	"fmt"
	"imgDownloader/helper"
	"log"
)

func main() {
	fmt.Println("Welcome to Image Downloader")

	url := "http://golang.org"

	//callling helpers
	data, err := helper.FetchData(url)
	if err != nil {
		log.Fatal(err)
	}

	imgBox, err := helper.ImgFetch(data)
	if err != nil {
		log.Fatal(err)
	}

	err = helper.ImgDownload(imgBox)
	if err != nil {
		log.Fatal(err)
	}

}
