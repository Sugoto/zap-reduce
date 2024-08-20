package api

import (
	"encoding/json"
	"io/ioutil"
	"github.com/sugoto/zap-reduce/mapreduce"
	"net/http"
)

type RequestBody struct {
	Text string `json:"text"`
}

type ResponseBody struct {
	WordCounts map[string]int `json:"word_counts"`
}

func ProcessText(w http.ResponseWriter, r *http.Request) {
	var reqBody RequestBody
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	wordCounts := mapreduce.Execute(reqBody.Text)
	respBody := ResponseBody{WordCounts: wordCounts}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respBody)
}
