package handlers

import (
	"net/http"

	"github.com/Sadotib/cc_project/handlers/helpers"
	pages "github.com/Sadotib/cc_project/views/pages"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	helpers.Render(w, r, pages.Home())
}
