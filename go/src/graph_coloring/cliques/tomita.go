package cliques

import (
	"graph_coloring/graph"
)

// Similar to BK except choose u P U X highest number of neigh in  P
// v in P \ N(u)
type TomitaAlgo struct{}

func (ta TomitaAlgo) FindAllMaximalCliques(graph *graph.G) Cliques {
	p := NewBitSet(graph.NumOfVertices)
	for idx := uint32(0); idx < graph.NumOfVertices; idx++ {
		p.Set(idx)
	}

	return *tomitaMaximalClique(make(Clique, 0, graph.NumOfVertices),
		p,
		NewBitSet(graph.NumOfVertices),
		neighborsBitSet(graph),
		&Cliques{Cliques: []Clique{}, NumOfVertices: graph.NumOfVertices})
}

func tomitaMaximalClique(r Clique,
	p *BitSet,
	x *BitSet,
	vertexToEdgeBitSet []*BitSet,
	result *Cliques) *Cliques {

	if p.IsZero() && x.IsZero() {
		result.Add(r.Clone())
		return result
	}

	numOfVertices := uint32(len(vertexToEdgeBitSet))
	pCopy := NewBitSet(numOfVertices)
	xCopy := NewBitSet(numOfVertices)

	for v := uint32(0); v < numOfVertices; v++ {
		if !p.IsSet(v) {
			continue
		}

		neighbors := vertexToEdgeBitSet[v]
		And(pCopy, neighbors, p)
		And(xCopy, neighbors, x)

		tomitaMaximalClique(append(r, v), pCopy, xCopy, vertexToEdgeBitSet, result)

		p.UnSet(v)
		x.Set(v)
	}
	return result
}
