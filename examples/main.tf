terraform {
  required_providers {
    yunjidft = {
      source = "yunji/yunjidft"
    }
  }
}

resource "yunjidft_demo" "test" {
  instance_dft= "dft-test"
  disk_size = 100
}