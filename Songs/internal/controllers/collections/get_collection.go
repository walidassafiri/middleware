package collections

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/repositories/collections"
	"net/http"
)

// GetSongs
// @Tags         collections
// @Summary      Get songs from collections.
// @Description  Get a list of songs from collections.
// @Success      200            {array}  models.Song
// @Failure      500            "Something went wrong"
// @Router       /songs [get] 
func GetSong(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	collectionId, _ := ctx.Value("collectionId").(uuid.UUID)

	collection, err := collections.GetSongById(collectionId)
	if err != nil {
		logrus.Errorf("error : %s", err.Error())
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
	body, _ := json.Marshal(collection)
	_, _ = w.Write(body)
	return
}
