# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Ckp < Formula
  desc ""
  homepage "https://github.com/elhmn/ckp"
  version "0.2.23"
  bottle :unneeded

  if OS.mac? && Hardware::CPU.intel?
    url "https://github.com/elhmn/ckp/releases/download/v0.2.23/ckp_0.2.23_darwin_amd64.tar.gz"
    sha256 "778787ac2da549d52175d691cd15aa9713a7664d9dd1dd3f5ff85f21b58ead82"
  end
  if OS.mac? && Hardware::CPU.arm?
    url "https://github.com/elhmn/ckp/releases/download/v0.2.23/ckp_0.2.23_darwin_arm64.tar.gz"
    sha256 "56e5ac08cf7faee580d6281a90d77ba45b2bf937346ef269d7d2de471656f86e"
  end
  if OS.linux? && Hardware::CPU.intel?
    url "https://github.com/elhmn/ckp/releases/download/v0.2.23/ckp_0.2.23_linux_amd64.tar.gz"
    sha256 "dcee730683cb55dfe1869266826fc4d5d30eac2f89fcc1695f7cb13a92774093"
  end
  if OS.linux? && Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
    url "https://github.com/elhmn/ckp/releases/download/v0.2.23/ckp_0.2.23_linux_arm64.tar.gz"
    sha256 "8516e994fca4b7b1469eb0cfa814a68e5194d142d28824d9b08ae69acb64efac"
  end

  def install
    bin.install "ckp"
  end
end