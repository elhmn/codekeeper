# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Ckp < Formula
  desc ""
  homepage "https://github.com/elhmn/ckp"
  version "0.2.22"
  bottle :unneeded

  if OS.mac? && Hardware::CPU.intel?
    url "https://github.com/elhmn/ckp/releases/download/v0.2.22/ckp_0.2.22_darwin_amd64.tar.gz"
    sha256 "3c71ee52560c04a16e1122f46f8b6ac9e2ee37618dff96725d5ffcd650f381b4"
  end
  if OS.mac? && Hardware::CPU.arm?
    url "https://github.com/elhmn/ckp/releases/download/v0.2.22/ckp_0.2.22_darwin_arm64.tar.gz"
    sha256 "b185bcf5ec6021202a4307e97e11f691215e98a58bd3bbe5294e31433ae04b98"
  end
  if OS.linux? && Hardware::CPU.intel?
    url "https://github.com/elhmn/ckp/releases/download/v0.2.22/ckp_0.2.22_linux_amd64.tar.gz"
    sha256 "5373cc9199c0749e9a0cce990c9730cd105e039de64963567d6e59ad34fc388f"
  end
  if OS.linux? && Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
    url "https://github.com/elhmn/ckp/releases/download/v0.2.22/ckp_0.2.22_linux_arm64.tar.gz"
    sha256 "d8518b951456ed5c9012fe25e07ea0e01dfbeee4444c39fa46320f1b42ff8945"
  end

  def install
    bin.install "ckp"
  end
end
