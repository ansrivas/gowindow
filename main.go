package main

import (
	"context"
	"flag"

	"github.com/ansrivas/gowindow/internal"
)

// Version represents version of the project
var Version string

// BuildTime represents the buildtime
var BuildTime string

var filePath string

func handleProcessing(filePath string, printStats bool) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// Lets create an input channel, with some buffer
	inputRecordChan := make(chan internal.InputRecord, 5)
	if printStats {
		internal.PrintHeader()
	}

	recordProcessor := internal.NewStatsRecord()
	go recordProcessor.Update(ctx, inputRecordChan, printStats)
	internal.ProcessFile(filePath, inputRecordChan)
	cancel()
}

func main() {
	flag.StringVar(&filePath, "filePath", "", "input file to parse")
	flag.Parse()
	handleProcessing(filePath, true)
}
