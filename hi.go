package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://api.openai.com/v1/chat/completions"
	payload := map[string]interface{}{
		"prompt":     "Write code in Golang to integrate with OpenAI's API",
		"max_tokens": 50,
		"model":      "gpt-3.5-turbo",
	}
	apiKey := "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"

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(respBody))
}
