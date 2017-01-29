package routers

import (
	"github.com/gorilla/mux"
	"championship-football/controllers"
)

func RouterStage(championship *mux.Router)  {
	championship.HandleFunc(
		"/stage/generateStage",
		controllers.GenerateStage,
	).Methods("POST")

	championship.HandleFunc(
		"/stage/playGroup",
		controllers.PlayGroupController,
	).Methods("GET")
}
