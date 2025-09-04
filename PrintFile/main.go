package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Check if the filename is being indicated in the terminal arguments
	if len(os.Args) < 2 {
		fmt.Println("Error: File name required")
		os.Exit(1)
	}

	// Try to read the file
	fn := os.Args[1]
	file, err := os.Open(fn)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Print the content of the file to the terminal
	io.Copy(os.Stdout, file)
}
