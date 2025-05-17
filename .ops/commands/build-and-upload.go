package commands

import (
	"os"
	"lesiw.io/cmdio/sys"
)

func (Ops) Buildandupload() {
	var rnr = sys.Runner().WithEnv(map[string]string{
		"PKGNAME":         "cmdio",
		"DOCKER_USERNAME": os.Getenv("DOCKER_USERNAME"),
		"DOCKER_PASSWORD": os.Getenv("DOCKER_PASSWORD"),
	})
	defer rnr.Close()

	Ops{}.Build()
	
	Ops{}.Upload()
}
