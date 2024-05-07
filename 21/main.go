package main

import (
	"fmt"
)

type MusicPlayer interface {
	Play(song string)
}

type MP3 struct{}

func (m *MP3) Play(song string) {
	fmt.Printf("MP3 plays song - \"%s\"\n", song)
}

type HiFi struct{}

func (h *HiFi) Music(song string) {
	fmt.Printf("HiFi plays music, song name - \"%s\"\n", song)
}

type HiFiAdapter struct {
	*HiFi
}

func (adapter *HiFiAdapter) Play(song string) {
	adapter.Music(song)
}

func NewHiFiAdapter(hifi *HiFi) MusicPlayer {
	return &HiFiAdapter{hifi}
}

func main() {
	player1 := &MP3{}

	player1.Play("The Beatles - Yesterday")

	player2 := NewHiFiAdapter(&HiFi{})

	player2.Play("The Who - My Generation")
}
