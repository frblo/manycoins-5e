package exchange

import (
	c "manycoins/coins"
	"testing"
)

func TestModiZeroByOne(t *testing.T) {
	expectedRes := 0
	expectedRem := 0
	gotRes, gotRem := modi(0, 1)

	if gotRem != expectedRem {
		t.Errorf("got %v as a remainder, wanted %v", gotRem, expectedRem)
	}
	if gotRes != expectedRes {
		t.Errorf("got %v as a result, wanted %v", gotRes, expectedRes)
	}
}

func TestModiOneByOne(t *testing.T) {
	expectedRes := 1
	expectedRem := 0
	gotRes, gotRem := modi(1, 1)

	if gotRem != expectedRem {
		t.Errorf("got %v as a remainder, wanted %v", gotRem, expectedRem)
	}
	if gotRes != expectedRes {
		t.Errorf("got %v as a result, wanted %v", gotRes, expectedRes)
	}
}

func TestModiTenBy3(t *testing.T) {
	expectedRes := 3
	expectedRem := 1
	gotRes, gotRem := modi(10, 3)

	if gotRem != expectedRem {
		t.Errorf("got %v as a remainder, wanted %v", gotRem, expectedRem)
	}
	if gotRes != expectedRes {
		t.Errorf("got %v as a result, wanted %v", gotRes, expectedRes)
	}
}

func TestPurseStringEmpty(t *testing.T) {
	purse := c.Purse{}

	expectedString := ""

	gotString, err := purseString(purse)

	if err == nil {
		t.Errorf("purseString with an emtpy purse should yield an error")
	}

	if gotString != expectedString {
		t.Errorf("got %v as a string, wanted %v", gotString, expectedString)
	}
}

func TestPurseStringOneDenomination(t *testing.T) {
	purses := []c.Purse{
		{Coins: []c.Coinstack{{Type: c.Platinum, Amount: 1}}},
		{Coins: []c.Coinstack{{Type: c.Gold, Amount: 32}}},
		{Coins: []c.Coinstack{{Type: c.Electrum, Amount: 9393399939393}}},
		{Coins: []c.Coinstack{{Type: c.Silver, Amount: 587293}}},
		{Coins: []c.Coinstack{{Type: c.Copper, Amount: 5}}},
	}

	expectedStrings := []string{
		"1pp",
		"32gp",
		"9393399939393ep",
		"587293sp",
		"5cp",
	}

	gotStrings := make([]string, len(purses))
	gotErrors := make([]error, len(purses))

	for i := 0; i < len(purses); i++ {
		gotStrings[i], gotErrors[i] = purseString(purses[i])
	}

	for i, gotString := range gotStrings {
		if gotString != expectedStrings[i] {
			t.Errorf("got %v as a string, wanted %v", gotString, expectedStrings[i])
		}
	}

	for i, gotError := range gotErrors {
		if gotError != nil {
			t.Errorf("purse with %v should not return error %v", purses[i].Coins[0].Type.Denomination, gotError)
		}
	}
}

func TestPurseStringFull(t *testing.T) {
	purse := c.Purse{
		Coins: []c.Coinstack{
			{Type: c.Platinum, Amount: 138},
			{Type: c.Gold, Amount: 83},
			{Type: c.Electrum, Amount: 32},
			{Type: c.Silver, Amount: 6},
			{Type: c.Copper, Amount: 39283},
		},
	}

	expected := "138pp 83gp 32ep 6sp 39283cp"
	got, err := purseString(purse)

	if got != expected {
		t.Errorf("got %v as a string, wanted %v", got, expected)
	}

	if err != nil {
		t.Errorf("purse should not return error")
	}
}

func TestPurseStringSpotty(t *testing.T) {
	purse := c.Purse{
		Coins: []c.Coinstack{
			{Type: c.Platinum, Amount: 138},
			{Type: c.Gold, Amount: 83},
			{Type: c.Copper, Amount: 39283},
		},
	}

	expected := "138pp 83gp 39283cp"
	got, err := purseString(purse)

	if got != expected {
		t.Errorf("got %v as a string, wanted %v", got, expected)
	}

	if err != nil {
		t.Errorf("purse should not return error")
	}
}
