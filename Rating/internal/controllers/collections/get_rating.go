package collections

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/repositories/collections"
	"net/http"
)

// GetRating
// @Tags         GetRating
// @Summary      Get a Rating.
// @Description  Get a Rating.
// @Param        id           	path      string  true  "Rating UUID formatted ID"
// @Success      200            {object}  models.Rating
// @Failure      500            "Something went wrong"
// @Router       /ratings/{id} [get]
func GetRating(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	collectionId, _ := ctx.Value("RatingId").(uuid.UUID)

	collection, err := collections.GetRatingById(collectionId)
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
