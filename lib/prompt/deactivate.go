package prompt

import (
	"fmt"
	"github.com/DualSpark/lowprofile/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/DualSpark/lowprofile/Godeps/_workspace/src/gopkg.in/mattes/go-expand-tilde.v1"
	"github.com/DualSpark/lowprofile/lib"
	"os"
	"errors"
)

func BeforeDeactivatePrompt(c *cli.Context) error {
	shell := os.Getenv("SHELL")
	if lowprofile.Shells()[shell] == "" {
		output := fmt.Sprintf("Sorry, %s is not a supported shell", shell)
		fmt.Println(output)
		return errors.New(output)
	}

	var filename, _ = tilde.Expand(lowprofile.Shells()[shell])
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		output := fmt.Sprintf("File %s not found", lowprofile.Shells()[shell])
		fmt.Println(output)
		return errors.New(output)
	}

	return nil
}

func DeactivatePrompt(c *cli.Context) {
	promptVariable := lowprofile.Prompts()[os.Getenv("SHELL")]

	fmt.Println("removing profile from prompt")

	var filename, err = lowprofile.CheckForShell()
	filename, err = tilde.Expand(filename)
	if err != nil {
		panic(err)
	}
	found, lines := lowprofile.ScanFileForVariableAndComment(filename, promptVariable)
	if found {
		lowprofile.WriteFile(filename, lines)
	}
}
