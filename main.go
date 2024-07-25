package main

import (
	"flag"
	"fmt"
	"log"
	c "manycoins/coins"
	"manycoins/exchange"
	"manycoins/parser"
	"os"
	"strings"
)

func main() {
	include := flag.Bool("i", false, "includes only the specified pieces when rebasing the purse")
	flag.Parse()

	args := os.Args[1:]

	sb := strings.Builder{}
	for _, s := range args {
		_, err := sb.WriteString(s)
		if err != nil {
			log.Printf("failed to parse argument %v", s)
		}
	}

	purse := parser.Parse(sb.String())
	purseString, err := exchange.PurseString(purse)
	if err != nil {
		log.Panic(err)
	}
	fmt.Printf("rebasing purse %v...\n", purseString)

	var exchanged string
	if *include {
		includedPieces := flag.Args()
		allowedPieces := parser.ParseIncludes(includedPieces)
		allowedPiecesString := strings.Builder{}
		allowedPiecesString.WriteString("only using ")
		for _, p := range allowedPieces {
			allowedPiecesString.WriteString(p.Denomination)
			allowedPiecesString.WriteRune(' ')
		}
		fmt.Println(allowedPiecesString.String())

		exchanged = exchange.Exchange(purse, allowedPieces)
	} else {
		exchanged = exchange.Exchange(purse, c.Pieces)
	}

	fmt.Println(exchanged)
}
