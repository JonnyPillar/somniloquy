provider "aws" {
  shared_credentials_file = "$HOME/.aws/somniloquy/credentials"
  profile                 = "default"
  region                  = "${var.aws_region}"
}
