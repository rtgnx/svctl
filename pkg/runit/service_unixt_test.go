package runit

import "testing"

func TestListServiceNames(t *testing.T) {
	names, err := ListServiceNames(DefaultSVDIR)
	if err != nil {
		t.Fatal(err)
	}

	if len(names) == 0 {
		t.Errorf("found %d services: %+v", len(names), names)
	}
}
