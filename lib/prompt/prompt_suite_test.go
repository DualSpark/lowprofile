package prompt

import (
	. "github.com/DualSpark/lowprofile/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/DualSpark/lowprofile/Godeps/_workspace/src/github.com/onsi/gomega"
	"testing"
)

func TestLowprofile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lowprofile Prompt Suite")
}
