package playlist

type Playlist struct {
	songs     []string
	maxLength int
}

func NewPlaylist(maxLength int) *Playlist {
	return &Playlist{
		songs:     make([]string, 0),
		maxLength: maxLength,
	}
}

func (p *Playlist) Play(song string) {
	// Check if the song already exists in the playlist
	for i, s := range p.songs {
		if s == song {
			// Move the existing song to the top of the playlist
			p.songs = append(p.songs[:i], p.songs[i+1:]...)
			break
		}
	}

	// Add the song to the top of the playlist
	p.songs = append([]string{song}, p.songs...)

	// Remove the last song if the playlist exceeds its maximum capacity
	if len(p.songs) > p.maxLength {
		p.songs = p.songs[:p.maxLength]
	}
}

func (p *Playlist) GetSongs() []string {
	return p.songs
}
