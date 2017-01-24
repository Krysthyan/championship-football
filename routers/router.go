package routers

import (
	//"championship-football/controllers"
	"github.com/gorilla/mux"
)

func InitRouter(championship *mux.Router)  {

	RouterTeam(championship)
	RouterPerson(championship)
	RouterChampionship(championship)

}
