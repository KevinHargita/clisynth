package streamer

type Streamer interface {
	Stream(samples [][2]float64) (n int, ok bool)
	Err() error
}
