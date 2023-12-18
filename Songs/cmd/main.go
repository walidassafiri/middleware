package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/controllers/collections"
	"middleware/example/internal/helpers"
	_ "middleware/example/internal/models"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	r.Route("/songs", func(r chi.Router) {
		r.Get("/", collections.GetAllSongs)
		//r.Route("/{id}", func(r chi.Router) {
			//r.Use(collections.Ctx)
			//r.Get("/", collections.GetSongsById)
			//r.Put("/", collections.UpdateSong)
			//r.Delete("/", collections.CreateSong)

		//})
		//r.Post("/", collections.CreateSong)
	})


	logrus.Info("[INFO] Web server started. Now listening on *:8080")
	logrus.Fatalln(http.ListenAndServe(":8082", r))
}

func init() {
	db, err := helpers.OpenDB()
	if err != nil {
		logrus.Fatalf("error while opening database : %s", err.Error())
	}
	schemes := []string{
		`CREATE TABLE IF NOT EXISTS songs (
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

