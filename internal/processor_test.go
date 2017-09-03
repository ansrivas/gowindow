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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_inputParser(t *testing.T) {
	assert := assert.New(t)

	unsafeInputPatterns := []string{"12311.1000", "12311 arc.asas", "ababc 1.111"}
	for _, pattern := range unsafeInputPatterns {
		_, err := inputParser(pattern)
		assert.NotNil(err, fmt.Sprintf("Should fail to parse: %s", pattern))
	}

	safeInputPatterns := []string{"1213123   1.1000", `1213123 1.1000`, `1213123 	1.1000`, `1213123      1.1000`}
	for _, pattern := range safeInputPatterns {
		record, err := inputParser(pattern)
		assert.Nil(err, fmt.Sprintf("Should successfully parse the input pattern %s", pattern))
		assert.Equal(record, InputRecord{timestamp: 1213123, priceRatio: 1.1000}, "msgAndArgs")
	}
}
