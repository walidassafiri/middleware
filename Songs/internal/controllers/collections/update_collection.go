package collections

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/repositories/collections"
	"net/http"
)

// UpdateSong
// @Tags         collections
// @Summary      Update a song in the collection.
// @Description  Update a song in the collection.
// @Param        id           	path      string  true  "Song UUID formatted ID"
// @Param        song         	body      object  true  "Song object to update"
// @Success      200            {object}  models.Song
// @Failure      400            "Invalid request payload"
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /collections/{id} [put]
func UpdateSong(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	collectionID, _ := ctx.Value("collectionId").(uuid.UUID)

	var updatedSong models.Song
	err := json.NewDecoder(r.Body).Decode(&updatedSong)
	if err != nil {
		logrus.Errorf("error decoding request body: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	song, err := collections.UpdateSong(collectionID, updatedSong)
	if err != nil {
		logrus.Errorf("error updating song: %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			w.WriteHeader(customError.Code)
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(song)
	_, _ = w.Write(body)
	return
}
