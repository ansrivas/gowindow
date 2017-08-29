package main

import (
	"fmt"
	"strings"
)

//printHeader will print a default header
func printHeader() {
	header := fmt.Sprintf("%5s %12s %8s %8s %11s %10s", "T", "V", "N", "RS", "MinV", "MaxV")
	fmt.Println(header, "\n", strings.Repeat("-", len(header)))
}
