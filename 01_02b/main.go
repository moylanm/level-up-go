package main

import (
	"log"
	"time"
	"strings"
)

const delay = 700 * time.Millisecond

// print outputs a message and then sleeps for a pre-determined amount
func print(msg string) {
	log.Println(msg)
	time.Sleep(delay)
}

// slowDown takes the given string and repeats its characters
// according to their index in the string.
func slowDown(msg string) {
	words := strings.Split(msg, " ")

	for _, word := range words {
		wordLength := len(word)
		letters := make([]string, (wordLength * (wordLength + 1)) / 2)

		for i := 0; i < wordLength; i++ {
			letters = append(letters, strings.Repeat(string(word[i]), i + 1))
		}

		print(strings.Join(letters, ""))
	}
}

func main() {
	msg := "Time to learn about Go strings!"
	slowDown(msg)
}
