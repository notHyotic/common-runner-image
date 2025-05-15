package commands

import (
	"encoding/json"
	"log"
	"os"

	"lesiw.io/cmdio/sys"
)

type Metadata struct {
	RunID        string `json:"run_id"`
	Repository   string `json:"repository"`
	RunStartedAt string `json:"run_started_at"`
	UpdatedAt    string `json:"updated_at"`
	Status       string `json:"status"`
	Conclusion   string `json:"conclusion"`
}

func (Ops) Exportdata() {
	var rnr = sys.Runner().WithEnv(map[string]string{
		"PKGNAME":        "cmdio",
		"QUEUE_URL":      "https://sqs.us-east-1.amazonaws.com/915898657279/runner-metrics-queue",
		"RUN_ID":         os.Getenv("RUN_ID"),
		"REPOSITORY":     os.Getenv("REPOSITORY"),
		"RUN_STARTED_AT": os.Getenv("RUN_STARTED_AT"),
		"UPDATED_AT":     os.Getenv("UPDATED_AT"),
		"STATUS":         os.Getenv("STATUS"),
		"CONCLUSION":     os.Getenv("CONCLUSION"),
	})
	defer rnr.Close()

	metadata := Metadata{
		RunID:        rnr.Env("RUN_ID"),
		Repository:   rnr.Env("REPOSITORY"),
		RunStartedAt: rnr.Env("RUN_STARTED_AT"),
		UpdatedAt:    rnr.Env("UPDATED_AT"),
		Status:       rnr.Env("STATUS"),
		Conclusion:   rnr.Env("CONCLUSION"),
	}

	jsonStr, err := json.Marshal(metadata)
	if err != nil {
		log.Fatal(err)
	}

	// Send the message to SQS using aws cli
	err = rnr.Run(
		"aws", "sqs", "send-message",
		"--queue-url", rnr.Env("QUEUE_URL"),
		"--message-body", string(jsonStr),
	)
	if err != nil {
		log.Fatal(err)
	}
}
