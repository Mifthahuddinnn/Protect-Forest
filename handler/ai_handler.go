package handler

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AIAPI struct {
	APIURL string
}

type ChatCompletionResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func NewAIAPI() *AIAPI {
	return &AIAPI{
		APIURL: "https://wgpt-production.up.railway.app/v1/chat/completions",
	}
}

func (a *AIAPI) GetChatCompletion(messages []map[string]string) (string, error) {
	payload, err := json.Marshal(map[string]interface{}{
		"model":    "gpt-3.5-turbo",
		"messages": messages,
	})
	if err != nil {
		return "", err
	}

	resp, err := http.Post(a.APIURL, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response ChatCompletionResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	if len(response.Choices) == 0 {
		return "", nil
	}

	content := response.Choices[0].Message.Content

	return content, nil
}

func HandleChatCompletion(c echo.Context) error {
	api := NewAIAPI()

	var messages []map[string]string
	if err := c.Bind(&messages); err != nil {
		return err
	}
	content, err := api.GetChatCompletion(messages)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"completion": content,
	})
}
