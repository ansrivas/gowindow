package main

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_filter(t *testing.T) {
	assert := assert.New(t)
	actual := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := []int{6, 7, 8, 9}
	var expectedWin []InputRecord
	for _, v := range expected {
		rec := InputRecord{timestamp: v}
		expectedWin = append(expectedWin, rec)
	}

	windowLen := 4
	st := NewStatsRecordWithLen(windowLen)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Lets create an input channel, with some buffer
	inputRecordChan := make(chan InputRecord, 5)

	//-----------------------------------------------------
	go func() {
		for i := 0; i < len(actual); i++ {
			st.Update(ctx, inputRecordChan, false)
		}
	}()
	//-----------------------------------------------------
	go func() {
		for _, v := range actual {
			rec := InputRecord{timestamp: v}
			inputRecordChan <- rec
		}
	}()
	//-----------------------------------------------------
	time.Sleep(time.Second * 5)

	assert.Equal(st.window, expectedWin, "Filter should always maintain a list of given length")
}
