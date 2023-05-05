package main

import (
	"encoding/json"
	"io"
	"net/http"
)

var (
	synonyms map[string]map[string]bool
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
			addSynonym(word, word2)
		}
	}
}

func addSynonym(w1, w2 string) {
	existingSynonyms, found := synonyms[w1]
	if !found {
		synonyms[w1] = map[string]bool{
			w2: true,
		}
	} else {
		existingSynonyms[w2] = true
	}

	existingSynonyms, found = synonyms[w2]
	if !found {
		synonyms[w2] = map[string]bool{
			w1: true,
		}
	} else {
		existingSynonyms[w1] = true
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
