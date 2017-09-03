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
