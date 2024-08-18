package main

import (
	"encoding/json"
	"log"
	"github.com/sugoto/zap-reduce/mapreduce"
	"net/http"
)

type RequestBody struct {
	Text string `json:"text"`
}

type ResponseBody struct {
	WordCounts map[string]int `json:"word_counts"`
}

func processTextHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	wordCounts := mapreduce.Execute(reqBody.Text)

	respBody := ResponseBody{WordCounts: wordCounts}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respBody)
}

func main() {
	// Serve static files (HTML, CSS, JS)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	// API endpoint for processing text
	http.HandleFunc("/process", processTextHandler)

	// Start the server on port 8080
	port := "8080"
	log.Printf("Starting server on :%s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
}
