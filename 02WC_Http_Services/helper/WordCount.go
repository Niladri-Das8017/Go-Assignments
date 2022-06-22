package helper

import (
	"regexp"
	"strings"
)

func WordCount(content string) map[string]int {

	//Removing all spcial charecters and white spaces from string
	reg, _ := regexp.Compile(`[^\w]`)

	content = reg.ReplaceAllString(content, " ")

	//making slice of a content
	strSlice := strings.Split(content, " ")

	//making dictionarry and store count in it
	wcMap := make(map[string]int)

	for _, word := range strSlice {
		if word != "" {

			wcMap[word]++ //if exist increment else initialize word woth value 1

		}

	}
	return wcMap
}
