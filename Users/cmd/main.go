package main

import (
	"middleware/example/internal/controllers/collections"
	"middleware/example/internal/helpers"
	_ "middleware/example/internal/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

func main() {
	r := chi.NewRouter()

	r.Route("/user", func(r chi.Router) {
		r.Get("/", collections.GetUsers)
		r.Post("/", collections.PostUsers)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(collections.Ctx)
			r.Get("/", collections.GetUser)
			r.Delete("/", collections.DeleteUser)
			r.Put("/", collections.UpdateUser)
		})
	})

	logrus.Info("[INFO] Web server started. Now listening on *:8080")
	logrus.Fatalln(http.ListenAndServe(":8088", r))
}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	schemes := []string{
		`CREATE TABLE IF NOT EXISTS collections (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			content VARCHAR(255) NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			name VARCHAR(255) NOT NULL ,
			mail VARCHAR(255) NOT NULL 
		
			
		);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)
}
