package main

import (
	"clisynth/synth"
)

func main() {
	synth.InitSynthInstance(44100, 2048)
	synth.StartKBSession(200)
}
