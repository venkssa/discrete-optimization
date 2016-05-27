package graph_coloring

import "testing"

func TestDomainVertex_SetAColor(t *testing.T) {
	d := NewDomainStore(2)

	verifyVertexColor(t, d, 0, UNSET)
	verifyVertexColor(t, d, 1, UNSET)

	d.Set(0, 1)

	verifyVertexColor(t, d, 0, 1)
	verifyVertexColor(t, d, 1, UNSET)
}

func TestDomain_Set_AlreadySetColorReturnsAnError(t *testing.T) {
	d := NewDomainStore(2)

	if err := d.Set(0, 1); err != nil {
		t.Errorf("Expected a successs but was an error %v", err)
	}

	if err := d.Set(0, 2); err == nil {
		t.Errorf("Expected an error but was nil")
	}
}

func TestCopy(t *testing.T) {
	domain := NewDomainStore(2)
	domain.Set(0, 1)

	copiedDomain := MakeACopy(domain)

	if len(copiedDomain.vertexColors) != len(domain.vertexColors) {
		t.Errorf("Expected the copied domain to have the same lenght as domain")
	}

	copiedDomain.Set(1, 1)
	verifyVertexColor(t, domain, 1, UNSET)
	verifyVertexColor(t, copiedDomain, 1, 1)
}

func verifyVertexColor(t *testing.T, d *DomainStore, vertex uint32, expectedColor color) {
	if d.IsSet(vertex) == (expectedColor == UNSET) {
		t.Errorf("Expected IsSet to be true but was false")
	}

	if actualColor := d.Color(vertex); actualColor != expectedColor {
		t.Errorf("Expected color to be %v but was %v", expectedColor, actualColor)
	}
}
