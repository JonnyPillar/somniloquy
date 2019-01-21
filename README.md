# Somniloquy

The Somniloquy project is a tool to record, transcribe and analyse peoples [Somniloquy](https://en.wikipedia.org/wiki/Somniloquy) or sleep talking as its better known.

## Development Setup

- Install PortAudio
  - `brew install portaudio`
- Install Dependancies
  - `make install`
- Run Services
  - `make run-services`
- Run Tests
  - `make test`

## Docker

The project uses Docker to containerize the client and service applications and manage environment dependancies

## Docker Installation

- Install Docker
  - `brew install docker`
- Build Services
  - `make build-services`
- Run Services Container
  `docker container run -p 7777:7777 jonnypillar/somniloquy-services`

## Terraform

The project uses Terraform to be able to startup/tear down infrastructure in AWS.

### Terraform Installation

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