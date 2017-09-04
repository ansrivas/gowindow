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
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_filter(t *testing.T) {
	assert := assert.New(t)

	//=================================================================
	actual := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := []int{6, 7, 8, 9}
	var expectedWin []InputRecord
	for _, v := range expected {
		rec := InputRecord{timestamp: v}
		expectedWin = append(expectedWin, rec)
	}
	//=================================================================
	windowLen := 4
	st := NewStatsRecordWithLen(windowLen)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Lets create an input channel, with some buffer
	inputRecordChan := make(chan InputRecord, 5)

	//=================================================================
	go st.Update(ctx, inputRecordChan, false)

	//=================================================================
	go func() {
		for _, v := range actual {
			rec := InputRecord{timestamp: v}
			inputRecordChan <- rec
		}
	}()
	//=================================================================
	time.Sleep(time.Second * 5)

	assert.Equal(st.window, expectedWin, "Filter should always maintain a list of given length")
}

func Test_NewStatsRecord(t *testing.T) {
	assert := assert.New(t)

	//=================================================================
	// We create a mock events of 70 instances.
	actual := make([]int, 0)
	for i := 0; i < 71; i++ {
		actual = append(actual, i)
	}

	//=================================================================
	// After a full update cycle of 60 events, we expect these 10 events
	// to be remaining as output.
	expected := make([]int, 0)
	for i := 11; i < 71; i++ {
		expected = append(expected, i)
	}

	//=================================================================
	// Now we prepare an expected result of input records
	var expectedWin []InputRecord
	for _, v := range expected {
		rec := InputRecord{timestamp: v}
		expectedWin = append(expectedWin, rec)
	}

	//=================================================================
	// Testing starts here
	st := NewStatsRecord()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Lets create an input channel, with some buffer
	inputRecordChan := make(chan InputRecord, 5)

	//=================================================================
	go st.Update(ctx, inputRecordChan, false)

	//=================================================================
	go func() {
		for _, v := range actual {
			rec := InputRecord{timestamp: v}
			inputRecordChan <- rec
		}
	}()
	//=================================================================
	time.Sleep(time.Second * 5)
	//=================================================================

	assert.Equal(st.window, expectedWin, "Filter should always maintain a list of given length")
}
