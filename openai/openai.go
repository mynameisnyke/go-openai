package openai

import (
	"net/http"

	"github.com/dghubble/sling"
)

const openaiBase = "https://api.openai.com/v1/"

// Client is a OpenAI client that allows you to make OpenAI API requests.
type Client struct {
	sling *sling.Sling
	//Openai API services
	Images *ImagesService
}

func NewClient(httpClient *http.Client, token string) *Client {
	base := sling.New().Client(httpClient).Base(openaiBase).Set("Authorization", "Bearer "+token)
	return &Client{
		sling:  base,
		Images: newImagesService(base.New()),
	}
}
