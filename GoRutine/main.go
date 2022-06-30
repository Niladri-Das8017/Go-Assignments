package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/sync/errgroup"
)

const (
	URL             = "http://golang.org"
	PATH            = " Saved Picture"
	DownloadChannel = 3
)

//var mut sync.Mutex //pointer

func varInit() map[string]bool {
	imagesUrl := make(map[string]bool)
	return imagesUrl
}

func urlToHtml(url string) (*goquery.Document, error) {
	resp, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func urlLink(url string) (map[string]struct{}, error) {
	resp, err := urlToHtml(url)
	if err != nil {
		return nil, err
	}

	imagesUrl := make(map[string]struct{})
	resp.Find("*").Each(func(index int, item *goquery.Selection) {
		tag := item.Find("img")
		link, _ := tag.Attr("src") //link ,bool
		if link != "" {
			imagesUrl[link] = struct{}{}
		}
	})
	return imagesUrl, nil
}

func downloader(imageUrl map[string]struct{}) error {
	//fmt.Println(imageUrl)
	var sem chan struct{}
	//	var wg sync.WaitGroup
	sem = make(chan struct{}, DownloadChannel)
	defer close(sem)
	//var errs chan error

	//Creating error group
	eg := new(errgroup.Group)
	for value, _ := range imageUrl { //val,bool
		//wg.Add(1)

		//sem <- struct{}{}

		val := value

		eg.Go(func() error {

			//fmt.Println("gorutine")
			err := downlaodImage(val)
			if err != nil {
				return err
			}

			//		defer wg.Done()
			// defer func() {

			// 	<-sem
			// }()

			return nil
		})
	}
	// Wait for all download to complete.
	err := eg.Wait()
	if err != nil {
		return err
	}
	fmt.Println("Successfully Downloaded All the images.")
	return nil
}

func downlaodImage(ImagesUrl string) error {

	err := os.Mkdir(PATH, os.FileMode(0777))
	var Addurl string
	if ImagesUrl[:4] != "http" {
		Addurl = "http:" + ImagesUrl
	} else {
		Addurl = ImagesUrl
	}
	parts := strings.Split(Addurl, "/")

	name := parts[len(parts)-1]

	resp, err := http.Get(Addurl)
	if err != nil {
		return errors.New("unable to fetch url")
	}

	//path := filepath.Join("Saved Pics", name) //Creating Path
	file, err := os.Create(string(PATH + "/" + name))
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return errors.New("unable to copy")
	}
	fmt.Printf("Saving %s  Link := %s\n", PATH+"/"+name, Addurl)
	return nil
}

func main() {
	imageurl, err := urlLink(URL)
	if err != nil {
		log.Fatal(err)
	}

	err = downloader(imageurl)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Downloaded completed")
}
