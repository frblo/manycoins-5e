package coins

import "testing"

func TestDenominationIndex(t *testing.T) {
	gotIndex := make([]int, len(Pieces))
	gotErrors := make([]error, len(Pieces))

	for i, p := range Pieces {
		gotIndex[i], gotErrors[i] = DenominationIndex(p)
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

	got, err := DenominationIndex(mud)
	if err == nil {
		t.Errorf("mud piece should generate an error")
	}

	if got != expected {
		t.Errorf("mud piece should give position %v, not %v", expected, got)
	}
}
