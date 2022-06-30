package main

import (
	"fmt"
	"imgDownloaderGorutine/helper"

	"log"
)

func main() {
	fmt.Println("Welcome to Image Downloader")

	url := "https://www.amazon.in/"

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
