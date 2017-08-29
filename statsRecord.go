package main

import (
	"context"
	"fmt"
	"math"
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
	if len(tempSlice) == 0 {
		st.min = math.MaxFloat64
		st.max = math.SmallestNonzeroFloat64
	} else {
		st.min = tempSlice[0].priceRatio
		st.max = tempSlice[0].priceRatio
	}
	//================================================
	for idx, v := range tempSlice {
		if v.priceRatio < st.min {
			st.min = v.priceRatio
		}
		if v.priceRatio > st.max {
			st.max = v.priceRatio
		}
		st.sum = st.sum + v.priceRatio
		st.window[idx] = v
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
	st.window = make([]InputRecord, 60)
	return &st
}

// NewStatsRecordWithLen ...
func NewStatsRecordWithLen(windowLen int) *StatsRecord {
	st := StatsRecord{windowLen: windowLen}
	return &st
}
