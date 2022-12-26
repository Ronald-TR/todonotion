require "formula"

class Todonotion < Formula
    desc "A todonotion CLI to improve productivity"
    homepage "https://github.com/Ronald-TR/todonotion"
    url "https://github.com/Ronald-TR/todonotion/releases/download/v0.1.0/todonotion-mac-arm64.tar.gz"
    sha256 "6d49c7975fadde47f4a00d5749ce1bf1c3dbe3d1d8d33be310c8e54cc1041a24"
    head "git@github.com:Ronald-TR/todonotion.git"
  
    def install
      bin.install "todonotion"
    end
  
    # Homebrew requires tests.
    test do
      assert_match "v0.2.0", shell_output("#{bin}/todonotion version", 2)
    end
  end
