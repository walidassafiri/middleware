package collections

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"middleware/example/internal/services/collections"
	"middleware/example/internal/models"
	"github.com/sirupsen/logrus"
	"net/http"
)

// DeleteRating
// @Tags         DeleteRating
// @Summary      Delete raiting.
// @Description  Delete raiting by id.
// @Param        id           	path      string  true  "Rating UUID formatted ID"
// @Success      204            ""
// @Failure      500            "Something went wrong"
// @Router       /ratings/{id} [delete]
func DeleteRating(w http.ResponseWriter, r *http.Request) {
	
	ctx := r.Context()
	RatingId, _ := ctx.Value("RatingId").(uuid.UUID)

	err := collections.DeleteRating(RatingId)
	
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

	w.WriteHeader(http.StatusNoContent)
	return
}
