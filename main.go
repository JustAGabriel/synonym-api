package main

import (
	"encoding/json"
	"io"
	"net/http"
)

var (
	synonyms Graph
)

type SynonymsPayload struct {
	Synonyms []string
}

func getCollectionWithoutElements[T comparable](collection []T, excludes []T) []T {
	r := []T{}
	for _, v := range collection {
		for _, ex := range excludes {
			if ex == v {
				continue
			}
		}
		r = append(r, v)
	}
	return r
}

func addSynonyms(syns []string) {
	for idx, word := range syns {
		synsWithoutWord := getCollectionWithoutElements(syns, syns[0:idx])
		for _, word2 := range synsWithoutWord {
			synonyms.ConnectNodes(word, word2)
		}
	}
}

func handleNew(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		synonymsPayload := &SynonymsPayload{}
		err = json.Unmarshal(body, synonymsPayload)
		if err != nil {
			panic(err)
		}

		addSynonyms(synonymsPayload.Synonyms)
	}
	if r.Method == http.MethodGet {

	}

	http.Error(w, "Invalid operation", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/synonym", handleNew)
}
