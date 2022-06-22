package helper

import (
	"testing"
)

func TestHelper(t *testing.T) {

	resultMap := WordCount("aa bb cc aa bb a na jfndj jdn ajnd aa aanj nn bb annjkfnjn n n n jndfnakl jnfjn njnf jnjnjd knfknkd ajnfjnj ajnjnjk ajnfnkls ajnnl nfnknf snfjnf anfkn knkfnk dkfnfn knka cc c c jndfna x")

	//Expeced Result
	expected := 3

	if resultMap["n"] != expected {
		t.Errorf("Word Count FAILED! \nExpected %d, got %d\n", expected, resultMap["n"])
	} else {
		t.Logf("Word Count PASSED \nExpected %d, got %d\n", expected, resultMap["n"])
	}

	err := PrintTopTen(resultMap)
	if err != nil {
		t.Log(err)
	}

}
