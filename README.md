# common runner image
A simple github actions runner image alternative.

Image link: https://hub.docker.com/repository/docker/hy0tic/common-runner-image

## Tools Installed
- Alpine
- Go
- bash
- git
- docker-cli
- op

## Usage
To use this image in your Github Actions runner, simply specify the image in the job:
```
jobs:
    build-and-upload:
        runs-on: ubuntu-latest
        container: 
          image: hy0tic/common-runner-image
```
full example [here](.github/workflows/build-and-upload.yml)