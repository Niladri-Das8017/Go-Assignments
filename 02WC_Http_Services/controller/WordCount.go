package controller

import (
	"encoding/json"
	"httpservices/helper"
	"io/ioutil"
	"net/http"
	"strings"
)

func WordCount(w http.ResponseWriter, r *http.Request) {

	//defer body
	defer r.Body.Close()

	//read response
	dataBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Could Not Read Body"))
		return
	}

	//Collecting text from request
	content := string(dataBytes)
	//Removing white spaces
	content = strings.TrimSpace(content)

	//Checking empty string by "POST"
	if content == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("500 - Empty String sent by POST!!!")
		return
	}

	//Calling Helper
	wcMap, err := helper.WordCount(content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("500 - Something wrong happened on our side")
		return
	}

	//Check if there is any valid word
	if len(wcMap) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("500 - No valid WORD!!!")
		return

	}

	//Calling Helper
	sortedWcMap := helper.SortWc(wcMap)

	//printing...
	json.NewEncoder(w).Encode(sortedWcMap)

}
