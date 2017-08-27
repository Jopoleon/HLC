package main

import (
	"log"
	"net/http"

	"fmt"

	"github.com/Jopoleon/HighLoadCup/app"
	"github.com/bmizerany/pat"
	"github.com/sirupsen/logrus"
)

//GET /<entity>/<id>
//GET /users/<id>/visits для получения списка посещений пользователем
//GET /locations/<id>/avg для получения средней оценки достопримечательности
//POST /<entity>/<id> на обновление
//POST /<entity>/new на создание

func main() {
	err := app.Run()
	if err != nil {
		log.Fatal("App fatal error: ", err)
	}
	m := pat.New()
	//для получения данных о сущности
	m.Get("/users/:id", http.HandlerFunc(app.GetUser))
	m.Get("/locations/:id", http.HandlerFunc(app.GetLocation))
	m.Get("/visits/:id", http.HandlerFunc(app.GetVisit))

	m.Get("/users/:id/visits", http.HandlerFunc(app.GetUsersVisits))
	m.Get("/locations/:id/avg", http.HandlerFunc(app.GetLocationAvg))
	m.Get("/usersall", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		logrus.Info("All endpoint used")
		fmt.Fprintf(w, "%v", app.UserID_User)
	}))

	m.Post("/:entity/:id", http.HandlerFunc(app.PostUpdateEntity))
	m.Post("/:entity/new", http.HandlerFunc(app.PostNewEntity))

	// Register this pat with the default serve mux so that other packages
	// may also be exported. (i.e. /debug/pprof/*)
	http.Handle("/", m)
	logrus.Println("Server started on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
