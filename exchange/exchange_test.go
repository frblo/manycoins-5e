package exchange

import (
	c "manycoins/coins"
	"testing"
)

func TestPurseToBaseFull(t *testing.T) {
	purse := c.Purse{
		Coins: []c.Coinstack{
			{Type: c.Platinum, Amout: 138},
			{Type: c.Gold, Amout: 83},
			{Type: c.Electrum, Amout: 32},
			{Type: c.Silver, Amout: 6},
			{Type: c.Copper, Amout: 3},
		},
	}

	expected := 147963
	got := purseToBase(purse)

	if got != expected {
		t.Errorf("got %v, wanted %v", got, expected)
	}
}

func TestPurseToBaseOne(t *testing.T) {
	purses := []c.Purse{
		{Coins: []c.Coinstack{{Type: c.Platinum, Amout: 1}}},
		{Coins: []c.Coinstack{{Type: c.Gold, Amout: 32}}},
		{Coins: []c.Coinstack{{Type: c.Electrum, Amout: 566}}},
		{Coins: []c.Coinstack{{Type: c.Silver, Amout: 587293}}},
		{Coins: []c.Coinstack{{Type: c.Copper, Amout: 5}}},
	}

	expectedInts := []int{
		1000,
		3200,
		28300,
		5872930,
		5,
	}

	gotInts := make([]int, len(purses))

	for i := 0; i < len(purses); i++ {
		gotInts[i] = purseToBase(purses[i])
	}

	for i, gotInt := range gotInts {
		if gotInt != expectedInts[i] {
			t.Errorf("got %v as a string, wanted %v", gotInt, expectedInts[i])
		}
	}
}

func TestPurseToBaseZero(t *testing.T) {
	purses := []c.Purse{
		{Coins: []c.Coinstack{{Type: c.Platinum, Amout: 0}}},
		{Coins: []c.Coinstack{{Type: c.Gold, Amout: 0}}},
		{Coins: []c.Coinstack{{Type: c.Electrum, Amout: 0}}},
		{Coins: []c.Coinstack{{Type: c.Silver, Amout: 0}}},
		{Coins: []c.Coinstack{{Type: c.Copper, Amout: 0}}},
	}

	expected := 0

	gotInts := make([]int, len(purses))

	for i := 0; i < len(purses); i++ {
		gotInts[i] = purseToBase(purses[i])
	}

	for _, gotInt := range gotInts {
		if gotInt != expected {
			t.Errorf("got %v as a string, wanted %v", gotInt, expected)
		}
	}
}
