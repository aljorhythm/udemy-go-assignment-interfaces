package main

import "testing"

func TestTriangleGetArea(t *testing.T) {
	area, e := triangle{2, 2}.getArea(), 2.0
	if area != e {
		t.Errorf("Expected %v got %v", e, area)
	}
}

func TestSquareGetArea(t *testing.T) {
	area, e := square{2}.getArea(), 4.0
	if area != e {
		t.Errorf("Expected %v got %v", e, area)
	}
}
