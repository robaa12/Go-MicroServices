package main

import (
	"log-service/data"
	"net/http"
)

type JSONPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (app *Config) WriteLog(w http.ResponseWriter, r *http.Request) {
	// read the JSON into var
	var requestPayLoad JSONPayload
	_ = app.readJSON(w, r, &requestPayLoad)

	// insert data
	event := data.LogEntry{
		Name: requestPayLoad.Name,
		Data: requestPayLoad.Data,
	}

	err := app.Models.LogEntry.Insert(&event)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	resp := jsonResponse{
		Error:   false,
		Message: "Message Logged",
	}

	app.writeJSON(w, http.StatusAccepted, resp)
}
