package main

import (
	"fmt"

	"github.com/adammck/memory"
)

func main() {
	printLimit()
	printUsage()
}

func printLimit() {
	limit, err := memory.Limit()
	if err != nil {
		fmt.Println("error reading limit:", err)
		return
	}

	if memory.IsNoLimit(limit) {
		fmt.Println("limit: none")
		return
	}

	fmt.Printf("limit: %dB\n", limit)
}

func printUsage() {
	usage, err := memory.Usage()
	if err != nil {
		fmt.Println("error reading usage:", err)
		return
	}

	fmt.Printf("usage: %dB\n", usage)
}
