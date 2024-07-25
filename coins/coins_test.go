package coins

import (
	"slices"
	"testing"
)

func TestPiecesArray(t *testing.T) {
	copperIndex := slices.Index(Pieces, Copper)
	if copperIndex == -1 {
		t.Error("copper is not present in the array of valid pieces")
	}
}

func TestDenominationIndex(t *testing.T) {
	gotIndex := make([]int, len(Pieces))
	gotErrors := make([]error, len(Pieces))

	for i, p := range Pieces {
		gotIndex[i], gotErrors[i] = DenominationIndex(p, Pieces)
	}

	for i := 0; i < len(gotIndex); i++ {
		if gotIndex[i] != i {
			t.Errorf("index of %v is %v, not %v", Pieces[i], gotIndex[i], i)
		}
		if gotErrors[i] != nil {
			t.Errorf("%v should not generate error %v", Pieces[i], gotErrors[i])
		}
	}
}

func TestDenominationIndexInvalid(t *testing.T) {
	mud := Piece{"mp", 9999}

	expected := 4

	got, err := DenominationIndex(mud, Pieces)
	if err == nil {
		t.Errorf("mud piece should generate an error")
	}

	if got != expected {
		t.Errorf("mud piece should give position %v, not %v", expected, got)
	}
}
