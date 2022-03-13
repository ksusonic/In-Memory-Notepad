package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var capacity uint
	fmt.Print("Enter the maximum number of notes: ")
	_, err := fmt.Scanf("%d", &capacity)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}

	memory := make([]string, 0, capacity)
	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Print("Enter command and data: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("%s", err)
			return
		}

		line = strings.TrimSuffix(line, "\n")
		command := strings.Split(line, " ")[0]
		if command == "create" {
			data := strings.TrimSpace(strings.TrimPrefix(line, "create"))
			if len(memory) == cap(memory) {
				fmt.Println("[Error] Notepad is full")
			} else {
				if data == "" {
					fmt.Println("[Error] Missing note argument")
				} else {
					memory = append(memory, data)
					fmt.Println("[OK] The note was successfully created")
				}
			}
		} else if command == "list" {
			if len(memory) == 0 {
				fmt.Println("[Info] Notepad is empty")
			} else {
				for i, item := range memory {
					fmt.Println("[Info] " + strconv.Itoa(i+1) + ": " + item)
				}
			}
		} else if command == "update" {
			splittedData := strings.Split(strings.TrimSpace(strings.TrimPrefix(line, "update")), " ")
			if len(splittedData) == 1 && splittedData[0] == "" {
				fmt.Println("[Error] Missing position argument")
				continue
			}
			i, err := strconv.Atoi(splittedData[0])
			if err != nil {
				fmt.Printf("[Error] Invalid position: %s\n", splittedData[0])
				continue
			}
			text := strings.Join(splittedData[1:], " ")
			if text == "" || len(splittedData) < 2 {
				fmt.Println("[Error] Missing note argument")
				continue
			}
			if i <= 0 || len(memory) < i {
				fmt.Println("[Error] There is nothing to update")
				continue
			}
			memory[i-1] = text
			fmt.Printf("[OK] The note at position %d was successfully updated\n", i)
		} else if command == "delete" {
			splittedData := strings.Split(strings.TrimSpace(strings.TrimPrefix(line, "delete")), " ")
			if len(splittedData) == 1 && splittedData[0] == "" {
				fmt.Println("[Error] Missing position argument")
				continue
			}
			i, err := strconv.Atoi(splittedData[0])
			if err != nil {
				fmt.Printf("[Error] Invalid position: %s\n", splittedData[0])
				continue
			}
			if i <= 0 || len(memory) < i {
				fmt.Println("[Error] There is nothing to delete")
				continue
			}
			memory = deleteKey(&memory, i-1, capacity)
			fmt.Printf("[OK] The note at position %d was successfully deleted\n", i)
		} else if command == "clear" {
			memory = make([]string, 0, 5)
			fmt.Println("[OK] All notes were successfully deleted")
		} else if command == "exit" {
			fmt.Println("[Info] Bye!")
			os.Exit(0)
		} else {
			fmt.Println("[Error] Unknown command")
		}
	}

}

func deleteKey(memory *[]string, i int, capacity uint) []string {
	newMemory := make([]string, len(*memory)-1, capacity)
	newIndex := 0
	for j, item := range *memory {
		if j != i {
			newMemory[newIndex] = item
			newIndex++
		}
	}
	return newMemory
}
