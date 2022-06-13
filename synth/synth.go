package synth

import (
	"clisynth/keyboard"
	"clisynth/osc"
	"clisynth/speaker"
	"clisynth/streamer"
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type synth struct {
	sr         int
	bufferSize int
	oscs       []osc.Osc
}

var synthInstance *synth

func InitSynthInstance(sr int, bufferSize int) {
	if synthInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if synthInstance == nil {
			fmt.Println("Starting synth.")
			synthInstance = &synth{}
		} else {
			fmt.Println("Synth already created.")
			return
		}
	} else {
		fmt.Println("Synth already created.")
		return
	}

	synthInstance.sr = sr
	synthInstance.bufferSize = bufferSize
	synthInstance.oscs = []osc.Osc{osc.NewOsc(sr, 0, "sine", 200, 0)}
	speaker.Init(synthInstance.sr, synthInstance.bufferSize)
	Play()
}

func Mute() {
	fmt.Println(len(synthInstance.oscs))
	speaker.Clear()
}

func Play() {
	streamers := getStreamers()
	speaker.Clear()
	speaker.Play(streamers[:]...)
}

func StartKBSession(centerNote float64) {
	k := keyboard.New(centerNote)
	for !k.ExitSig {
		select {
		case pressedKeys := <-k.PressedKeys:
			var streamers []streamer.Streamer = nil
			for _, osc := range synthInstance.oscs {
				for _, freq := range pressedKeys {
					if freq != 0 {
						streamers = append(streamers, osc.CreateStreamer(freq))
					}
				}
			}
			speaker.Clear()
			speaker.Play(streamers[:]...)
		default:
		}
	}
	Play()
}

func getStreamers() []streamer.Streamer {
	var streamers []streamer.Streamer
	for _, o := range synthInstance.oscs {
		streamers = append(streamers, o.CreateStreamer(1))
	}
	return streamers
}

func NewOsc(waveTypeString string, frequency float64, volume float64) {
	synthInstance.oscs = append(synthInstance.oscs, osc.NewOsc(
		synthInstance.sr,
		len(synthInstance.oscs),
		waveTypeString,
		frequency,
		volume,
	))
	Play()
}
