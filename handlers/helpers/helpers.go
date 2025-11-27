package helpers

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/a-h/templ"
)

func Render(w http.ResponseWriter, r *http.Request, c templ.Component) error {
	return c.Render(r.Context(), w)
} //function to render the templ component

func RandomImage() (string, error) {
	var x int = rand.Intn(5) + 1 // generates a random number between 1 and 5
	fmt.Printf("%d\n", x)
	var imgPath string
	switch x {
	case 1:
		// imgBytes, err = os.ReadFile("assets/1.webp") // returns []byte
		imgPath = "public/assets/1.png"
	case 2:
		// imgBytes, err = os.ReadFile("assets/2.jpg")
		imgPath = "public/assets/2.png"
	case 3:
		// imgBytes, err = os.ReadFile("assets/3.jpg")
		imgPath = "public/assets/3.png"
	case 4:
		// imgBytes, err = os.ReadFile("assets/4.png")
		imgPath = "public/assets/4.png"
	case 5:
		// imgBytes, err = os.ReadFile("assets/5.png")
		imgPath = "public/assets/5.png"
	}

	return imgPath, nil
} //function to get the paths of random images
