package graph_coloring

import "fmt"

type color int

const (
	UNSET color = iota
)

type domain struct {
	vertexColor []color
}

func (d *domain) IsSet(vertex uint32) bool {
	return d.vertexColor[vertex] != UNSET
}

func (d *domain) Set(vertex uint32, color color) error {
	if d.vertexColor[vertex] != UNSET {
		return fmt.Errorf("Trying to set a color when vertex is already colored.")
	}
	d.vertexColor[vertex] = color
	return nil
}

func (d *domain) Color(vertex uint32) color {
	return d.vertexColor[vertex]
}

func newDomain(numVertices uint32) *domain {
	return &domain{vertexColor: make([]color, numVertices)}
}

