package collections

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"middleware/example/internal/services/collections"
	"middleware/example/internal/models"
	"github.com/sirupsen/logrus"
	"net/http"
	"io/ioutil"
	"fmt"
)

// GetCollection
// @Tags         collections
// @Summary      Get a collection.
// @Description  Get a collection.
// @Param        id           	path      string  true  "Collection UUID formatted ID"
// @Success      200            {object}  models.Collection
// @Failure      422            "Cannot parse id"
// @Failure      500            "Something went wrong"
// @Router       /collections/{id} [get]
func UpdateRating(w http.ResponseWriter, r *http.Request) {
	
	ctx := r.Context()
	RatingId, _ := ctx.Value("RatingId").(uuid.UUID)

	reqBody, _ := ioutil.ReadAll(r.Body)
	var upmodel models.UpdateRating

	err := json.Unmarshal(reqBody, &upmodel)
	if err != nil {
		fmt.Println("Erreur lors du décodage JSON:", err)
		return
	}
	
	err = collections.UpdateRating(RatingId,upmodel)
	
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
