package parser

import (
	c "manycoins/coins"
	"testing"
)

func TestParseDenomination(t *testing.T) {
	str := "5pp 3gp 1sp 193ep 2gp 0cp"

	expected := []c.Coinstack{
		{Type: c.Platinum, Amount: 5},
		{Type: c.Gold, Amount: 5},
		{Type: c.Electrum, Amount: 193},
		{Type: c.Silver, Amount: 1},
		{Type: c.Copper, Amount: 0},
	}

	got := make([]c.Coinstack, len(expected))

	for i := 0; i < len(got); i++ {
		got[i] = parseDenomination(str, c.Pieces[i])
	}

	for i := 0; i < len(got); i++ {
		if got[i] != expected[i] {
			t.Errorf("got %v, wanted %v", got[i], expected[i])
		}
	}
}
