# Somniloquy

The Somniloquy project is a tool to record, transcribe and analyse peoples [Somniloquy](https://en.wikipedia.org/wiki/Somniloquy) or sleep talking as its better known.

## Prerequisite

This solution uses the [PortAudio](http://portaudio.com/) library for Microphone I/O. This needs to be installed on any machine that is running the Client application.

- OSX
  - `brew install portaudio`
- Linux
  - `apt-get install portaudio19-dev`

## Development Setup

- Install Dependancies
  - `make install`
- Run Tests
  - `make test`
- Run Services
  - `make run-record-service`
  - `make run-conversion-service`
  - `make run-transcription-service`
- Run Client
  - `make run-client`

## Docker

The project uses Docker to containerize the client and service applications and manage environment dependancies

### Docker Installation

- Install Docker
  - [Docker Desktop](https://www.docker.com/products/docker-desktop)
- Install Docker Cli
  - `brew install docker`

### Docker Commands

- Build Services
  - `make build-services`
- Run Services Container
  `docker container run -p 7777:7777 jonnypillar/somniloquy-services`
- Push Services To Docker
  `make push-services`

## Terraform

The project uses Terraform to be able to startup/tear down infrastructure in AWS. We have two environments, dev and live. Run `terraform apply` in the dev dir for development

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