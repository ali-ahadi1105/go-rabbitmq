package main

import "net/http"

func (app *Application) authentication(w http.ResponseWriter, _ *http.Request) {
	payload := JsonResponse{
		Error:   false,
		Message: "authentication service working correctly",
	}

	err := app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

}
