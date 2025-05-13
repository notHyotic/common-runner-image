package commands

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"lesiw.io/cmdio"
	"lesiw.io/cmdio/sys"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
		log.Println("Skipping loading .env file")
	}
}

func (Ops) Buildandupload() {
	var rnr = sys.Runner().WithEnv(map[string]string{
		"PKGNAME":         "cmdio",
		"DOCKER_USERNAME": os.Getenv("DOCKER_USERNAME"),
		"DOCKER_PASSWORD": os.Getenv("DOCKER_PASSWORD"),
	})
	defer rnr.Close()

	// Build the image
	err := rnr.Run("docker", "build", "-t", "hy0tic/common-runner-image", ".")
	if err != nil {
		log.Fatal(err)
	}

	// Log into docker
	err = cmdio.Pipe(
		rnr.Command("echo", "$DOCKER_PASSWORD"),
		rnr.Command("docker", "login",
			"-u", "$DOCKER_USERNAME", "--password-stdin"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Push the image
	err = rnr.Run("docker", "push", "hy0tic/common-runner-image:latest")
	if err != nil {
		log.Fatal(err)
	}

	// Log the image size
	err = rnr.Run("docker", "images", "hy0tic/common-runner-image",
		"--format", "Image Size: {{.Size}}")
	if err != nil {
		log.Fatal(err)
	}
}
