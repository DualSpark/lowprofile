package lowprofile

import (
  "os"
  "bufio"
  "regexp"
  "fmt"
  "log"
  "strings"
  "errors"
)

const zsh                 = "/bin/zsh"
const zshrc               = "~/.zshrc"
const zshPrompt         = "PROMPT"
const bash_profile        = "~/.bash_profile"
const bash                = "/bin/bash"
const bashPrompt          = "PS1"
const ProfileVariable     = "AWS_PROFILE"
const dot_aws_credentials = "~/.aws/credentials"

var Debug bool            = false

func Shells() map[string]string {
  return map[string]string {bash: bash_profile, zsh: zshrc}
}

func Prompts() map[string]string {
  return map[string]string {bash: bashPrompt, zsh: zshPrompt}
}

func Debugln(str string) {
  if Debug {
    fmt.Printf("DEBUG: %v\n", str)
  }
}

func Debugf(str string, args ...interface{}) {
  if Debug {
    var format = fmt.Sprintf(str, args...)
    fmt.Printf("DEBUG: %v\n", format)
  }
}

func ActiveProfile() string {
  return os.Getenv(ProfileVariable)
}

func WriteFile(filename string, lines []string) {
  file, err := os.OpenFile(filename, os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()
  Debugf("Writing to file %s", filename)

  w := bufio.NewWriter(file)
  for index := range lines {
    Debugf("Writing: %s", lines[index])
    fmt.Fprintln(w, lines[index])
  }

	w.Flush()
}

func CheckForShell()(string, error) {
	Debugln("checking shell")
	shell := os.Getenv("SHELL")
	Debugf("the shell is %s", shell)
	var err error
	var filename string
	if strings.Contains(shell, zsh) {
		Debugln("checking for variable in ~/.zshrc")
		filename = zshrc
	} else if strings.Contains(shell, bash) {
		Debugln("checking for variable in ~/.bash_profile")
		filename = bash_profile
	} else {
		 err = errors.New(fmt.Sprintf("Sorry, %s is not a supported shell", shell))
	}

	return filename, err
}

func ScanFileForVariable(filename string, variable string, value string) (bool, []string) {

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

  var newlines = LinesInVariable(variable)
	var lines []string
	found := false
	regex := regexp.MustCompile(fmt.Sprintf("\\#*\\s*(export\\s+%s=).*", variable))
	replace := fmt.Sprintf("export %s=%s",variable, value)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		if regex.MatchString(text) {
			found = true
			text = replace
      lines = append(lines, text)
		} else if found && newlines > 1 {
      newlines--
    } else {
      lines = append(lines, text)
    }

		Debugln(text)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return found, lines
}

func LinesInVariable(variable string) int {
	return len(strings.Split(os.Getenv(variable), "\n"))
}

func ScanFileForVariableAndComment(filename string, variable string) (bool, []string) {

  var newlines = LinesInVariable(variable)

  file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	found := false
  regex := regexp.MustCompile(fmt.Sprintf("\\#*\\s*(export\\s+%s=\\w*)", variable))
	replace := "# ${1}"
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
    if found && newlines > 1 {
      text = "# " + text
      newlines--
    }
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
