//package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const filename string = "todo.txt"
const endCommand string = "end"

func loadList(filename string) (result []string, err error) {
	file, err := os.Open(filename)
	if err != nil {
		if err == os.ErrNotExist {
			createdFile, errCreate := os.Create(filename)
			createdFile.Close()
			return []string{}, errCreate
		}
		return []string{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result, scanner.Err()
}

func showList(list []string) {
	for i, text := range list {
		fmt.Printf("%d. %s\n", i+1, text)
	}
}

func nextCommand(buffer *bufio.Scanner, str *string) bool {
	fmt.Print("> ")
	if buffer.Scan() {
		*str = buffer.Text()
		return true
	} else {
		return false
	}
}

func isDelCommand(command string, tasks *[]string) bool {
	if !strings.HasPrefix(command, "del ") {
		return false
	}

	if parts := strings.Split(command, " "); len(parts) > 1 {
		if index, err := strconv.Atoi(strings.TrimSpace(parts[1])); err == nil && index > 0 && index <= len(*tasks) {
			fmt.Println("succes deleted!")
			*tasks = append((*tasks)[:index-1], (*tasks)[index:]...)
		} else {
			fmt.Printf("cant convert to int or use index: %s\n", parts[1])
		}
		return true
	} else {
		fmt.Print("needed format: del [COMMAND NUBER]")
		return true
	}

}

func AddCommand(command string, tasks *[]string) {
	*tasks = append(*tasks, strings.TrimSpace(command))
}

func saveList(text []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, task := range text {
		if _, err := fmt.Fprintf(file, "%s\n", task); err != nil {
			return err
		}
	}
	return nil
}

func isListCommand(command string, text []string) bool {
	if command == "list" {
		showList(text)
		return true
	}
	return false
}

func main() {
	text, _ := loadList(filename)
	scanner := bufio.NewScanner(os.Stdin)
	var command string
	showList(text)
	for nextCommand(scanner, &command) && command != endCommand {
		if !isDelCommand(command, &text) && !isListCommand(command, text) {
			AddCommand(command, &text)
		}
	}

	if err := saveList(text); err != nil {
		fmt.Printf("error with saving text: %s\n", err)
	}

}
