package coins

import "errors"

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

func DenominationIndex(piece Piece) (int, error) {
	switch piece {
	case Platinum:
		return 0, nil
	case Gold:
		return 1, nil
	case Electrum:
		return 2, nil
	case Silver:
		return 3, nil
	case Copper:
		return 4, nil
	default:
		return 4, errors.New("not a valid denomination")
	}
}

type Coinstack struct {
	Type   Piece
	Amount int
}

type Purse struct {
	Coins []Coinstack
}
