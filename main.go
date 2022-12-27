package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func pausePrompt() {
	fmt.Println("Press the Enter Key to continue")
	fmt.Scanln() // wait for Enter Key
}

func cmda() {
	fmt.Println("> Command a chosen")

	cmd := exec.Command("ls", "./")

	// omit the error assignment
	out, _ := cmd.Output()
	fmt.Println("")
	fmt.Println("Your command's output:")
	fmt.Println(string(out))

	pausePrompt()
}

func cmdb() {
	fmt.Println("> Command b chosen")

	cmd := exec.Command("df", "-h")

	// omit the error assignment
	out, _ := cmd.Output()
	fmt.Println("")
	fmt.Println("Your command's output:")
	fmt.Println(string(out))

	pausePrompt()
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("\x1bc")
		// Clear screen from bottom
		fmt.Println("")
		fmt.Println("")
		fmt.Println("--------------------")
		fmt.Println("PC TOOLS - MAIN MENU")
		fmt.Println("--------------------")
		fmt.Println("")
		fmt.Println("Type a for ls")
		fmt.Println("Type b for command df -h")
		fmt.Println("")
		fmt.Print("or press Ctrl-C to quit")
		fmt.Println("")
		fmt.Println("")
		fmt.Print("Enter choice and press Enter: ")

		// Scans from Stdin(Console)
		scanner.Scan()
		// Holds the string that was scanned
		text := scanner.Text()

		fmt.Println("")

		switch {
		case text == "a":
			cmda()
		case text == "b":
			cmdb()

		}

	}

}
