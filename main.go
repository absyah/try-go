package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"try-go/playlist"
)

func main() {
	// Read the max capacity of the playlist from command line arguments
	var maxLength int
	if len(os.Args) > 1 {
		fmt.Sscanf(os.Args[1], "%d", &maxLength)
	}

	// Create a new playlist
	playlist := playlist.NewPlaylist(maxLength)

	// Read commands from STDIN
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()
		parts := strings.Split(command, " ")
		if len(parts) != 2 {
			fmt.Println("Invalid command format")
			continue
		}

		if parts[0] == "PLAY" {
			song := parts[1]
			playlist.Play(song)
			fmt.Printf("%s Played\n", song)
			fmt.Printf("Playlist: %v\n", playlist.GetSongs())
		} else {
			fmt.Println("Invalid command")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
