package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Cat struct {
	files []*os.File
	numberLines bool 
	squezeBlanc bool 
	showsEnd bool
	numberNotBlanc bool
}

func newCat(args []string) (Cat, error) {
	result := Cat {
		files: make([]*os.File,0),
		numberLines: false,
		squezeBlanc: false,
		showsEnd:  false,
		numberNotBlanc: false,
	}
	
	for _, arg := range args[1:] {
		if arg[0] == '-' {
			if len(arg) == 1 {
				result.files = append(result.files, nil)
			} else if arg[1] == '-' {
				switch arg[2:] {
				case "line-number": result.numberLines = true
				case "squeze-blanc": result.squezeBlanc = true
				case "show-end": result.showsEnd = true
				case "number-non-blanc": result.numberNotBlanc = true
				default: return result, fmt.Errorf("unknown flag: %s", arg)
				}
			} else {
				for _, ar := range arg[1:] {
					switch ar {
					case 'l': result.numberLines = true
					case 'q': result.squezeBlanc = true
					case 'E': result.showsEnd = true
					case 'b': result.numberNotBlanc = true
					default: return result, fmt.Errorf("unknown flag: %c", ar)
					}
				}
			}
		} else {
			if file, err := os.Open(arg); err != nil {
				fmt.Printf("error with open file(%s): %s\n", arg, err)
			} else {
				result.files = append(result.files, file)		
			}
		}
	}
	return result, nil
}

func (cat Cat) run() error {
	if len(cat.files) == 0 {
		cat.files = append(cat.files, nil)
	}
	for _, file := range cat.files {
		var input *bufio.Scanner	
		needClose := false
		if file == nil {
			input = bufio.NewScanner(os.Stdin)
		} else {
			input = bufio.NewScanner(file)
			needClose = true
		}
		
		
		if cat.numberLines && cat.numberNotBlanc {
			cat.numberLines = false
		}

		linenumber := 1
		lastBlanc := false

		for input.Scan() {
			text := input.Text()
			if (len(text) == 0) {
				if lastBlanc {
					continue
				} else {
					lastBlanc = true
				} 
			} else {lastBlanc = false}

			if (cat.numberLines || (cat.numberNotBlanc && len(text) != 0)) {
				fmt.Printf("%d: ", linenumber)
				linenumber++
			}

			if !cat.showsEnd {
				fmt.Println(text)	
			} else {
				fmt.Printf( "%s$\n", strings.TrimSuffix(text, "\n"))
			}
		}
		if err:= input.Err(); err != nil {
			return err
		}
		if needClose {
			file.Close()
		}
	}
	return nil
}

