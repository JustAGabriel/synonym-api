package main

import "errors"

type Graph map[string]map[string]bool

func (g Graph) ConnectNodes(node1, node2 string) {
	directNeighbours, found := g[node1]
	if !found {
		g[node1] = map[string]bool{
			node2: true,
		}
	} else {
		directNeighbours[node2] = true
	}

	directNeighbours, found = g[node2]
	if !found {
		g[node2] = map[string]bool{
			node1: true,
		}
	} else {
		directNeighbours[node1] = true
	}
}

func (g Graph) GetDirectNeighbours(node string) ([]string, error) {
	neighbourNodes, found := g[node]
	if !found {
		return nil, errors.New("could not find node")
	}

	nodes := []string{}
	for k := range neighbourNodes {
		nodes = append(nodes, k)
	}

	return nodes, nil
}

func (g Graph) GetDirectAndSecondLevelNeighbours(node string) ([]string, error) {
	directNeighbours, err := g.GetDirectNeighbours(node)
	if err != nil {
		return nil, err
	}

	subnodes := []string{}
	for _, directNeighbour := range directNeighbours {
		secondLevelNeighbours, err2 := g.GetDirectNeighbours(directNeighbour)
		if err2 != nil {
			return nil, err2
		}

		secondLevelNeighbours = GetCollectionWithoutElements(secondLevelNeighbours, node)
		subnodes = append(subnodes, secondLevelNeighbours...)
	}

	return append(directNeighbours, subnodes...), nil
}
