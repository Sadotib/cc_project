package handlers

import (
	"net/http"

	"github.com/Sadotib/cc_project/handlers/helpers"
	pages "github.com/Sadotib/cc_project/views/pages"
)

func CloudHandler(w http.ResponseWriter, r *http.Request) {

	imgPath, err := helpers.RandomImage()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	// w.Header().Set("Content-Type", "image/png")

	// w.Header().Set("Content-Type", "text/html; charset=utf-8")

	helpers.Render(w, r, pages.Cloud(imgPath))
	// w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// w.Write([]byte("hello"))
	// w.Header().Set("Content-Type", "image/png")
	// w.Write(imgBytes)
	// fmt.Printf(string(imgBytes))

}
