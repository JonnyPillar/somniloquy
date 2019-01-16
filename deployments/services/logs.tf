# Set up cloudwatch group and log stream and retain logs for 30 days
resource "aws_cloudwatch_log_group" "somniloquy_service_log_group" {
  name              = "/ecs/somniloquy-service-app"
  retention_in_days = 30

  tags {
    Name = "somniloquy-service-log-group"
  }
}

resource "aws_cloudwatch_log_stream" "somniloquy_service_log_stream" {
  name           = "somniloquy-service-log-stream"
  log_group_name = "${aws_cloudwatch_log_group.somniloquy_service_log_group.name}"
}