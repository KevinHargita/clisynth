package main

import (
	"clisynth/synth"
	"time"
)

func main() {
	synth.InitSynthInstance(44100, time.Second/20)
	synth.StartKBSession(200)
}
