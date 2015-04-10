package mathp

import (
	"testing"
)

func TestMaxI(t *testing.T) {
	if MaxI(2, 3) != 3 {
		t.Errorf("MaxI(2, 3) returns %v instead of 3", MaxI(2, 3))
	}
	if MaxI(3, 2) != 3 {
		t.Errorf("MaxI(3, 2) returns %v instead of 3", MaxI(3, 2))
	}
}

func TestMaxI32(t *testing.T) {
	if MaxI32(2, 3) != 3 {
		t.Errorf("MaxI32(2, 3) returns %v instead of 3", MaxI32(2, 3))
	}
	if MaxI32(3, 2) != 3 {
		t.Errorf("MaxI32(3, 2) returns %v instead of 3", MaxI32(3, 2))
	}
}

func TestMaxI64(t *testing.T) {
	if MaxI64(2, 3) != 3 {
		t.Errorf("MaxI64(2, 3) returns %v instead of 3", MaxI64(2, 3))
	}
	if MaxI64(3, 2) != 3 {
		t.Errorf("MaxI64(3, 2) returns %v instead of 3", MaxI64(3, 2))
	}
}

func TestMinI(t *testing.T) {
	if MinI(2, 3) != 2 {
		t.Errorf("MinI(2, 3) returns %v instead of 2", MinI(2, 3))
	}
	if MinI(3, 2) != 2 {
		t.Errorf("MinI(3, 2) returns %v instead of 2", MinI(3, 2))
	}
}

func TestMinI32(t *testing.T) {
	if MinI32(2, 3) != 2 {
		t.Errorf("MinI32(2, 3) returns %v instead of 2", MinI32(2, 3))
	}
	if MinI32(3, 2) != 2 {
		t.Errorf("MinI32(3, 2) returns %v instead of 2", MinI32(3, 2))
	}
}

func TestMinI64(t *testing.T) {
	if MinI64(2, 3) != 2 {
		t.Errorf("MinI64(2, 3) returns %v instead of 2", MinI64(2, 3))
	}
	if MinI64(3, 2) != 2 {
		t.Errorf("MinI64(3, 2) returns %v instead of 2", MinI64(3, 2))
	}
}
