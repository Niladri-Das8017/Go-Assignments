package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var dict map[string]int

func TestWC(t *testing.T) {

	content := "aa aa bb aa cc bb aa"
	dict, err := WordCount(content)
	assert.Nil(t, err)

	//expected Result
	expected := 4

	if dict["aa"] != expected {
		t.Errorf("\nWord Count FAILED! \nExpected %d, got %d\n", expected, dict["aa"])
	} else {
		t.Logf("\nWord Count PASSED \nExpected %d, got %d\n", expected, dict["aa"])
	}

}

func TestSort(t *testing.T) {
	sortedWc_Dict := SortWc(dict)

	for i := 0; i < len(sortedWc_Dict)-1; i++ {
		wc := sortedWc_Dict[i]
		nextWc := sortedWc_Dict[i+1]

		if wc.Count < nextWc.Count {
			t.Errorf("\nWordCount Sorting FAILED!\n")
		} else {
			t.Logf("\nWordCount Sorting PASSED \n")
		}
	}
}
