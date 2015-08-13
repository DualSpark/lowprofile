package lowprofile

import (
	"fmt"
	"errors"
	"github.com/DualSpark/lowprofile/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/DualSpark/lowprofile/Godeps/_workspace/src/gopkg.in/mattes/go-expand-tilde.v1"
	"os"
)

func BeforeActivateProfile(c *cli.Context) error {
	if !c.IsSet("profile") {
		cli.ShowSubcommandHelp(c)
		output := "Missing required option 'profile'"
		fmt.Println(output)
		return errors.New(output)
	}

	shell := os.Getenv("SHELL")
	if Shells()[shell] == "" {
		output := fmt.Sprintf("Sorry, %s is not a supported shell", shell)
		fmt.Println(output)
		return errors.New(output)
	}

	var filename, _ = tilde.Expand(Shells()[shell])
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		output := fmt.Sprintf("File %s not found", Shells()[shell])
		fmt.Println(output)
		return errors.New(output)
	}

	return nil
}

func ActivateProfile(c *cli.Context) {
	profile := c.String("profile")

	fmt.Printf("activating profile %s\n", profile)

	var filename, err = CheckForShell()

	filename, err = tilde.Expand(filename)
	if err != nil {
		panic(err)
	}
	found, lines := ScanFileForVariable(filename, ProfileVariable, profile)
	if !found {
		lines = append(lines, fmt.Sprintf("export %s=%s", ProfileVariable, profile))
	}

	WriteFile(filename, lines)
}
