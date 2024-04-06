package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"template-go-vercel/utils"

	openai "github.com/sashabaranov/go-openai"
)

type Preparation struct {
	Total           string
	PreparationTime string `json:"preparation_time"`
	Cooking         string
}

type PromptData struct {
	Prompt       string      `json:"prompt"`
	Description  string      `json:"description"`
	Preparation  Preparation `json:"preparation"`
	Ingredients  []string    `json:"ingredients"`
	Instructions []string    `json:"instructions"`
	Nutrition    []string    `json:"nutrition"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	utils.LoadEnvVariable()

	SystemPrompt := `sen yemek tarifleri veren bir aracsın kullanıcı sana sormak istediği yemek ile bilgileri JSON şemasında gelecek ve buna göre sende böyle bir JSON şeması gelecek:

gelen lang  değişkenine göre o dilde belirt!

{
  "prompt": "oluşturmak istediği tarif ile ilgili prompt",
"lang":  "string "
}

sende ise şu formatta JSON döndüreceksin:


{
  "results": {
  "prompt": "food name",
  "description": "food description",
  "preparation": {
    "total": "total preparation time",
    "preparation_time": "preparation time",
    "cooking": "cooking time"
  },
    "ingredients": [ingredients...],
    "instructions": [instructions...],
    "nutrition": [nutrition...]              
}
}}

prompt ile bir değer üretemiyorsan results dizisini boş döndürebilirsin.

Sadece JSON data döndür!`

	UserPropmtS := `{
  "prompt": "%s",
"lang":  "tr"
}`

	UserPropmt := fmt.Sprintf(UserPropmtS, r.FormValue("food"))

	client := openai.NewClient(os.Getenv("TOKEN"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: SystemPrompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: UserPropmt,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	w.Write([]byte(resp.Choices[0].Message.Content))
}
