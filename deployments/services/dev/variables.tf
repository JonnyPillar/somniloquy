variable "aws_region" {
  description = "The AWS region things are created in"
  default     = "eu-west-1"
}

variable "audio_upload_bucket_name" {
  description = "The name of the Audio Upload S3 bucket"
  default     = "dev-somniloquy-uploads"
}
