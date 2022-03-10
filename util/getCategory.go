package util

import (
	"encoding/json"
	"io"
	"net/http"
	// "github.com/Elbi123/telegram-bot/util"
)

func GetQuizCategory(url string) []string {
	var array []map[string]interface{}
	category := make([]string, 3)
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	jsonError := json.Unmarshal([]byte(body), &array)
	if jsonError != nil {
		panic(err)
	}

	for _, arrayValue := range array {
		for key, value := range arrayValue {
			if key == "category" {
				category = append(category, value.(string))
			}
		}
	}

	return RemoveDuplicateStr(category)
}
