package cmd

import "time"

// TODO: decide on 1 or more portfolios
// Make multiple portfolios (1 per chain?)
type Portfolio []struct {
	positions []Position
}

type Position struct {
	Ticker     string
	Size       float64
	EntryPrice float64
	EntryTime  time.Time // time.Now().UTC() for instantiated
}
