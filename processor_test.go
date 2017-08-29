package main

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
