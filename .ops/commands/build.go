package commands

import (
	"log"
	"os"

	// "github.com/joho/godotenv"
	"lesiw.io/cmdio/sys"
)

// func init() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// }

func (Ops) Build() {
	var rnr = sys.Runner().WithEnv(map[string]string{
		"PKGNAME": "cmdio",
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
	err = rnr.Run("sh", "-c", "echo $DOCKER_PASSWORD | docker login -u $DOCKER_USERNAME --password-stdin")
	if err != nil {
		log.Fatal(err)
	}

	// Push the image
	err = rnr.Run("docker", "push", "hy0tic/common-runner-image:latest")
	if err != nil {
		log.Fatal(err)
	}
}
