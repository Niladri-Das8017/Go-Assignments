package helper

import (
	"sort"
)

type wc struct {
	Word  string `json:"word"`
	Count int    `json:"count"`
}

func SortWc(dict map[string]int) []wc {

	var wC_Dict []wc
	var wc wc

	//A slice that contains words
	words := make([]string, 0, len(dict))
	for w := range dict {
		words = append(words, w)
	}

	//Sorting Words by count in slice
	sort.Slice(words, func(i, j int) bool {
		return dict[words[i]] > dict[words[j]]
	})

	//creating Slice of wordcount in sorted order
	for i := 0; i < len(words); i++ {
		word := words[i]
		wc.Word = word
		wc.Count = dict[word]
		wC_Dict = append(wC_Dict, wc)
	}

	return wC_Dict

}
