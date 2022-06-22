package helper

import (
	"regexp"
	"strings"
)

func WordCount(str string) map[string]int {

	//Splitting string into slice
	strSlice := strings.Fields(str)

	//Removing all spcial charecters and white spaces
	reg, _ := regexp.Compile(`[^\w]`)
	for i := range strSlice {
		strSlice[i] = reg.ReplaceAllString(strSlice[i], "")
	}
	//making dictionarry and store count in it
	dict := make(map[string]int)

	for _, word := range strSlice {

		if word != "" {
			dict[word]++ //if exist increment else initialize word woth value 1
		}
		
	}

	return dict
}
