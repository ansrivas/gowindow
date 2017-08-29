package main

// func Test_filter(t *testing.T) {
// 	assert := assert.New(t)
// 	actual := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	expected := []int{6, 7, 8, 9}
// 	var expectedWin []InputRecord
// 	for _, v := range expected {
// 		rec := InputRecord{timestamp: v}
// 		expectedWin = append(expectedWin, rec)
// 	}
//
// 	windowLen := 4
// 	st := NewStatsRecordWithLen(windowLen)
// 	for _, v := range actual {
// 		rec := InputRecord{timestamp: v}
// 		st.Update(rec)
// 		st.filter(rec)
// 	}
// 	assert.Equal(st.window, expectedWin, "Filter should always maintain a list of given length")
// }
//
// func Test_Update(t *testing.T) {
// 	assert := assert.New(t)
// 	st := NewStatsRecordWithLen(4)
// 	rec1 := InputRecord{timestamp: 123, priceRatio: 32.1}
// 	rec2 := InputRecord{timestamp: 123, priceRatio: 32.2}
//
// 	st.Update(rec1)
// 	assert.Equal(st.window, []InputRecord{rec1}, "Update should insert values in place")
//
// 	st.Update(rec2)
// 	assert.Equal(st.window, []InputRecord{rec1, rec2}, "Update should insert values in place")
// }
