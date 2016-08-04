package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Alexander-Attar/jumblesort/worker"
)

func main() {
	inputs := os.Args[1:]
	output := make([]string, len(inputs))
	var o = worker.JumbleSort(inputs, output)
	fmt.Println(strings.Join(o[:], " "))
}
