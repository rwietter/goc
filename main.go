package main

import (
	"errors"
	"fmt"

	"github.com/manifoldco/promptui"

	"rwietter/goc/cmd"
)

func prompt() {
	validate := func(input string) error {
		len := len(input)

		if len == 0 {
			return errors.New("Please enter a path")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Path of your local repository",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
}

func main() {
	cmd.Execute()
}
