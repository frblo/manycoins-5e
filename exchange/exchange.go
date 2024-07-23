package exchange

import (
	c "manycoins/coins"
	"strings"
)

func Exchange(purse c.Purse) string {
	baseValue := purseToBase(purse)

	var purses []c.Purse

	purses = append(purses, baseToPurse(baseValue, c.Platinum))

	for p, err := c.NextDenomination(c.Platinum); err == nil; p, err = c.NextDenomination(p) {
		purses = append(purses, baseToPurse(baseValue, p))
	}

	purses = append(purses, baseToPurse(baseValue, c.Copper))

	sb := strings.Builder{}
	for _, p := range purses {
		ps, err := purseString(p)
		if err != nil {
			continue
		}
		sb.WriteString(ps)
	}

	return sb.String()
}

// Calculates the combined value of a purse in copper
func purseToBase(purse c.Purse) int {
	sum := 0
	for _, v := range purse.Coins {
		sum += v.Amout * v.Type.ExchangeRate
	}

	return sum
}

func baseToCoin(base int, purse *c.Purse, p c.Piece) int {
	amount, rem := modi(base, p.ExchangeRate)
	if amount != 0 {
		purse.Coins = append(purse.Coins, c.Coinstack{Type: p, Amout: amount})
	}
	return rem
}

func baseToPurse(base int, maxPiece c.Piece) c.Purse {
	purse := c.Purse{}

	if maxPiece.Denomination == c.Copper.Denomination {
		purse.Coins = append(purse.Coins, c.Coinstack{Type: c.Copper, Amout: base})
	}

	rem := baseToCoin(base, &purse, maxPiece)
	if rem == 0 && rem == base {
		return purse
	}
	base = rem

	for p, err := c.NextDenomination(maxPiece); err == nil; p, err = c.NextDenomination(p) {
		rem := baseToCoin(base, &purse, p)
		if rem == 0 {
			break
		}
		base = rem
	}

	return purse
}
