package collections

import (
	"encoding/json"
	"middleware/example/internal/models"
	"middleware/example/internal/repositories/collections"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// GetUser
// @Tags         users
// @Summary      Get a user.
// @Description  Get a user.
// @Param        id           	path      string  true  "User UUID formatted ID"
// @Success      200            {object}  models.User
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /user/{id} [get]
func GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	collectionId, _ := ctx.Value("userId").(uuid.UUID)

	collection, err := collections.GetUserById(collectionId)
	if err != nil {
		logrus.Errorf("error: %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(customError.Code)

			if encodeErr := json.NewEncoder(w).Encode(customError); encodeErr != nil {
				logrus.Errorf("error encoding custom error: %s", encodeErr.Error())
				// Gérez l'erreur d'encodage ici si nécessaire
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(collection)
	_, _ = w.Write(body)
	return
}
