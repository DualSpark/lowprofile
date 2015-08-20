package prompt

import (
	. "github.com/DualSpark/lowprofile/Godeps/_workspace/src/github.com/onsi/ginkgo"
	. "github.com/DualSpark/lowprofile/Godeps/_workspace/src/github.com/onsi/gomega"
	"flag"
	"github.com/DualSpark/lowprofile/Godeps/_workspace/src/github.com/codegangsta/cli"
	"github.com/DualSpark/lowprofile/lib"
	"io/ioutil"
	"os"
	"strings"
	"errors"
)

var _ = Describe("Activate", func() {
	var (
		context *cli.Context
		set *flag.FlagSet
	)

	BeforeEach(func() {
		os.Clearenv()

		set = flag.NewFlagSet("test", 0)
		set.String("profile", "profile-name", "doc")
		command := cli.Command{Name: "ap"}
		context = cli.NewContext(nil, set, nil)
		context.Command = command
  })

	Context("When the shell is not supported", func() {
    It("should error", func() {
				os.Setenv("SHELL", "not_supported_shell")
				Expect(BeforeActivatePrompt(context)).Should(Equal(errors.New("Sorry, not_supported_shell is not a supported shell")))
    })
  })

	Context("When the resource file doesn't exist", func() {
    It("should error", func() {
				os.Setenv("SHELL", "/bin/bash")
				os.Setenv("HOME", "/tmp")
        Expect(BeforeActivatePrompt(context)).Should(Equal(errors.New("File ~/.bash_profile not found")))
    })
  })

	Context("When the bash resource file exists", func() {
		var (
			profilePath string = "/tmp/.bash_profile"
			shellName string = "/bin/bash"
			homePath string = "/tmp"
		)

		BeforeEach(func(){
			os.Setenv("SHELL", shellName)
			os.Setenv("HOME", homePath)
			os.Setenv(lowprofile.ProfileVariable, "some-profile")

			var bytes []byte
			ioutil.WriteFile(profilePath, bytes, 0660)
		})

    It("should activate", func() {
        Expect(func(){ActivatePrompt(context)}).ShouldNot(Panic())
				contents, _ := ioutil.ReadFile(profilePath)
				Expect(strings.TrimSpace(string(contents))).To(Equal("export lowprofile_pp=\"$(lowprofile pp) \"\nexport PS1=\"${lowprofile_pp}\""))
    })

		It("should reactivate", func() {
				os.Setenv(lowprofile.ProfileVariable, "some-other-profile")

				Expect(func(){ActivatePrompt(context)}).ShouldNot(Panic())
				contents, _ := ioutil.ReadFile(profilePath)
				Expect(strings.TrimSpace(string(contents))).To(Equal("export lowprofile_pp=\"$(lowprofile pp) \"\nexport PS1=\"${lowprofile_pp}\""))
		})

		It("should reactivate after deactivation", func() {
				os.Setenv(lowprofile.ProfileVariable, "some-other-profile")

				Expect(func(){ActivatePrompt(context)}).ShouldNot(Panic())
				Expect(func(){DeactivatePrompt(context)}).ShouldNot(Panic())
				Expect(func(){ActivatePrompt(context)}).ShouldNot(Panic())
				contents, _ := ioutil.ReadFile(profilePath)
				Expect(strings.TrimSpace(string(contents))).To(Equal("export lowprofile_pp=\"$(lowprofile pp) \"\nexport PS1=\"${lowprofile_pp}\""))
		})

		AfterEach(func(){
			os.Remove(profilePath)
		})
  })

	Context("When the zsh resource file exists", func() {
		var (
			profilePath string = "/tmp/.zshrc"
			shellName string = "/bin/zsh"
			homePath string = "/tmp"
		)

		BeforeEach(func(){
			os.Setenv("SHELL", shellName)
			os.Setenv("HOME", homePath)
			os.Setenv(lowprofile.ProfileVariable, "some-profile")

			var bytes []byte
			ioutil.WriteFile(profilePath, bytes, 0660)
		})

		It("should activate", func() {
				Expect(func(){ActivatePrompt(context)}).ShouldNot(Panic())
				contents, _ := ioutil.ReadFile(profilePath)
				Expect(strings.TrimSpace(string(contents))).To(Equal("export lowprofile_pp=\"$(lowprofile pp) \"\nexport PROMPT=\"${lowprofile_pp}\""))
		})

		It("should reactivate", func() {
				os.Setenv(lowprofile.ProfileVariable, "some-other-profile")
				Expect(func(){ActivatePrompt(context)}).ShouldNot(Panic())
				contents, _ := ioutil.ReadFile(profilePath)
				Expect(strings.TrimSpace(string(contents))).To(Equal("export lowprofile_pp=\"$(lowprofile pp) \"\nexport PROMPT=\"${lowprofile_pp}\""))
		})

		It("should reactivate after deactivation", func() {
				os.Setenv(lowprofile.ProfileVariable, "some-other-profile")

				Expect(func(){ActivatePrompt(context)}).ShouldNot(Panic())
				Expect(func(){DeactivatePrompt(context)}).ShouldNot(Panic())
				Expect(func(){ActivatePrompt(context)}).ShouldNot(Panic())
				contents, _ := ioutil.ReadFile(profilePath)
				Expect(strings.TrimSpace(string(contents))).To(Equal("export lowprofile_pp=\"$(lowprofile pp) \"\nexport PROMPT=\"${lowprofile_pp}\""))
		})

		AfterEach(func(){
			os.Remove(profilePath)
		})
	})

	AfterEach(func() {

  })
})
