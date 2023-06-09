package clients

import (
	"github.com/juanfer2/whorship_band/src/shared/utilities"
)

func CreateFactoryClientGpt3() Gpt3Client {
	baseUrl := utilities.GetEnv("GPT3_URL")
	token := utilities.GetEnv("GPT3_API_KEY")

	return NewGpt3Client(baseUrl, token)
}
