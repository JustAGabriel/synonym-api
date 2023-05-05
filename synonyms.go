package main

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
