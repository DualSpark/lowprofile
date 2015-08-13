package prompt

import (
	"fmt"
	"github.com/DualSpark/lowprofile/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/DualSpark/lowprofile/lib"
	"os"
)

func DescribePrompt(c *cli.Context) {
	fmt.Print(buildPrompt())
}

func buildPrompt() string {
	return fmt.Sprintf("%s%s%s", os.Getenv("AWS_PROMPT_PREFIX"), lowprofile.ActiveProfile(), os.Getenv("AWS_PROMPT_SUFFIX"))
}
