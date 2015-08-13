package prompt

import (
	"fmt"
	"errors"
	"github.com/DualSpark/lowprofile/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/DualSpark/lowprofile/Godeps/_workspace/src/gopkg.in/mattes/go-expand-tilde.v1"
	"github.com/DualSpark/lowprofile/lib"
	"os"
)

func BeforeActivatePrompt(c *cli.Context) error {
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

func ActivatePrompt(c *cli.Context) {
	profile := "$(lowprofile pp)"

	fmt.Println("adding profile to prompt")

	var filename, err = lowprofile.CheckForShell()
	var promptVariable = lowprofile.Prompts()[os.Getenv("SHELL")]
	var promptValue = os.Getenv(promptVariable)
	lowprofile.Debugf("current %s is %s", promptVariable, promptValue)
	var prompt = fmt.Sprintf("\"%s%s \"", promptValue, profile)

	filename, err = tilde.Expand(filename)
	if err != nil {
		panic(err)
	}
	found, lines := lowprofile.ScanFileForVariable(filename, promptVariable, prompt)
	if !found {
		lines = append(lines, fmt.Sprintf("export %s=%s", promptVariable, prompt))
	}

	lowprofile.WriteFile(filename, lines)
}
