package slack

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func SendWebHook(msg string) {
	url := os.Getenv("SLACK_WEBHOOK_URL")
	message := map[string]interface{}{
		"text": msg,
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	http.Post(url, "application/json", bytes.NewBuffer(bytesRepresentation))
}
