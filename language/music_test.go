package language

import (
	"testing"
)

func TestSearchMusic(t *testing.T) {
	sentence := "Play In Your Eyes from The weeknd on Spotify"
	music, artist := "In Your Eyes", "The weeknd"
	foundMusic, foundArtist := SearchMusic("en", sentence)

	if music != foundMusic || artist != foundArtist {
		t.Errorf("SearchMusic() failed, excepted %s, %s got %s, %s", music, artist, foundMusic, foundArtist)
	}
}
