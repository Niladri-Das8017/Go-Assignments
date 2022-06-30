package helper

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

const PATH = "Saved Picture"
const DownloadChannel = 3

func Downloader(imageUrl []string) error {
		fmt.Println(imageUrl)
	var sem chan struct{}
	var wg sync.WaitGroup
	sem = make(chan struct{}, DownloadChannel)
	defer close(sem)
	var errs chan error

	//Creating error group
	//eg := new(errgroup.Group)
	for _, value := range imageUrl { //val,bool
		wg.Add(1)
		select {
		case sem <- struct{}{}:
		case x := <-errs:

			fmt.Println("err")
			return x

		}

		go func(val string) {
			err := downlaodImage(val)
			if err != nil {
				errs <- err
			}

			defer wg.Done()
			defer func() {

				<-sem
			}()

		}(value)
	}
	wg.Wait()

	return nil
}

func downlaodImage(ImagesUrl string) error {
	err := os.Mkdir("PATH", os.FileMode(0777))
	if err != nil {
		return errors.New("unable to create file directory ")
	}
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

	file, err := os.Create(string(PATH + "/" + name))
	if err != nil {
		return errors.New("unable to create file ")

	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return errors.New("unable to copy")
	}
	fmt.Printf("Saving %s \n", PATH+"/"+name)
	return nil
}
