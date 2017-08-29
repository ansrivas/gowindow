package internal

import (
	"fmt"
	"strings"
)

//PrintHeader will print a default header
func PrintHeader() {
	header := fmt.Sprintf("%5s %12s %8s %8s %11s %10s", "T", "V", "N", "RS", "MinV", "MaxV")
	fmt.Println(header, "\n", strings.Repeat("-", len(header)))
}
