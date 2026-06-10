package main

import (
	"fmt"
	"strings"
)

func Repeat(str string, iterationCount int) string {
	var repeated string
	for counter := 0; counter < iterationCount; counter++ {
		repeated = repeated + str
	}

	return repeated
}

func RepeatWithBuilder(str string, iterationCount int) string {
	var repeated strings.Builder
	for counter := 0; counter < iterationCount; counter++ {
		repeated.WriteString(str)
	}
	return repeated.String()
}

func main() {
	fmt.Println(Repeat("a", 5))
}