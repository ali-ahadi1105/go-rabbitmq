package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type JsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    any `json:"data,omitempty"`
}

func (app *Application) readJSON(w http.ResponseWriter, r *http.Request, payload any) error {
	maxBytes := 1048576 // one megabyte

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(payload)

	if err != nil {
		return err
	}

	err = decoder.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("Body must have only a single json value")
	}

	return nil

}

func (app *Application) writeJSON(w http.ResponseWriter, status int, payload any, headers ...http.Header) error {
	out, err := json.Marshal(payload)

	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header()[k] = v
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil

}

func (app *Application) errorJSON(w http.ResponseWriter, err error, status ...int) error {

	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload JsonResponse
	payload.Error = true
	payload.Message = err.Error()

	return app.writeJSON(w, statusCode, payload)

}
