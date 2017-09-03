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
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// InputRecord represents a record in input file
type InputRecord struct {
	timestamp  int
	priceRatio float64
}

// inputParser parses a line from the input file and converts it into an InputRecord
func inputParser(line string) (InputRecord, error) {
	var inpRec InputRecord

	record := strings.Fields(line)
	if len(record) < 2 {
		return inpRec, fmt.Errorf("Unable to parse: %s", line)
	}
	ts, err1 := strconv.Atoi(record[0])
	pr, err2 := strconv.ParseFloat(record[1], 64)

	if (err1 != nil) || (err2 != nil) {
		return inpRec, fmt.Errorf("Unable to parse: %s", line)
	}

	return InputRecord{
		timestamp:  ts,
		priceRatio: pr,
	}, nil

}

// ProcessFile reads a file line by line and publishes to an input channel
func ProcessFile(filePath string, output chan<- InputRecord) {
	fileHandle, err := os.Open(filePath)
	defer fileHandle.Close()
	if err != nil {
		log.Fatal(err)
	}
	fileReader := bufio.NewScanner(fileHandle)

	for fileReader.Scan() {
		line := fileReader.Text()
		record, err := inputParser(line)
		if err != nil {
			log.Printf("Unable to parse the given line: %s", line)
			continue
		}
		output <- record
	}
}
