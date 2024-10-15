package main

import (
	"container/heap"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

const path = "songs.json"

// Song stores all the song related information
type Song struct {
	Name      string `json:"name"`
	Album     string `json:"album"`
	PlayCount int64  `json:"play_count"`

	AlbumCount, SongCount int
}

type PlaylistHeap []Song

func (h PlaylistHeap) Len() int {
	return len(h)
}

func (h PlaylistHeap) Less(i, j int) bool {
	return h[i].PlayCount > h[j].PlayCount
}

func (h PlaylistHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PlaylistHeap) Push(x any) {
	*h = append(*h, x.(Song))
}

func (h *PlaylistHeap) Pop() any {
	original := *h
	n := len(original)
	x := original[n - 1]
	*h = original[0:n - 1]
	return x
}

// makePlaylist makes the merged sorted list of songs
func makePlaylist(albums [][]Song) []Song {
	var playlist []Song
	pheap := &PlaylistHeap{}
	if len(albums) == 0 {
		return playlist
	}

	heap.Init(pheap)
	for i, f := range albums {
		firstSong := f[0]
		firstSong.AlbumCount, firstSong.SongCount = i, 0
		heap.Push(pheap, firstSong)
	}

	for pheap.Len() != 0 {
		p := heap.Pop(pheap)
		song := p.(Song)
		playlist = append(playlist, song)

		if song.SongCount < len(albums[song.AlbumCount]) - 1 {
			nextSong := albums[song.AlbumCount][song.SongCount + 1]
			nextSong.AlbumCount, nextSong.SongCount = song.AlbumCount, song.SongCount + 1
			heap.Push(pheap, nextSong)
		}
	}

	return playlist
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
