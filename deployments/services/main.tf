provider "aws" {
  shared_credentials_file = "$HOME/.aws/Somniloquy/credentials"
  profile                 = "default"
  region                  = "${var.aws_region}"
}
