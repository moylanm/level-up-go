package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"
)

const path = "entries.json"

// raffleEntry is the struct we unmarshal raffle entries into
type raffleEntry struct {
	ID		int	   `json:"id,string"`
	Name	string `json:"name"`
}

// importData reads the raffle entries from file and creates the entries slice.
func importData() []raffleEntry {
	jsonData, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	entries := []raffleEntry{}

	err = json.Unmarshal(jsonData, &entries)
	if err != nil {
		log.Fatal(err)
	}

	return entries
}

// getWinner returns a random winner from a slice of raffle entries.
func getWinner(entries []raffleEntry) raffleEntry {
	rand.Seed(time.Now().Unix())
	wi := rand.Intn(len(entries))
	return entries[wi]
}

func main() {
	entries := importData()
	log.Println("And... the raffle winning entry is...")
	winner := getWinner(entries)
	time.Sleep(500 * time.Millisecond)
	log.Println(winner)
}
