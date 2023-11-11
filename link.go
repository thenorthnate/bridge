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
	start    time.Time
	isActive bool
	err      error
	// save the previous X (e.g. 1000) execution times in this
	// array and then compute the quartiles of it periodically
	// and check where the current execution duration is relative
	// to those values (and some cutoff... > 1min different) then
	// alert based on execution times.
	historyCount   int
	historyIndex   int
	successHistory []float64
}

func (link *Link) Start() {
	link.start = time.Now().UTC()
	link.isActive = true
}

func (link *Link) Done(err error) {
	link.isActive = false
	link.err = err
	if err != nil {
		return
	}
	link.successHistory[link.historyIndex] = time.Since(link.start).Seconds()
	link.historyIndex = (link.historyIndex + 1) % len(link.successHistory)
	if link.historyCount < len(link.successHistory) {
		link.historyCount++
	}
}
