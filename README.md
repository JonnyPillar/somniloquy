# Somniloquy

The Somniloquy project is a tool to record, transcribe and analyse peoples [Somniloquy](https://en.wikipedia.org/wiki/Somniloquy) or sleep talking as its better known.

## Development Setup

- Install Dependancies
  - `make install`
- Run Services
  - `make run-services`
- Run Tests
  - `make test`

## Terraform

The project uses Terraform to be able to startup/tear down infrastructure in AWS.

### Installation

- Install Terraform
  - `brew install terraform`
- Set AWS Credentials
  - `mkdir ~/.aws/somniloquy`
  - `nano credentials`
  - Set `aws_access_key_id` & `aws_secret_access_key`

### Commands

- Check infrastructure changes are valid
  - `terraform plan`
- Apply infrastructure changes
  - `terraform apply`
- Tear down infrastructure changes
  - `terraform destroy`