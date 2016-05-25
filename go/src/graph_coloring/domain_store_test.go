package graph_coloring

import "testing"

func TestDomainVertex_SetAColor(t *testing.T) {
	d := newDomain(2)

	verifyVertexColor(t, d, 0, UNSET)
	verifyVertexColor(t, d, 1, UNSET)

	d.Set(0, 1)

	verifyVertexColor(t, d, 0, 1)
	verifyVertexColor(t, d, 1, UNSET)
}

func TestDomain_Set_AlreadySetColorReturnsAnError(t *testing.T) {
	d := newDomain(2)

	if err := d.Set(0, 1); err != nil {
		t.Errorf("Expected a successs but was an error %v", err)
	}

	if err := d.Set(0, 2); err == nil {
		t.Errorf("Expected an error but was nil")
	}
}

func verifyVertexColor(t *testing.T, d *domain, vertex uint32, expectedColor color) {
	if d.IsSet(vertex) == (expectedColor == UNSET) {
		t.Errorf("Expected IsSet to be true but was false")
	}

	if actualColor := d.Color(vertex); actualColor != expectedColor {
		t.Errorf("Expected color to be %v but was %v", expectedColor, actualColor)
	}
}
