package main

import (
	"flag"

	"github.com/ansrivas/gowindow/internal"
)

// Version represents version of the project
var Version string

// BuildTime represents the buildtime
var BuildTime string

var filePath string

func handleProcessing(filePath string, printStats bool) {

	if printStats {
		internal.PrintHeader()
	}

	internal.ProcessFile(filePath, printStats)

}

func main() {
	flag.StringVar(&filePath, "filePath", "", "input file to parse")
	flag.Parse()
	handleProcessing(filePath, true)
}
