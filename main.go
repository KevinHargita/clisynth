package main

import (
	"clisynth/cli"
	"clisynth/synth"
)

func main() {
	synth.InitSynthInstance(44100, 2048)
	cli.CliListen()
}
