require "formula"

class Rdpcli < Formula
    desc "A todonotion CLI to improve productivity"
    homepage "https://github.com/Ronald-TR/todonotion"
    url "https://github.com/Ronald-TR/todonotion/releases/download/v0.1.0/todonotion-mac-arm64.tar.gz",
    sha256 "066555966e17587246c0bb854f63936335dedf5028dbbc6f27bf9ff5d1e3f6b6"
    head "git@github.com:Ronald-TR/todonotion.git"
  
    def install
      bin.install "todonotion"
    end
  
    # Homebrew requires tests.
    test do
      assert_match "v0.2.0", shell_output("#{bin}/todonotion version", 2)
    end
  end