package routers

import (
	"github.com/gorilla/mux"
)

func InitRouter(championship *mux.Router)  {

	RouterTeam(championship)
	RouterPerson(championship)
	RouterChampionship(championship)

}
