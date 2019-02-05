# Somniloquy

The Somniloquy project is a tool to record, transcribe and analyse peoples [Somniloquy](https://en.wikipedia.org/wiki/Somniloquy) or sleep talking as its better known. We say some highly amusing things in our sleep.

Our subconscious likes to say many funny/embarrassing/disturbing things in our sleep. It would be a shame if we didn't keep a transcribed record of those moments for everyone to enjoy.

This project was inspired by the [PISleepTalk](https://thomaskekeisen.de/en/blog/record-sleeptalk-pisleeptalk/) and is still a work in progress.

## Prerequisite

### Port Audio

This audio stream has been hacked together using a Raspberry PI & USB microphone. As a result we are using the [PortAudio](http://portaudio.com/) library for microphone I/O. This needs to be installed on any PI that is running the Client application.

- OSX
  - `brew install portaudio`
- Linux
  - `apt-get install portaudio19-dev`

### Dep

- Dependency management is done using [Dep](https://golang.github.io/dep/).

- OSX
  - `brew install dep`
- Linux
  - `apt-get install dep`

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

The project uses Terraform to be able to startup/tear down infrastructure in AWS. We have two environments, dev and live. 

If you are running locally and using AWS, run `terraform apply` in the dev dir for development to create the required basic AWS infrastructure.

### Terraform Installation

- Install Terraform
  - `brew install terraform`
- Set AWS Credentials
  - `mkdir ~/.aws/somniloquy`
  - `nano credentials`
  - Set `aws_access_key_id` & `aws_secret_access_key`

### Commands

- Go to the environment folder
  - Dev `cd deployments/services/dev`
  - Live `cd deployments/services/live`
- Check infrastructure changes are valid
  - `terraform plan`
- Apply infrastructure changes
  - `terraform apply`
- Tear down infrastructure changes
  - `terraform destroy`

## TODO

- [ ] Create a sleep talking Gopher. Is it a real Go project if it doesn't?
- [ ] Isolate the PortAudio dependency
- [ ] Add DST Analysis to remove empty recordings
- [ ] Remove Stream->AIFF->Flac step
- [ ] Update Terraform to use Kubernetes instead of AWS Fargate
- [ ] Add website
