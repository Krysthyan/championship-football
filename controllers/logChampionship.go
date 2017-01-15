package controllers

import (
	"github.com/fatih/color"
	"log"
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
		log.Println(blue("GET ") + "----> " + yellow(url) + " ===> " + cyan(status))
	} else {
		log.Println(green("POST ") + "----> " + yellow(url) + " ===> " + cyan(status))
	}
}
