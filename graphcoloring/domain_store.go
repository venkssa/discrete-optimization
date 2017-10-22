package graphcoloring

import "fmt"

type color int

const (
	UNSET color = iota
)

type DomainStore struct {
	vertexColors       []color
	numOfColoredVertex uint32
}

func (d *DomainStore) IsAllVertexColored() bool {
	return d.numOfColoredVertex == uint32(len(d.vertexColors))
}

func (d *DomainStore) IsSet(vertex uint32) bool {
	return d.vertexColors[vertex] != UNSET
}

func (d *DomainStore) Set(vertex uint32, color color) error {
	if d.vertexColors[vertex] != UNSET {
		return fmt.Errorf("Trying to set a color when vertex is already colored.")
	}
	d.numOfColoredVertex++
	d.vertexColors[vertex] = color
	return nil
}

func (d *DomainStore) Color(vertex uint32) color {
	return d.vertexColors[vertex]
}

func NewDomainStore(numVertices uint32) *DomainStore {
	return &DomainStore{vertexColors: make([]color, numVertices)}
}

func MakeACopy(d *DomainStore) *DomainStore {
	colors := make([]color, len(d.vertexColors))
	copy(colors, d.vertexColors)
	return &DomainStore{vertexColors: colors, numOfColoredVertex: d.numOfColoredVertex}
}
