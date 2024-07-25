package parser

import (
	"fmt"
	c "manycoins/coins"
	"regexp"
	"slices"
	"strconv"
)

func Parse(str string) c.Purse {
	var coins []c.Coinstack

	for _, p := range c.Pieces {
		stack := parseDenomination(str, p)
		if stack.Amount != 0 {
			coins = append(coins, stack)
		}
	}

	return c.Purse{Coins: coins}
}

func parseDenomination(str string, piece c.Piece) c.Coinstack {
	re := regexp.MustCompile(fmt.Sprintf(`-?\d+%v`, piece.Denomination))
	occurrences := re.FindAll([]byte(str), -1)

	if occurrences == nil {
		return c.Coinstack{}
	}

	sum := 0
	numberRe := regexp.MustCompile(`-?\d+`)
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

func ParseIncludes(args []string) []c.Piece {
	includePieces := make([]bool, len(c.Pieces))

	// always include lowest denomination
	includePieces[len(includePieces)-1] = true

	for i, p := range c.Pieces {
		if slices.Contains(args, p.Denomination) {
			includePieces[i] = true
		}
	}

	var pieces []c.Piece
	for i, p := range includePieces {
		if p {
			pieces = append(pieces, c.Pieces[i])
		}
	}

	return pieces
}
