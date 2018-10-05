package main

// hello 

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func execInput(input string) error {
	// Remove the newline character.
	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")
	switch args[0] {
	case "":
		return nil
	case "cd":
		if len(args) < 2 {
			return errors.New("Path Required")
		}
		err := os.Chdir(args[1])
		if err != nil {
			return err
		}
		return nil
	case "exit":
		os.Exit(0)
	}

	// Prepare the command to execute.
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and save it's output.
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
func checkGit() {
	cmd := "git branch | grep \\*"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Print(err)
	}
	text := string(out)
	printGitArrow(text)
}

func printGitArrow(text string) {
	text = strings.Replace(text, "\n", "", -1)
	BackArrowColor := color.New(color.FgBlue, color.BgYellow)
	BackArrowColor.Print("")
	DircColour := color.New(color.BgYellow, color.FgBlack)
	DircColour.Printf(" " + text + " ")
	FrontArrowColour := color.New(color.FgYellow)
	FrontArrowColour.Print(" ")
}
func printDictArrow(text string) {
	text = strings.Replace(text, "\n", "", -1)
	DircColour := color.New(color.BgBlue, color.FgBlack)
	DircColour.Printf(" " + text + " ")
	out, _ := exec.Command("git", "rev-parse", "--is-inside-work-tree").Output()
	b, _ := strconv.ParseBool(strings.Replace(string(out), "\n", "", -1))
	if b {
		checkGit()
	} else {
		FrontArrowColour := color.New(color.FgBlue)
		FrontArrowColour.Print(" ")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		out, err := exec.Command("pwd", ".").Output()
		if err != nil {
			fmt.Printf("%s", err)
		}
		s := string(out)
		printDictArrow(s)
		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		// Handle the execution of the input.
		err = execInput(input)
		if err != nil {
			fmt.Println(err)
		}
	}
}
