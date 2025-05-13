package commands

import (
	"log"

	"lesiw.io/cmdio/sys"
)

func (Ops) Prune() {
	var rnr = sys.Runner()

	defer rnr.Close()

	// Prune unused Docker images
	err := rnr.Run("docker", "image", "prune", "-f")
	if err != nil {
		log.Fatal(err)
	}
}