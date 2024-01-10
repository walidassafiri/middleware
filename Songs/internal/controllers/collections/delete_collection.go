package collections

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/repositories/collections"
	"net/http"
)

// DeleteSong
// @Tags         collections
// @Summary      Delete a song from a collection.
// @Description  Delete a song from a collection.
// @Param        id           	path      string  true  "Song UUID formatted ID"
// @Success      204            "No Content"
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /songs/{id} [delete]
func DeleteSong(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	collectionID, _ := ctx.Value("collectionId").(uuid.UUID)

	err := collections.DeleteSong(collectionID)
	if err != nil {
		logrus.Errorf("error: %s", err.Error())
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

	w.WriteHeader(http.StatusNoContent)
	return
}
