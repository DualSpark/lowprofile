package lowprofile

import (
	"fmt"
	"github.com/DualSpark/lowprofile/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/DualSpark/lowprofile/Godeps/_workspace/src/gopkg.in/mattes/go-expand-tilde.v1"
	"os"
	"errors"
)

func BeforeDeactivateProfile(c *cli.Context) error {
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

func DeactivateProfile(c *cli.Context) {
	Debugln("checking shell")
	shell := os.Getenv("SHELL")
	Debugf("the shell is %s", shell)

	profile := os.Getenv(ProfileVariable)
	if len(profile) > 0 {
		fmt.Printf("deactivating profile %s\n", profile)
	} else {
		fmt.Println("there is currently no active profile")
	}

	var filename, err = CheckForShell()
	filename, err = tilde.Expand(filename)
	if err != nil {
		panic(err)
	}
	found, lines := ScanFileForVariableAndComment(filename, ProfileVariable)
	if found {
		WriteFile(filename, lines)
	}
}
