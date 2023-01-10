package discord

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func SendWebHook(msg string) {
	url := os.Getenv("DISCORD_WEBHOOK_URL")
	message := map[string]interface{}{
		"content": msg,
	}

	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}

	http.Post(url, "application/json", bytes.NewBuffer(bytesRepresentation))
}
