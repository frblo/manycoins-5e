package exchange

import (
	"log"
	c "manycoins/coins"
	"strings"
)

func Exchange(purse c.Purse) string {
	baseValue := purseToBase(purse)

	var purses []c.Purse

	for i, p := range c.Pieces {
		rebasedPurse := baseToPurse(baseValue, p)
		if len(rebasedPurse.Coins) == 0 {
			continue
		}
		if rebasedPurse.Coins[0].Type != c.Pieces[i] {
			continue
		}

		purses = append(purses, rebasedPurse)
	}

	sb := strings.Builder{}
	for _, p := range purses {
		ps, err := purseString(p)
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

func baseToPurse(base int, maxPiece c.Piece) c.Purse {
	purse := c.Purse{}

	maxIndex, err := c.DenominationIndex(maxPiece)
	if err != nil {
		log.Println(err)
	}

	for i := maxIndex; i < len(c.Pieces); i++ {
		amount, rem := modi(base, c.Pieces[i].ExchangeRate)
		if amount != 0 {
			purse.Coins = append(purse.Coins, c.Coinstack{Type: c.Pieces[i], Amount: amount})
		}
		if rem == 0 {
			break
		}

		base = rem
	}

	return purse
}
