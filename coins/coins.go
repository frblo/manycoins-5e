package coins

import (
	"errors"
	"log"
	"slices"
)

type Piece struct {
	Denomination string
	ExchangeRate int // based on copper
}

var (
	Platinum = Piece{"pp", 1000}
	Gold     = Piece{"gp", 100}
	Electrum = Piece{"ep", 50}
	Silver   = Piece{"sp", 10}
	Copper   = Piece{"cp", 1}
	Pieces   = []Piece{Platinum, Gold, Electrum, Silver, Copper}
)

func DenominationIndex(piece Piece, pieces []Piece) (int, error) {
	index := slices.Index(pieces, piece)
	if index != -1 {
		return index, nil
	}

	copperIndex := slices.Index(Pieces, Copper)
	if copperIndex != -1 {
		return copperIndex, errors.New("could not find %v in given array, defaulting to copper")
	}
	log.Fatal("copper not in list of valid pieces")
	return -1, errors.New("copper not in list of valid pieces")
}

type Coinstack struct {
	Type   Piece
	Amount int
}

type Purse struct {
	Coins []Coinstack
}
