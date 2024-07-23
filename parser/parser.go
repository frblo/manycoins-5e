package parser

import (
	"fmt"
	c "manycoins/coins"
	"regexp"
	"strconv"
)

func Parse(args string) c.Purse {
	var coins []c.Coinstack

	for _, p := range c.Pieces {
		stack := parseDenomination(args, p)
		if stack.Amount != 0 {
			coins = append(coins, stack)
		}
	}

	return c.Purse{Coins: coins}
}

func parseDenomination(str string, piece c.Piece) c.Coinstack {
	re := regexp.MustCompile(fmt.Sprintf(`\d+\s*%v`, piece.Denomination))
	occurrences := re.FindAll([]byte(str), -1)

	if occurrences == nil {
		return c.Coinstack{}
	}

	sum := 0
	numberRe := regexp.MustCompile(`\d+`)
	for _, o := range occurrences {
		number := numberRe.Find(o)
		amount, err := strconv.Atoi(string(number))
		if err != nil {
			continue
		}
		sum += amount
	}

	return c.Coinstack{Type: piece, Amount: sum}
}
