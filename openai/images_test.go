package openai

import (
	"fmt"
	"net/http"
	"testing"
)

func TestImages(t *testing.T) {

	openai := NewClient(http.DefaultClient, "token")
	query := "Illustrate a multi-panel comic with a jewish guy who is a chinese spy tricking his friends to sign up for chinese spyware in under the guise of a free egg foamer"

	t.Run("Should return N number of iamges for given query", func(t *testing.T) {

		const n = 4
		config := CreateImageAPIRequest{
			Prompt:         query,
			N:              n,
			ResponseFormat: "url",
		}
		images, resp, err := openai.Images.Create(&config)

		if err != nil || resp.StatusCode != 200 {
			t.Errorf("FAILED STATUS CODE: %d", resp.StatusCode)
		}

		got := len(images.Data)
		want := n

		if got != want {
			t.Errorf("Got %d, but wanted %d", got, want)
		}

		for i, img := range images.Data {
			fmt.Printf("IMG #%d: %s\n", i, img.URL)
		}
	})

}
