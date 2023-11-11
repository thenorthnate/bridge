package bridge

import (
	"time"
)

// If you have > 0 active count, where the last addition
// was > X duration ago, open a warning... if greater than
// Y duration ago, open an error.

// If in the above state, the next Done() call resets the
// active count to 0 and subsequent calls to Done() are ignored
// until another Start call is made.

type Link struct {
	start       time.Time
	duration    time.Duration
	err         error
	errDuration time.Duration
	// save the previous X (e.g. 1000) execution times in this
	// array and then compute the quartiles of it periodically
	// and check where the current execution duration is relative
	// to those values (and some cutoff... > 1min different) then
	// alert based on execution times.
	historyCount int
	historyIndex int
	history      []float64
}

func (link *Link) Start() {
	link.start = time.Now().UTC()
}

func (link *Link) Done(err error) {
	link.err = err
	link.duration = time.Since(link.start)
	if err != nil {
		link.errDuration = link.duration
	}
}
