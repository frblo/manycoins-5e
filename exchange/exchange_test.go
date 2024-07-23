package exchange

import (
	c "manycoins/coins"
	"testing"
)

func TestPurseToBaseFull(t *testing.T) {
	purse := c.Purse{
		Coins: []c.Coinstack{
			{Type: c.Platinum, Amount: 138},
			{Type: c.Gold, Amount: 83},
			{Type: c.Electrum, Amount: 32},
			{Type: c.Silver, Amount: 6},
			{Type: c.Copper, Amount: 3},
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
		{Coins: []c.Coinstack{{Type: c.Platinum, Amount: 1}}},
		{Coins: []c.Coinstack{{Type: c.Gold, Amount: 32}}},
		{Coins: []c.Coinstack{{Type: c.Electrum, Amount: 566}}},
		{Coins: []c.Coinstack{{Type: c.Silver, Amount: 587293}}},
		{Coins: []c.Coinstack{{Type: c.Copper, Amount: 5}}},
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
		{Coins: []c.Coinstack{{Type: c.Platinum, Amount: 0}}},
		{Coins: []c.Coinstack{{Type: c.Gold, Amount: 0}}},
		{Coins: []c.Coinstack{{Type: c.Electrum, Amount: 0}}},
		{Coins: []c.Coinstack{{Type: c.Silver, Amount: 0}}},
		{Coins: []c.Coinstack{{Type: c.Copper, Amount: 0}}},
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

func TestBaseToPurse(t *testing.T) {
	base := 445382

	expected := []c.Purse{
		{
			Coins: []c.Coinstack{
				{Type: c.Platinum, Amount: 445},
				{Type: c.Gold, Amount: 3},
				{Type: c.Electrum, Amount: 1},
				{Type: c.Silver, Amount: 3},
				{Type: c.Copper, Amount: 2},
			},
		},
		{
			Coins: []c.Coinstack{
				{Type: c.Gold, Amount: 4453},
				{Type: c.Electrum, Amount: 1},
				{Type: c.Silver, Amount: 3},
				{Type: c.Copper, Amount: 2},
			},
		},
		{
			Coins: []c.Coinstack{
				{Type: c.Electrum, Amount: 8907},
				{Type: c.Silver, Amount: 3},
				{Type: c.Copper, Amount: 2},
			},
		},
		{
			Coins: []c.Coinstack{
				{Type: c.Silver, Amount: 44538},
				{Type: c.Copper, Amount: 2},
			},
		},
		{
			Coins: []c.Coinstack{
				{Type: c.Copper, Amount: 445382},
			},
		},
	}

	got := make([]c.Purse, len(expected))
	for i := 0; i < len(got); i++ {
		got[i] = baseToPurse(base, c.Pieces[i])
	}

	for i := 0; i < len(got); i++ {
		for j := 0; j < len(got[i].Coins); j++ {
			if got[i].Coins[j] != expected[i].Coins[j] {
				t.Errorf("baseToPurse does not rebase purse correctly. For purse with max piece %v the coin stack is %v when it should be %v",
					c.Pieces[i], got[i].Coins[j], expected[i].Coins[j])
			}
		}
	}
}
