package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"
)

const path = "songs.json"

// Song stores all the song related information
type Song struct {
	Name      string `json:"name"`
	Album     string `json:"album"`
	PlayCount int64  `json:"play_count"`
}

// makePlaylist makes the merged sorted list of songs
func makePlaylist(albums [][]Song) []Song {
	songs := make([]Song, 0)

	for _, album := range albums {
		for _, song := range album {
			songs = append(songs, song)
		}
	}

	sort.Slice(songs, func(i, j int) bool {
		return songs[i].PlayCount > songs[j].PlayCount
	})

	return songs
}

func main() {
	albums := importData()
	printTable(makePlaylist(albums))
}

// printTable prints merged playlist as a table
func printTable(songs []Song) {
	w := tabwriter.NewWriter(os.Stdout, 3, 3, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "####\tSong\tAlbum\tPlay count")
	for i, s := range songs {
		fmt.Fprintf(w, "[%d]:\t%s\t%s\t%d\n", i+1, s.Name, s.Album, s.PlayCount)
	}
	w.Flush()

}

// importData reads the input data from file and creates the friends map
func importData() [][]Song {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data [][]Song
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
