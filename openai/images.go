package openai

import (
	"net/http"

	"github.com/dghubble/sling"
)

type Image struct {
	URL string `json:"url"`
}

type ImageAPIResponse struct {
	Created int64   `json:"created"`
	Data    []Image `json:"data"`
}
type CreateImageAPIRequest struct {
	Prompt         string `json:"prompt,omitempty"` // maximum length is 1000
	N              int    `json:"n,omitempty"`
	ResponseFormat string `json:"response_format,omitempty"`
	User           string `json:"user,omitempty"`
}

type ImagesService struct {
	sling *sling.Sling
}

func newImagesService(sling *sling.Sling) *ImagesService {
	return &ImagesService{
		sling: sling.Path("images/"),
	}
}

func (i *ImagesService) Create(config *CreateImageAPIRequest) (*ImageAPIResponse, *http.Response, error) {
	images := new(ImageAPIResponse)

	resp, err := i.sling.New().Post("generations").BodyJSON(config).Receive(images, nil)
	return images, resp, err
}
