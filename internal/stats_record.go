// MIT License
//
// Copyright (c) 2017 Ankur Srivastava
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

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
	// timestamp represents the last timestamp in current window
	timestamp int
	// count represents the number of elements in current window
	count int
	//priceRatio is the last record in the current window
	priceRatio float64
}

// String converts a stats record in its formatted string representation
// 1355270609    1.80215     1    1.80215    1.80215    1.80215
func (st *StatsRecord) String() string {

	return fmt.Sprintf("%10.10d %10.5f %5d %10.5f %10.5f %10.5f",
		st.timestamp,
		st.priceRatio,
		st.count,
		st.sum,
		st.min,
		st.max)
}

// filter will filter out everything lesser than timestamp `ts` and populate
// the StatsRecord struct
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

	for i := 0; i < st.count; i++ {
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

// NewStatsRecord creates a new StatsRecord object with default window length of 60
func NewStatsRecord() *StatsRecord {

	st := StatsRecord{windowLen: 60}
	st.window = make([]InputRecord, 0)
	return &st
}

// NewStatsRecordWithLen creates a new StatsRecord object with default window length of windowLen
func NewStatsRecordWithLen(windowLen int) *StatsRecord {
	st := StatsRecord{windowLen: windowLen}
	st.window = make([]InputRecord, 0)
	return &st
}
