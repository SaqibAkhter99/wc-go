package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	funcFlag := os.Args[1]
	if len(os.Args) < 3 {
		file, err := os.Open(os.Args[1])
		fName := strings.Split(os.Args[1], "/")
		AbsfileName := fName[len(fName)-1]
		if err != nil {
			log.Fatalf("Failed to open file: %s", err)
			os.Exit(1)
		} else {
			lineCount := count_lines(file)
			defer file.Close()
			file1, _ := os.Open(os.Args[1])
			wordCount := count_words(file1)
			defer file.Close()
			file, _ := os.Open(os.Args[1])
			byteCount := count_bytes(file)
			defer file.Close()
			fmt.Printf("%d %d %d %s \n", lineCount, wordCount, byteCount, AbsfileName)

		}
	} else {
		resFlag := string(funcFlag[1])
		fName := strings.Split(os.Args[2], "/")
		AbsfileName := fName[len(fName)-1]
		file, err := os.Open(os.Args[2])
		if err != nil {
			log.Fatalf("Failed to open file: %s", err)
		}
		if resFlag == "c" {
			byteCount := count_bytes(file)
			defer file.Close()
			fmt.Printf("%d %s \n", byteCount, AbsfileName)
		} else if resFlag == "l" {
			lineCount := count_lines(file)
			defer file.Close()
			fmt.Printf("%d %s \n", lineCount, AbsfileName)
		} else if resFlag == "w" {
			wordCount := count_words(file)
			defer file.Close()
			fmt.Printf("%d %s \n", wordCount, AbsfileName)
		} else if resFlag == "m" {
			charCount := count_chars(file)
			defer file.Close()
			fmt.Printf("%d %s \n", charCount, AbsfileName)

		} else {
			fmt.Println("invalid flag")
		}
	}
}

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
}

func count_lines(file *os.File) int {
	lineCount := 0
	reader := bufio.NewReader(file)
	// Read the file line by line
	for {
		_, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		lineCount++
	}
	return lineCount
}

func count_words(file *os.File) int {
	wordCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wordCount += len(strings.Fields(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return wordCount
}

func count_chars(file *os.File) int {
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("failed to read file: %v", err)
	}

	// Count the characters
	charCount := utf8.RuneCount(content)
	return charCount
}
