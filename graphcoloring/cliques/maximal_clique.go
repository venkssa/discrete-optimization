package cliques

import "github.com/venkssa/discrete-optimization/graphcoloring/graph"

type Clique []uint32

func (c Clique) Clone() Clique {
	return append(Clique{}, c...)
}

type Cliques struct {
	Cliques       []Clique
	NumOfVertices uint32
}

func (cls *Cliques) Add(c ...Clique) {
	cls.Cliques = append(cls.Cliques, c...)
}

type MaximalCliqueFinder interface {
	FindAllMaximalCliques(graph *graph.G) *Cliques
}

func neighborsBitSet(graph *graph.G) []*BitSet {
	neighborsBitSet := make([]*BitSet, graph.NumOfVertices)

	for vertexIdx := uint32(0); vertexIdx < graph.NumOfVertices; vertexIdx++ {
		neighborsBitSet[vertexIdx] = NewBitSet(graph.NumOfVertices)
		for _, neighborIdx := range graph.Neighbors(vertexIdx) {
			neighborsBitSet[vertexIdx].Set(neighborIdx)
		}
	}

	return neighborsBitSet
}
