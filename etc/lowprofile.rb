require "language/go"

class Lowprofile < Formula
  desc "Simple profile management for AWS."
  homepage "https://github.com/DualSpark/lowprofile"
  url "https://github.com/DualSpark/lowprofile/archive/9bb404054ddd743a37470c9e1a477d7c24bf3620.tar.gz"
  version "0.2"
  sha256 "64f83dd6282f371a55515d21b0eab853c9691f3c403949748fb38d3af604ed07"
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

    rm "#{HOMEBREW_PREFIX}/etc/lowprofile"
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
