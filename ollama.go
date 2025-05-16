package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type OllamaResponse struct {
	Response string `json:"response"`
}

func AskOllama(prompt string) (string, error) {
	reqBody := OllamaRequest{
		Model:  "llama3", // zorg dat je dit model lokaal hebt draaien
		Prompt: prompt,
		Stream: false,
	}

	jsonBody, _ := json.Marshal(reqBody)
	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var res OllamaResponse
	json.Unmarshal(body, &res)

	return res.Response, nil
}
