package controllers

import (
	"github.com/fatih/color"
	"log"
	"net/http"
)

func LogChampionship(
	type_request string,
	url string,
	status string,
) {

	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()

	if type_request == "GET" {
		log.Println(blue("GET ") + "----> " + yellow("localhost:8080" + url) + " ===> " + cyan(status))
	} else {
		log.Println(green("POST ") + "----> " + yellow("localhost:8080" + url) + " ===> " + cyan(status))
	}
}

func Set_ResponseWrite(w http.ResponseWriter) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return w
}
