// Package speaker implements playback of beep.Streamer values through physical speakers.
package speaker

import (
	"clisynth/mixer"
	"clisynth/streamer"
	"math"
	"sync"

	"github.com/hajimehoshi/oto/v2"
)

var (
	mu      sync.Mutex
	mix     mixer.Mixer
	samples [][2]float64
	context *oto.Context
	done    chan struct{}
)

// Init initializes audio playback through speaker. Must be called before using this package.
//
// The bufferSize argument specifies the number of samples of the speaker's buffer. Bigger
// bufferSize means lower CPU usage and more reliable playback. Lower bufferSize means better
// responsiveness and less delay.
func Init(sampleRate int, bufferSize int) error {
	mu.Lock()
	defer mu.Unlock()

	mix = mixer.Mixer{SampleRate: sampleRate}

	samples = make([][2]float64, bufferSize)

	var err error
	var ready chan struct{}
	context, ready, err = oto.NewContext(sampleRate, 2, 2)
	if err != nil {
		return err
	}
	<-ready

	done = make(chan struct{})

	go func() {
		for {
			select {
			default:
				update()
			case <-done:
				return
			}
		}
	}()

	return nil
}

// Lock locks the speaker. While locked, speaker won't pull new data from the playing Stramers. Lock
// if you want to modify any currently playing Streamers to avoid race conditions.
//
// Always lock speaker for as little time as possible, to avoid playback glitches.
func Lock() {
	mu.Lock()
}

// Unlock unlocks the speaker. Call after modifying any currently playing Streamer.
func Unlock() {
	mu.Unlock()
}

// Play starts playing all provided Streamers through the speaker.
func Play(s ...streamer.Streamer) {
	mu.Lock()
	mix.UpdateChannelCount(len(s))
	mix.Add(s...)
	mu.Unlock()
}

// Clear removes all currently playing Streamers from the speaker.
func Clear() {
	mu.Lock()
	mix.Clear()
	mu.Unlock()
}

// update pulls new data from the playing Streamers and sends it to the speaker. Blocks until the
// data is sent and started playing.
func update() {
	mu.Lock()
	mix.Stream(samples)
	p := context.NewPlayer(&Speaker{samples: &samples})
	p.Play()

	mu.Unlock()
}

type Speaker struct {
	samples *[][2]float64
}

func (s *Speaker) Read(buf []byte) (int, error) {
	newsamples := *s.samples
	for i := 0; i < len(*s.samples); i++ {
		const max = 32767
		b1 := int16(math.Sin(2*math.Pi*newsamples[i][0]) * max)
		b2 := int16(math.Sin(2*math.Pi*newsamples[i][1]) * max)

		buf[4*i] = byte(b1)
		buf[4*i+1] = byte(b1 >> 8)

		buf[4*i+2] = byte(b2)
		buf[4*i+3] = byte(b2 >> 8)

	}

	return len(newsamples), nil
}
