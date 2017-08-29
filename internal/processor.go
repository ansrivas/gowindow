package internal

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// InputRecord ...
type InputRecord struct {
	timestamp  int
	priceRatio float64
}

var re *regexp.Regexp

func init() {
	re = regexp.MustCompile(`([0-9]+)([ \t]+)(\d*.?\d*)`)
}

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

// ProcessFile ...
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
