package main

import (
	"flag"
	"log"
)

type operatorType int

const (
	openBracket operatorType = iota
	closedBracket
	otherOperator
)

var bracketPairs = map[rune]rune {
	'(': ')',
	'[': ']',
	'{': '}',
}

func getOperatorType(op rune) operatorType {
	for ob, cb := range bracketPairs {
		switch op {
		case ob:
			return openBracket
		case cb:
			return closedBracket
		}
	}

	return otherOperator
}

// isBalanced returns whether the given expression
// has balanced brackets.
func isBalanced(expr string) bool {
	return false
}

// printResult prints whether the expression is balanced.
func printResult(expr string, balanced bool){ 
	if balanced {
		log.Printf("%s is balanced.\n", expr)
		return
	}
	log.Printf("%s is not balanced.\n", expr)
}

func main() {
	expr := flag.String("expr", "", "The expression to validate brackets on.")
	flag.Parse()
	printResult(*expr, isBalanced(*expr))
}
