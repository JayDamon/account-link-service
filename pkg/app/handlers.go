package app

import (
	"net/http"

	httptoolbox "github.com/jaydamon/http-toolbox"
)

func (app *App) Broker(w http.ResponseWriter, r *http.Request) {
	payload := httptoolbox.JsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	var tools httptoolbox.Tools

	_ = tools.WriteJSON(w, http.StatusOK, payload)
}
