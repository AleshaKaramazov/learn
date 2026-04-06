package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Grep struct {
	pattern    string
	searchFile *os.File
	ignoreCase bool
	count      bool
	lineNumber bool
}

func (grep Grep) run() error {
	var (
		input   *os.File
		scanner *bufio.Scanner
	)

	if grep.searchFile == nil {
		input = os.Stdin
	} else {
		input = grep.searchFile
		defer input.Close()
	}

	if grep.ignoreCase {
		grep.pattern = strings.ToLower(grep.pattern)
	}

	scanner = bufio.NewScanner(input)
	line_number := 1
	count := 0
	for scanner.Scan() {
		text := scanner.Text()
		if grep.ignoreCase {
			text = strings.ToLower(text)
		}
		if strings.Contains(text, grep.pattern) {
			if grep.count {
				count++
			} else {
				if grep.lineNumber {
					fmt.Printf("%d: ", line_number)
				}
				fmt.Printf("%s\n", text)
			}
		}
		line_number++
	}
	if grep.count {
		fmt.Println(count)
	}
	return scanner.Err()
}

func NewGrep(args []string) (Grep, error) {
	if len(args) < 2 {
		return Grep{}, errors.New("empty args")
	}

	grep := Grep{
		pattern:    "",
		searchFile: nil,
		ignoreCase: false,
		count:      false,
		lineNumber: false,
	}

	for _, arg := range args[1:] {
		if arg[0] == '-' {
			if len(arg) == 1 {
				grep.searchFile = os.Stdin	
			} else if arg[1] == '-' {
				switch arg[2:] {
				case "ignore-case":
					grep.ignoreCase = true
				case "count":
					grep.count = true
				case "line-number":
					grep.lineNumber = true
				default:
					return grep, fmt.Errorf("unknown flag: %s\n", arg)
				}
			} else {
				for _, ch := range arg[1:] {
					switch ch {
					case 'i':
						grep.ignoreCase = true
					case 'c':
						grep.count = true
					case 'n':
						grep.lineNumber = true
					default:
						return grep, fmt.Errorf("unknown flag: %s\n", ch)
					}
				}
			}

		} else if grep.pattern == "" {
			grep.pattern = arg
		} else if file, err := os.Open(arg); err != nil {
			return grep, fmt.Errorf("error with open file: %s\n", arg)
		} else {
			grep.searchFile = file
		}
	}

	if grep.pattern == "" {
		return grep, errors.New("pattern not found")
	}
	return grep, nil
}
