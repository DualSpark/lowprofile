package lowprofile

import (
	"fmt"
	"github.com/DualSpark/lowprofile/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/DualSpark/lowprofile/Godeps/_workspace/src/github.com/vaughan0/go-ini"
	"github.com/DualSpark/lowprofile/Godeps/_workspace/src/gopkg.in/mattes/go-expand-tilde.v1"
	"os"
	"errors"
)

const AWS_DEFAULT_PROFILE = "AWS_DEFAULT_PROFILE"
const dot_aws_credentials = "~/.aws/credentials"

func BeforeDescribeProfiles(c *cli.Context) error {
	var filename, _ = tilde.Expand(dot_aws_credentials)
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		Debugln("AWS credentials file isn't there")
		return errors.New(fmt.Sprintf("File %s not found, make sure to run 'aws configure'", filename))
	}
	return nil
}

func DescribeProfiles(c *cli.Context) {
	Debugln("reading profiles from ~/.aws/credentials")
	filename, err := tilde.Expand(dot_aws_credentials)
	if err != nil {
		panic(err)
	}

	profiles := getProfiles(filename)
	for _, profile := range profiles {
		println(profile)
	}
}

func DescribeActiveProfile(c *cli.Context) {
	Debugln("reading variable AWS_DEFAULT_PROFILE")
	profile := os.Getenv(AWS_DEFAULT_PROFILE)
	if len(profile) > 0 {
		fmt.Printf("current profile is %s\n", profile)
	} else {
		fmt.Println("there is currently no active profile")
	}
}

func getProfiles(filename string) []string {
	config, err := ini.LoadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Unable to load %s make sure the file exists and is valid.", filename))
	}
	var profiles []string
	for profile := range config {
		profiles = append(profiles, profile)
	}
	return profiles
}
