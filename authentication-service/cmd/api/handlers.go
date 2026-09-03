package main

import "net/http"

func (app *Application) authentication(w http.ResponseWriter, _ *http.Request) {
	payload := JsonResponse{
		Error:   false,
		Message: "authentication Service is running",
	}

	err := app.writeJSON(w, http.StatusOK, payload)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
