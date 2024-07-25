package exchange

import (
	"log"
	c "manycoins/coins"
	"strings"
)

func Exchange(purse c.Purse, allowedPieces []c.Piece) string {
	if len(allowedPieces) == 0 {
		allowedPieces = c.Pieces
	}

	baseValue := purseToBase(purse)

	var purses []c.Purse

	for i, p := range allowedPieces {
		rebasedPurse := baseToPurse(baseValue, p, allowedPieces)
		if len(rebasedPurse.Coins) == 0 {
			continue
		}
		if rebasedPurse.Coins[0].Type != allowedPieces[i] {
			continue
		}

		purses = append(purses, rebasedPurse)
	}

	sb := strings.Builder{}
	for _, p := range purses {
		ps, err := PurseString(p)
		if err != nil {
			continue
		}
		sb.WriteString(ps)
		sb.WriteRune('\n')
	}

	return sb.String()
}

// Calculates the combined value of a purse in copper
func purseToBase(purse c.Purse) int {
	sum := 0
	for _, v := range purse.Coins {
		sum += v.Amount * v.Type.ExchangeRate
	}

	return sum
}

func baseToPurse(base int, maxPiece c.Piece, pieces []c.Piece) c.Purse {
	purse := c.Purse{}

	maxIndex, err := c.DenominationIndex(maxPiece, pieces)
	if err != nil {
		log.Println(err)
	}

	for i := maxIndex; i < len(pieces); i++ {
		amount, rem := modi(base, pieces[i].ExchangeRate)
		if amount != 0 {
			purse.Coins = append(purse.Coins, c.Coinstack{Type: pieces[i], Amount: amount})
		}
		if rem == 0 {
			break
		}

		base = rem
	}

	return purse
}
