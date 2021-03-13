package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func sendViewedMessage(videoPath string) error {
	requestBody, err := json.Marshal(
		map[string]string{
			"videoPath": videoPath,
		},
	)
	if err != nil {
		return fmt.Errorf("unable to construct request body: %w", err)
	}

	resp, err := http.Post("http://history/viewed", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("error sending message to history service: %w", err)
	}
	defer resp.Body.Close()

	return nil
}
