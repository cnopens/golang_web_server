package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/olivere/elastic"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/yangtinngting/golang_web_server/monitor"
)

func main() {

	http.HandleFunc("/", monitor.Monitor(SayHello))
	http.HandleFunc("/queryES", monitor.Monitor(SearchLineByKeyword("http://localhost:9200")))
	http.HandleFunc("/sleep", monitor.Monitor(Sleep))
	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":8080", nil)
}

// SayHello ...
func SayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world")
}

// SearchLineByKeyword search shakespeare book by keyword, and return the first one with high score.
func SearchLineByKeyword(urls ...string) http.HandlerFunc {
	// lazy load elasticsearch
	es, err := elastic.NewClient(elastic.SetURL(urls...))
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to establish connections to elasticsearch")
	}

	/*
		type line struct {
			LineId       int    `json:"line_id, omitempty"`
			LineNumber   string `json:"line_number, omitempty"`
			PlayName     string `json:"play_name, omitempty"`
			Speaker      string `json:"speaker, omitempty"`
			SpeechNumber int    `json:"speech_number, omitempty"`
			TextEntry    string `json:"text_entry", omitempty"`
			Type         string `json:"type, omitempty"`
		}
	*/

	return func(w http.ResponseWriter, r *http.Request) {
		keyword := r.URL.Query().Get("keyword")
		query := elastic.NewMatchQuery("text_entry", keyword)
		bytes, err := monitor.ESQuery(es, "shakespeare", "doc", query)
		if err != nil {
			http.Error(w, "internal failure", 500)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			fmt.Fprint(w, string(bytes))
		}

	}
}

// Sleep ...
func Sleep(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 10)
	io.WriteString(w, "slept 10 seconds")
}
