variable "aws_region" {
  description = "The AWS region things are created in"
  default     = "eu-west-1"
}

variable "az_count" {
  description = "Number of AZs to cover in a given region"
  default     = "2"
}

variable "app_image" {
  description = "Docker image to run in the ECS cluster"
  default     = "jonnypillar/somniloquy-services"
}

variable "app_port" {
  description = "Port exposed by the docker image to redirect traffic to"
  default     = 7777
}

variable "app_count" {
  description = "Number of docker containers to run"
  default     = 1
}

variable "ecs_autoscale_role" {
  description = "Role arn for the ecsAutocaleRole"
  default     = "arn:aws:iam::466126549927:role/ecsAutoscaleRoles"
}

variable "ecs_task_execution_role" {
  description = "Role arn for the ecsTaskExecutionRole"
  default     = "arn:aws:iam::466126549927:role/ecsTaskExecutionRole"
}

variable "health_check_path" {
  default = "/"
}

variable "fargate_cpu" {
  description = "Fargate instance CPU units to provision (1 vCPU = 1024 CPU units)"
  default     = "256"
}

variable "fargate_memory" {
  description = "Fargate instance memory to provision (in MiB)"
  default     = "512"
}


variable "audio_upload_bucket_name" {
  description = "The name of the Audio Upload S3 bucket"
  default     = "somniloquy-uploads"
}
