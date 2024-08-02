package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	cont, err := io.ReadAll(os.Stdin)
	checkPipe := 0
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read from stdin: %s\n", err)
		os.Exit(1)
	}
	checkPipe = 1
	// Define flags
	byteFlag := flag.Bool("c", false, "Count bytes")
	lineFlag := flag.Bool("l", false, "Count lines")
	wordFlag := flag.Bool("w", false, "Count words")
	charFlag := flag.Bool("m", false, "Count characters")
	flag.Parse()
	// Check for file argument
	if flag.NArg() < 1 && checkPipe == 0 {
		log.Fatalf("Usage: %s [-c|-l|-w|-m] <filename>", os.Args[0])
	}
	if checkPipe == 1 {
		if *byteFlag {
			fmt.Printf("%d  \n", len(cont))
		}
		if *lineFlag {
			lineCount := countLines(cont)
			fmt.Printf("%d  \n", lineCount)
		}
		if *wordFlag {
			wordCount := countWords(cont)
			fmt.Printf("%d  \n", wordCount)
		}
		if *charFlag {
			charCount := utf8.RuneCount(cont)
			fmt.Printf("%d \n", charCount)
		}

		// If no specific flag is provided, default to all counts
		if !*byteFlag && !*lineFlag && !*wordFlag && !*charFlag {
			lineCount := countLines(cont)
			wordCount := countWords(cont)
			charCount := utf8.RuneCount(cont)
			fmt.Printf("%d %d %d\n", lineCount, wordCount, charCount)
		}
	} else {

		// Get the file name
		fileName := flag.Arg(0)

		// Open the file
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatalf("Failed to open file: %s", err)
		}
		defer file.Close()

		// Read the file content
		content, err := io.ReadAll(file)
		if err != nil {
			log.Fatalf("Failed to read file: %s", err)
		}

		// Count and print results based on flags
		if *byteFlag {
			fmt.Printf("%d %s\n", len(content), fileName)
		}
		if *lineFlag {
			lineCount := countLines(content)
			fmt.Printf("%d %s\n", lineCount, fileName)
		}
		if *wordFlag {
			wordCount := countWords(content)
			fmt.Printf("%d %s\n", wordCount, fileName)
		}
		if *charFlag {
			charCount := utf8.RuneCount(content)
			fmt.Printf("%d %s\n", charCount, fileName)
		}

		// If no specific flag is provided, default to all counts
		if !*byteFlag && !*lineFlag && !*wordFlag && !*charFlag {
			lineCount := countLines(content)
			wordCount := countWords(content)
			charCount := utf8.RuneCount(content)
			fmt.Printf("%d %d %d %s\n", lineCount, wordCount, charCount, fileName)
		}
	}
}

func countLines(content []byte) int {
	lineCount := 0
	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	for scanner.Scan() {
		lineCount++
	}
	return lineCount
}

func countWords(content []byte) int {
	wordCount := 0
	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	for scanner.Scan() {
		wordCount += len(strings.Fields(scanner.Text()))
	}
	return wordCount
}
