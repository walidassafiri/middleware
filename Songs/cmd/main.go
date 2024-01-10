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

	r.Route("/songs", func(r chi.Router) {
		r.Get("/", collections.GetSongs)
		r.Route("/{id}", func(r chi.Router) {
			r.Use(collections.Ctx)
			r.Get("/", collections.GetSong)
			r.Put("/", collections.UpdateSong)
			r.Delete("/", collections.DeleteSong)
		})
		r.Post("/", collections.CreateSong)
	})

	logrus.Info("[INFO] Web server started. Now listening on *:8092")
	logrus.Fatalln(http.ListenAndServe(":8092", r))
}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	schemes := []string{
		`CREATE TABLE IF NOT EXISTS song (
			id VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
			artist VARCHAR(255) NOT NULL,
			title VARCHAR(255) NOT NULL,
			album VARCHAR(255) NOT NULL,
			content VARCHAR(255) NOT NULL
		);`,
	}
	for _, scheme := range schemes {
		if _, err := db.Exec(scheme); err != nil {
			logrus.Fatalln("Could not generate table ! Error was : " + err.Error())
		}
	}
	helpers.CloseDB(db)
}
