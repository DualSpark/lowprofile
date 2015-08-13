package lowprofile

import (
	. "github.com/DualSpark/lowprofile/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/DualSpark/lowprofile/Godeps/_workspace/src/github.com/onsi/gomega"
	"os"
)

var _ = Describe("Util", func() {
	var (

	)

	BeforeEach(func() {
    Debug = true
  })

	Context("When the debug flag is set", func() {
    It("should write out debug statements", func() {
        Expect(func(){Debugln("testing")}).ShouldNot(Panic())
        Expect(func(){Debugf("testing %s", "f")}).ShouldNot(Panic())
    })
  })

	Context("When the profile variable is requested", func(){
		It("should return the value if it exists", func() {
			os.Clearenv()
			os.Setenv(ProfileVariable, "Some-Profile")
			Expect(ActiveProfile()).Should(Equal("Some-Profile"))
		})

		It("should not return the value if it does not exist", func() {
			os.Clearenv()
			Expect(ActiveProfile()).Should(Equal(""))
		})
	})

	AfterEach(func() {
    Debug = false
  })
})
