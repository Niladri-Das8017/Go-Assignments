package helper

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sync/errgroup"
)

func ImgDownload(imgBox []string) error {

	fmt.Println("Downloading Images...")

	//Creating error group
	eg := new(errgroup.Group)

	for _, src := range imgBox {

		imgUrl := src

		eg.Go(func() error {

			//Seperating img source, make it easy to identify name of image
			//eg: /images/logos/google.svg as google.svg => [ images logos google.svg ]
			splitSrc := strings.Split(imgUrl, "/")

			//Identidy the name from image source using splitSrc
			imgName := splitSrc[len(splitSrc)-1]

			//Image source parsing
			result, err := url.Parse(imgUrl)
			if err != nil {
				return err
			}

			//If the source have "https" scheme, then
			if result.Scheme == "https" {

				//SAVE IMAGE
				imgFolderPath := "C:/Users/Niladri Das/go/Go-Assignments/04ImgDownloader/img"
				err := os.Mkdir(imgFolderPath, os.FileMode(0777)) //creating path
				filePath := filepath.Join(imgFolderPath, imgName) //Creating Path
				file, err := os.Create(filePath)
				if err != nil {
					return err
				}

				defer file.Close()

				response, err := http.Get(imgUrl)
				if err != nil {
					return err
				}
				defer response.Body.Close()

				//Read from reesponse.Body and write to file
				file.ReadFrom(response.Body)
				fmt.Println("Downloaded: ", imgName, "\tSRC: ", imgUrl)
			}

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
