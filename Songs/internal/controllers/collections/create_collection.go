package collections

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/repositories/collections"
	"net/http"
)

// CreateSong
// @Tags         collections
// @Summary      Create a new song.
// @Description  Create a new song.
// @Param        song         	body      models.Song   true  "Song object to be created"
// @Success      201            {object}  models.Song
// @Failure      400            "Invalid request body"
// @Failure      500            "Something went wrong"
// @Router       /collections [post]
func CreateSong(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the song details
	var newSong models.Song
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newSong); err != nil {
		logrus.Errorf("Error decoding request body: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Validate the song details (you can implement your own validation logic)

	// Create the song in the database
	createdSong, err := collections.CreateSong(newSong)
	if err != nil {
		logrus.Errorf("Error creating song: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return the created song as JSON
	w.WriteHeader(http.StatusCreated)
	responseBody, _ := json.Marshal(createdSong)
	_, _ = w.Write(responseBody)
}
