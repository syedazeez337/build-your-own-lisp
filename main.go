package main

import (
	"fmt"
	"os"

	"github.com/peterh/liner"
)

func main() {
	// input := bufio.NewReader(os.Stdin)
	fmt.Println("Lisp version 0.0.1")
	fmt.Println("Press CTRL+c or 'exit'")

	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)
	line.SetCompleter(func(line string) (c []string) {
		if line == "pri" {
			c = append(c, "print", "println")
		}
		return
	})

	for {
		// fmt.Print("lisp>>")
		// message, err := input.ReadString('\n')

		input, err := line.Prompt("lisp>>")
		if err != nil {
			if err == liner.ErrPromptAborted {
				fmt.Println("\nAborted...")
				break
			}
			// panic(err)
			fmt.Fprintln(os.Stderr, "Error reading line: ", err)
		}

		if input == "exit" {
			fmt.Println("Goodbye...")
			break
		}

		fmt.Printf("your input is %s\n", input)
		line.AppendHistory(input)
	}
}
