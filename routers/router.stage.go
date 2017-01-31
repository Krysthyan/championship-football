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

	championship.HandleFunc(
		"/stage/getListChampionship",
		controllers.GetListStageFromChampionship,
	).Methods("GET")

	championship.HandleFunc(
		"/stage/getTablaPosicion",
		controllers.GetTablePosicion,
	).Methods("GET")

	championship.HandleFunc(
		"/stage/getPlayOffs",
		controllers.GetPlayOff,
	).Methods("GET")

	championship.HandleFunc(
		"/stage/getTeamPlayOffs",
		controllers.GetTeamPlayOffs,
	).Methods("GET")
}
