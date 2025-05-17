package commands

import (
	"log"

	"lesiw.io/cmdio/sys"
)

func (Ops) Build() {
	var rnr = sys.Runner()
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
}
