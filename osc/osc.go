package osc

import (
	"fmt"

	"clisynth/waves"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
)

type WaveType int

const (
	sine WaveType = iota
	square
	sawtooth
	triangle
)

var waveTypeMap map[string]WaveType = map[string]WaveType{
	"sine":     sine,
	"square":   square,
	"sawtooth": sawtooth,
	"triangle": triangle,
}

func (wt WaveType) String() string {
	return []string{"sine", "square", "sawtooth", "triangle"}[wt]
}

type Osc struct {
	sr        int
	OscId     int
	waveType  WaveType
	frequency float64
	volume    float64
}

func NewOsc(sr int, oscId int, waveTypeString string, frequency float64, volume float64) Osc {
	waveType := waveTypeMap[waveTypeString]

	newOsc := Osc{
		sr,
		oscId,
		waveType,
		frequency,
		volume,
	}
	return newOsc
}

func (o *Osc) CreateStreamer(freqRatio float64) beep.Streamer {
	var streamer beep.Streamer

	f := freqRatio * o.frequency

	switch o.waveType {
	case sine:
		newStreamer, err := waves.SineTone(o.sr, f)
		if err != nil {
			fmt.Println(err)
		}
		streamer = newStreamer
	case square:
		newStreamer, err := waves.SquareTone(o.sr, f)
		if err != nil {
			fmt.Println(err)
		}
		streamer = newStreamer

	case sawtooth:
		newStreamer, err := waves.SawtoothTone(o.sr, f)
		if err != nil {
			fmt.Println(err)
		}
		streamer = newStreamer

	case triangle:
		newStreamer, err := waves.TriangleTone(o.sr, f)
		if err != nil {
			fmt.Println(err)
		}
		streamer = newStreamer

	default:
		panic("not a valid oscillator type")

	}

	streamer = &effects.Volume{
		Streamer: streamer,
		Base:     2,
		Volume:   o.volume,
		Silent:   false,
	}

	return streamer
}

func (o *Osc) PrintDetails() {
	fmt.Printf("osc-%d\t%s\t%g\t%g\n", o.OscId, o.waveType.String(), o.frequency, o.volume)
}
