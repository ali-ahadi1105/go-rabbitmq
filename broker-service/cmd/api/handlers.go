package main

import "net/http"

func (app *Application) broker(w http.ResponseWriter, r *http.Request) {
	payload := JsonResponse{
		Error:   false,
		Message: "Broker service working correctly",
	}

	err := app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

}
