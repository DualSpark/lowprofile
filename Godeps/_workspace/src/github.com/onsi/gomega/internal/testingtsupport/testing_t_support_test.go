package testingtsupport_test

import (
	. "github.com/DualSpark/lowprofile/Godeps/_workspace/src/github.com/onsi/gomega"

	"testing"
)

func TestTestingT(t *testing.T) {
	RegisterTestingT(t)
	Ω(true).Should(BeTrue())
}
