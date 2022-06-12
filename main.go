package main

import (
	"clisynth/cli"
	"clisynth/synth"
	"time"
)

func main() {
	synth.InitSynthInstance(44100, time.Second/20)
	cli.CliListen()
}
