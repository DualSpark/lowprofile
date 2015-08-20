require "language/go"

class Lowprofile < Formula
  desc "Simple profile management for AWS."
  homepage "https://github.com/DualSpark/lowprofile"
  url "https://github.com/DualSpark/lowprofile/archive/feature-promptly.tar.gz"
  version "0.2"
  sha256 "3e8082a3bb33145fe70112ab082d4664b457493fd1a40733a7fcee59aeb62678"
  depends_on "go" => :build

  go_resource "github.com/DualSpark/lowprofile" do
    url "https://github.com/DualSpark/lowprofile.git", :branch => "feature-promptly"
  end

  def install
    ENV["GOPATH"] = buildpath
    Language::Go.stage_deps resources, buildpath/"src"
    # Build and install lowprofile
    system "go", "build", "-v", "-o", "./bin/lowprofile-#{version}", "main.go"
    bin.install "bin/lowprofile-#{version}"

    rm "#{HOMEBREW_PREFIX}/etc/lowprofile" if File.exist?("#{HOMEBREW_PREFIX}/etc/lowprofile")
    etc.install "etc/lowprofile"
  end

  def caveats; <<-EOS.undent
    Add the following to your bash_profile or zshrc to complete the install:

      . #{HOMEBREW_PREFIX}/etc/lowprofile

    and source the file to pick up the change.

    if you don't already have it in there feel free to add (if not lowprofile
    will append it for you):

      export AWS_PROFILE=default

    that's it lowprofile with take it from there!

    You can now switch AWS profiles simply by typing

      lowprofile activate-profile --profile new-profile

    EOS
  end

  test do
    system "#{bin}/lowprofile-#{version}", "--help"
  end
end
