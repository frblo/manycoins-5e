package main

import (
	"fmt"
	"log"
	"manycoins/exchange"
	"manycoins/parser"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	sb := strings.Builder{}
	for _, s := range args {
		_, err := sb.WriteString(s)
		if err != nil {
			log.Printf("failed to parse argument %v", s)
		}
	}

	purse := parser.Parse(sb.String())
	exchanged := exchange.Exchange(purse)

	fmt.Println(exchanged)
}
