package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)


type Grep struct {
	pattern string
	searchFile *os.File
	ignoreCase bool
	count bool
	lineNumber bool
}

func (g Grep) run() error {
	var (
		input *os.File;
		scanner *bufio.Scanner;
	)

	if g.searchFile == nil {
		input = os.Stdin
	} else {
		input = g.searchFile;
		defer input.Close()
	}

	if g.ignoreCase {
		g.pattern = strings.ToLower(g.pattern)
	} 

	scanner = bufio.NewScanner(input)
	line_number := 1;
	count := 0;
	for scanner.Scan() {
		text := scanner.Text();
		if g.ignoreCase {
			text = strings.ToLower(text);
		}
		if strings.Contains(text, g.pattern) {
			if g.count {
				count++;	
			} else {
				if g.lineNumber {
					fmt.Printf("%d: ", line_number);
				}
				fmt.Printf("%s\n",  text);
			}
		}
		line_number++;
	}
	if g.count {
		fmt.Println(count);
	}
	return scanner.Err()
}

func New(args []string) (Grep, error) {
	if len(args) < 2 {return Grep{}, errors.New("empty args")}

	g := Grep {
		pattern: "",
		searchFile: nil,
		ignoreCase: false,
		count: false,
		lineNumber: false,
	}

	for _, arg := range args[1:] {
		if arg[0] == '-' {
			if arg[1] == '-' {
				switch arg[2:] {
				case "ignore-case": g.ignoreCase = true;
			 	case "count": g.count = true;
				case "line-number": g.lineNumber = true;
				default: return g, fmt.Errorf("unknown flag: %s\n", arg,)
				}
			} else {
				for _, ch := range arg[1:] {
					switch ch {
					case 'i': g.ignoreCase = true;
					case 'c': g.count = true;
					case 'n': g.lineNumber = true;
					default: return g, fmt.Errorf("unknown flag: %s\n", ch)
					}	
				}	
			}
		
		} else if g.pattern == "" {
			g.pattern = arg;
		} else if file, err := os.Open(arg); err != nil {
			return g, fmt.Errorf("error with open file: %s\n", arg);
		} else {
			g.searchFile = file
		}
	}

	if g.pattern == "" {
		return g, errors.New("pattern not found");
	}
	return g, nil
}

func main() {
	if grep, err := New(os.Args); err != nil {
		fmt.Printf("GREP: %s\n", err);
	} else if err = grep.run(); err!=nil {
		fmt.Printf("GREP: %s\n", err);
	} 
}
