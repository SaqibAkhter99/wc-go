package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fileName := flag.String("c", "", "Path to the file to read")
	fmt.Println(*fileName)
	flag.Parse()

	// Extracting name of file
	fName := strings.Split(*fileName, "/")
	AbsfileName := fName[len(fName)-1]

	if *fileName == "" {
		fmt.Println("Please provide a file name using the -file flag.")
		os.Exit(1)
	}
	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	byteCount := count_bytes(file)
	defer file.Close()
	fmt.Printf("The file %s contains %d bytes.\n", AbsfileName, byteCount)

}

// Create a buffered reader

func count_bytes(file *os.File) int {
	reader := bufio.NewReader(file)
	// Initialize byte count
	byteCount := 0
	// Read the file byte by byte
	for {
		_, err := reader.ReadByte()
		if err != nil {
			break
		}
		byteCount++
	}
	return byteCount

	// Print the byte count

}
