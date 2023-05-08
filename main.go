package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type AddSynonymsPayload struct {
	Synonyms []string
}

type GetSynonymsPayload struct {
	Word     string
	Transive bool
}

type GetSynonymsResponsePayload struct {
	Synonyms []string
}

var synonyms Graph = Graph{}

func handleNew(w http.ResponseWriter, r *http.Request) {
	log.Default().Printf("graph content: %+v", synonyms)
	if r.Method == http.MethodPost {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		synonymsPayload := &AddSynonymsPayload{}
		err = json.Unmarshal(body, synonymsPayload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		syns := synonymsPayload.Synonyms
		for idx, word := range syns {
			synsWithoutWord := GetCollectionWithoutElements(syns, syns[idx:]...)
			for _, word2 := range synsWithoutWord {
				if word2 != word {
					synonyms.ConnectNodes(word, word2)
				}
			}
		}

		return
	}
	if r.Method == http.MethodGet {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		getSynonymsPayload := &GetSynonymsPayload{}
		err = json.Unmarshal(body, getSynonymsPayload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		log.Default().Printf("received payload: %+v", getSynonymsPayload)

		var responseBodyContent any
		if getSynonymsPayload.Transive {
			log.Default().Println("get synonyms (incl. transive ones)")
			synonyms2, err2 := synonyms.GetDirectAndSecondLevelNeighbours(getSynonymsPayload.Word)
			if err2 != nil {
				http.Error(w, err2.Error(), http.StatusInternalServerError)
			}

			responseBodyContent = GetSynonymsResponsePayload{
				Synonyms: synonyms2,
			}
		} else {
			log.Default().Println("get synonyms")
			synonyms2, err2 := synonyms.GetDirectNeighbours(getSynonymsPayload.Word)
			if err2 != nil {
				http.Error(w, err2.Error(), http.StatusInternalServerError)
			}

			responseBodyContent = GetSynonymsResponsePayload{
				Synonyms: synonyms2,
			}
		}

		responseBodyBytes, err2 := json.Marshal(responseBodyContent)
		if err2 != nil {
			http.Error(w, err2.Error(), http.StatusInternalServerError)
			return
		}

		_, err3 := w.Write(responseBodyBytes)
		if err3 != nil {
			http.Error(w, err3.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "Invalid operation", http.StatusBadRequest)
}

func main() {
	http.HandleFunc("/synonyms", handleNew)
	log.Fatal(http.ListenAndServe("localhost:5432", nil))
}
