package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fileName := flag.String("file", "", "Path to the file to read")
	flag.Parse()
	absName = fileName
	fmt.Println(absName)
	if *fileName == "" {
		fmt.Println("Please provide a file name using the -file flag.")
		os.Exit(1)
	}
	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	// Create a buffered reader
	reader := bufio.NewReader(file)
	fmt.Println(reader.ReadLine())
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

	// Print the byte count
	fmt.Printf("The file %s contains %d bytes.\n", *fileName, byteCount)
}
