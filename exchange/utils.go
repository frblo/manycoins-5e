package exchange

import (
	"errors"
	"fmt"
	c "manycoins/coins"
	"strings"
)

func modi(x int, y int) (int, int) {
	res := x / y
	rem := x % y

	return res, rem
}

func purseString(purse c.Purse) (string, error) {
	if len(purse.Coins) == 0 {
		return "", errors.New("empty coin purse")
	}

	sb := strings.Builder{}

	for i, p := range purse.Coins {
		if p.Amount == 0 {
			continue
		}

		s := fmt.Sprintf("%v%v", p.Amount, p.Type.Denomination)
		if _, err := sb.WriteString(s); err != nil {
			continue
		}
		if i+1 < len(purse.Coins) {
			if _, err := sb.WriteRune(' '); err != nil {
				continue
			}
		}
	}

	return sb.String(), nil
}
