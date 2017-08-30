package internal

import (
	"context"
	"fmt"
)

// StatsRecord creates a bunch of stats for a given input window
type StatsRecord struct {
	// Window len to initialize the window
	windowLen int
	// Window to keep all the InputRecords
	window []InputRecord
	// Get maximum entry from the current window
	max float64
	// Get min entry from the current window
	min float64
	// Get sum of all the entries from the current window
	sum float64
	//
	timestamp int
	//
	count int
	//priceRatio
	priceRatio float64
}

func (st *StatsRecord) String() string {

	return fmt.Sprintf("%10.10d %10.5f %5d %10.5f %10.5f %10.5f",
		st.timestamp,
		st.priceRatio,
		st.count,
		st.sum,
		st.min,
		st.max)
}

// filter will filter out everything lesser than timestamp `ts`
func (st *StatsRecord) filter(input InputRecord) {
	//===============================================
	tempSlice := st.window[:0]
	for _, v := range st.window {
		if v.timestamp > (input.timestamp - st.windowLen) {
			tempSlice = append(tempSlice, v)
		}
	}

	//===============================================
	st.window = tempSlice[:]
	st.sum = 0
	st.count = len(st.window)
	st.priceRatio = input.priceRatio
	st.timestamp = input.timestamp

	st.min = input.priceRatio
	st.max = input.priceRatio

	//================================================

	for i := 0; i < len(st.window); i++ {
		v := st.window[i].priceRatio
		if v < st.min {
			st.min = v
		} else if v > st.max {
			st.max = v
		}
		st.sum = st.sum + v
	}
	//===============================================

}

// Update will append any new entry in the window and generate stats on it.
func (st *StatsRecord) Update(ctx context.Context, input <-chan InputRecord, printStats bool) {

	for {
		select {
		case record := <-input:
			st.window = append(st.window, record)
			st.filter(record)
			if printStats {
				fmt.Println(st)
			}
		case <-ctx.Done():
			return
		}
	}
}

// NewStatsRecord ...
func NewStatsRecord() *StatsRecord {

	st := StatsRecord{windowLen: 60}
	st.window = make([]InputRecord, 0)
	return &st
}

// NewStatsRecordWithLen ...
func NewStatsRecordWithLen(windowLen int) *StatsRecord {
	st := StatsRecord{windowLen: windowLen}
	return &st
}
