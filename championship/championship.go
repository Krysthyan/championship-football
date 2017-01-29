package championship

import (
	"championship-football/routers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func Run() {
	championship := mux.NewRouter().StrictSlash(false)
	routers.InitRouter(championship)
	championship.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))
	http.Handle("/", championship)

	server := &http.Server{
		Addr:           ":8080",
		Handler:        championship,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 50,
	}

	LogMain()
	log.Println(server.ListenAndServe())
}

func LogMain() {
	log.Println("\t    UNIVERSIDAD DE CUENCA")
	log.Println("\t    Facultad de ingenieria")
	log.Println("\t   Proyecto de base de datos 2")
	log.Println("\t       GO and Angular 2")
	log.Println("\t          Gorilla Mux")
	log.Println("   Listening http://localhost:8080 .....")
}
