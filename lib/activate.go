package lowprofile

import (
	"bufio"
	"fmt"
	"errors"
	"github.com/DualSpark/lowprofile/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/DualSpark/lowprofile/Godeps/_workspace/src/gopkg.in/mattes/go-expand-tilde.v1"
	"os"
	"regexp"
)

func BeforeActivateProfile(c *cli.Context) error {
	if !c.IsSet("profile") {
		cli.ShowSubcommandHelp(c)
		Debugln("Missing profile")
		return errors.New("Missing required option 'profile'")
	}

	shell := os.Getenv("SHELL")
	if Shells()[shell] == "" {
		Debugln("Unsupported shell")
		return errors.New(fmt.Sprintf("Sorry, %s is not a supported shell", shell))
	}

	var filename, _ = tilde.Expand(Shells()[shell])
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		Debugln("Profile file isn't there")
		return errors.New(fmt.Sprintf("File %s not found", Shells()[shell]))
	}

	return nil
}

func ActivateProfile(c *cli.Context) {
	profile := c.String("profile")

	fmt.Printf("activating profile %s\n", profile)

	var filename, err = checkForShell()

	filename, err = tilde.Expand(filename)
	if err != nil {
		panic(err)
	}
	found, lines := scanFileForVariable(filename, profileVariable, profile)
	if !found {
		lines = append(lines, fmt.Sprintf("export %s=%s", profileVariable, profile))
	}

	writeFile(filename, lines)
}

func scanFileForVariable(filename string, variable string, profile string) (bool, []string) {

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	found := false
	regex := regexp.MustCompile(fmt.Sprintf("\\#*\\s*(export\\s+%s=).*", variable))
	replace := fmt.Sprintf("${1}%s", profile)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if regex.MatchString(text) {
			found = true
			text = regex.ReplaceAllString(text, replace)
		}
		lines = append(lines, text)
		Debugln(text)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return found, lines
}
