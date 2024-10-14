package main

import (
	"errors"
	"flag"
	"log"
)

type stack []string

func (s stack) Push(v string) {
	s = append(s, v)
}

func (s stack) Pop() (string, error) {
	length := len(s)

	if length == 0 {
		return "", errors.New("Empty stack")
	}

	result := s[length - 1]
	s = s[:length - 1]

	return result, nil
}

func (s stack) Peek() (string, error) {
	length := len(s)

	if length == 0 {
		return "", errors.New("Empty stack")
	}

	return s[length - 1], nil
}

func (s stack) Length() int {
	return len(s)
}

// isBalanced returns whether the given expression
// has balanced brackets.
func isBalanced(expr string) bool {
	panic("NOT IMPLEMENTED")
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
