package collections

import (
	"encoding/json"
	"middleware/example/internal/models"
	"middleware/example/internal/repositories/collections"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// DeleteUser
// @Tags         deleteUser
// @Summary      Delete user.
// @Description  Delete user.
// @Success      200
// @Failure      500             "Something went wrong"
// @Router       /user/{id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	collectionId, _ := ctx.Value("userId").(uuid.UUID)

	err := collections.DeleteUserById(collectionId)
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

	return
}
