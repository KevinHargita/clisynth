package synth

import (
	"clisynth/keyboard"
	"clisynth/osc"
	"clisynth/speaker"
	"fmt"
	"sync"
	"time"

	"github.com/faiface/beep"
)

var lock = &sync.Mutex{}

type synth struct {
	sr         beep.SampleRate
	bufferSize int
	oscs       []osc.Osc
}

var synthInstance *synth

func InitSynthInstance(sr int, bufferSize time.Duration) {
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

	synthInstance.sr = beep.SampleRate(sr)
	synthInstance.bufferSize = synthInstance.sr.N(bufferSize)
	synthInstance.oscs = []osc.Osc{osc.NewOsc(beep.SampleRate(sr), 0, "sine", 400, 0)}
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
	speaker.Play(len(synthInstance.oscs), streamers[:]...)
}

func StartKBSession(centerNote float64) {
	k := keyboard.New(centerNote)
	for !k.ExitSig {
		select {
		case pressedKeys := <-k.PressedKeys:
			var streamers []beep.Streamer = nil
			for _, osc := range synthInstance.oscs {
				for _, freq := range pressedKeys {
					if freq != 0 {
						streamers = append(streamers, osc.CreateStreamer(freq))
					}
				}
			}
			speaker.Clear()
			speaker.Play(len(synthInstance.oscs), streamers[:]...)
		default:
		}
	}
	Play()
}

func getStreamers() []beep.Streamer {
	var streamers []beep.Streamer
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
