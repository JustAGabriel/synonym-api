package main

import "testing"

// todo: test graph logic:
// 1. add nodes -> retrieve synonyms (direct neighbors)
// 2. add nodes -> retrive synonyms transistively (2. level neighbors)

func TestSimpleSynonymRetrieval(t *testing.T) {
	// arrange
	g := Graph{}
	word := "test"
	word2 := "test2"

	g.ConnectNodes(word, word2)

	// act
	synonyms, err := g.GetDirectNeighbours(word)

	// assert
	if err != nil {
		t.Error("should not throw an error")
	}

	if len(synonyms) != 1 {
		t.Error("should contain exactly one synonym")
	}

	if synonyms[0] != word2 {
		t.Errorf("registered synonym should be '%s'", word2)
	}
}

func TestTransiveSynonymRetrieval(t *testing.T) {
	// arrange
	g := Graph{}
	word := "test"
	word2 := "test2"
	word3 := "test3"

	g.ConnectNodes(word, word2)
	g.ConnectNodes(word2, word3)

	// act
	synonyms, err := g.GetDirectAndSecondLevelNeighbours(word)

	// assert
	if err != nil {
		t.Error("should not throw an error")
	}

	if len(synonyms) != 2 {
		t.Errorf("should contain exactly two synonym, found %d", len(synonyms))
	}

	if synonyms[0] != word2 && synonyms[1] != word3 {
		t.Errorf("registered synonym should be '%s'", word2)
	}
}
