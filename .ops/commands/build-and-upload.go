package commands

import (
	"log"
	"os"
	"strings"

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
	var err error

	err = rnr.Run("docker", "build", "-t", "hy0tic/common-runner-image", ".")
	if err != nil {
		log.Fatal(err)
	}

	err = rnr.Run("docker", "images", "hy0tic/common-runner-image", 
		"--format", "Image Size: {{.Size}}")
	if err != nil {
		log.Fatal(err)
	}
	
	err = cmdio.Pipe(
		strings.NewReader(rnr.Env("DOCKER_PASSWORD")),
		rnr.Command("docker", "login",
			"-u", rnr.Env("DOCKER_USERNAME"),
			"--password-stdin",
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = rnr.Run("docker", "push", "hy0tic/common-runner-image:latest")
	if err != nil {
		log.Fatal(err)
	}
}
