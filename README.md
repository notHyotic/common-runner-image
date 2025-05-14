# common runner image
A simple github actions runner image alternative.

Image definition: [Dockerfile](Dockerfile)

Image link: https://hub.docker.com/repository/docker/hy0tic/common-runner-image

## Tools Installed
- [Alpine](https://github.com/alpinelinux/alpine-make-vm-image)
- [Go](https://github.com/golang/go)
- [golangci-lint](https://github.com/golangci/golangci-lint)
- [bash](https://github.com/bminor/bash)
- [git](https://github.com/git/git)
- [docker-cli](https://github.com/docker/cli)
- [op](https://github.com/lesiw/ops)
- [node](https://github.com/nodejs/node)
- [npm](https://github.com/npm/cli)

## Usage
To use this image in your Github Actions runner, simply specify the image in the job:
```
jobs:
    build-and-upload:
        runs-on: ubuntu-latest  # this is still required
        container: 
          image: hy0tic/common-runner-image
```
full example [here](.github/workflows/build-and-upload.yml)