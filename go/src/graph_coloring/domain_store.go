package graph_coloring

import "fmt"

type color int

const (
	UNSET color = iota
)

type DomainStore struct {
	vertexColor []color
}

func (d *DomainStore) IsSet(vertex uint32) bool {
	return d.vertexColor[vertex] != UNSET
}

func (d *DomainStore) Set(vertex uint32, color color) error {
	if d.vertexColor[vertex] != UNSET {
		return fmt.Errorf("Trying to set a color when vertex is already colored.")
	}
	d.vertexColor[vertex] = color
	return nil
}

func (d *DomainStore) Color(vertex uint32) color {
	return d.vertexColor[vertex]
}

func newDomainStore(numVertices uint32) *DomainStore {
	return &DomainStore{vertexColor: make([]color, numVertices)}
}

func MakeACopy(d *DomainStore) *DomainStore {
	colors := make([]color, len(d.vertexColor))
	copy(colors, d.vertexColor)
	return &DomainStore{vertexColor: colors}
}