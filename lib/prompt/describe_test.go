package prompt

import (
	. "github.com/DualSpark/lowprofile/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/DualSpark/lowprofile/Godeps/_workspace/src/github.com/onsi/gomega"
	"github.com/DualSpark/lowprofile/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/DualSpark/lowprofile/lib"
	"os"
)

var _ = Describe("DescribePrompt", func() {
	var (
		context *cli.Context
	)

	BeforeEach(func() {
		os.Clearenv()
  })

	Context("When the profile variable is set", func() {
    It("should describe a prompt with the value", func() {
				os.Setenv(lowprofile.ProfileVariable, "Some-Profile")
				Expect(buildPrompt()).Should(Equal("Some-Profile"))
    })

		It("should not error", func() {
				os.Setenv(lowprofile.ProfileVariable, "Some-Profile")
				Expect(func(){DescribePrompt(context)}).ShouldNot(Panic())
    })
  })

	AfterEach(func() {

  })
})
