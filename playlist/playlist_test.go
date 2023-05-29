package playlist

import (
	"reflect"
	"testing"
)

func TestPlaylist_Play(t *testing.T) {
	playlist := NewPlaylist(3)

	// Test adding songs to an empty playlist
	playlist.Play("abc")
	expected := []string{"abc"}
	if !reflect.DeepEqual(playlist.GetSongs(), expected) {
		t.Errorf("Expected playlist songs: %v, but got: %v", expected, playlist.GetSongs())
	}

	playlist.Play("def")
	expected = []string{"def", "abc"}
	if !reflect.DeepEqual(playlist.GetSongs(), expected) {
		t.Errorf("Expected playlist songs: %v, but got: %v", expected, playlist.GetSongs())
	}

	playlist.Play("ghi")
	expected = []string{"ghi", "def", "abc"}
	if !reflect.DeepEqual(playlist.GetSongs(), expected) {
		t.Errorf("Expected playlist songs: %v, but got: %v", expected, playlist.GetSongs())
	}

	// Test adding a song that already exists in the playlist
	playlist.Play("abc")
	expected = []string{"abc", "ghi", "def"}
	if !reflect.DeepEqual(playlist.GetSongs(), expected) {
		t.Errorf("Expected playlist songs: %v, but got: %v", expected, playlist.GetSongs())
	}

	// Test adding songs beyond the maximum capacity
	playlist.Play("jkl")
	expected = []string{"jkl", "abc", "ghi"}
	if !reflect.DeepEqual(playlist.GetSongs(), expected) {
		t.Errorf("Expected playlist songs: %v, but got: %v", expected, playlist.GetSongs())
	}

	playlist.Play("mno")
	expected = []string{"mno", "jkl", "abc"}
	if !reflect.DeepEqual(playlist.GetSongs(), expected) {
		t.Errorf("Expected playlist songs: %v, but got: %v", expected, playlist.GetSongs())
	}
}
