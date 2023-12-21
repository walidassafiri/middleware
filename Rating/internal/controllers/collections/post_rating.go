package collections

import (
	"encoding/json"
	"middleware/example/internal/models"
	"middleware/example/internal/services/collections"
	"github.com/sirupsen/logrus"
	"net/http"
	"fmt"
	"io/ioutil"
)

// PostRating
// @Tags         PostRating
// @Summary      Post a Rating.
// @Description  Post a Rating.
// @Body 		 json	UpdateRating	 
// @Success      200            {object}  models.Rating
// @Failure      500            "Something went wrong"
// @Router       /ratings/ [post]
func PostRating(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body)
	var post models.InsertRating

	err := json.Unmarshal(reqBody, &post)
	if err != nil {
		fmt.Println("Erreur lors du décodage JSON:", err)
		return
	}
	
	collection,err := collections.PostRating(post)
	
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

	w.WriteHeader(http.StatusCreated)
	body, _ := json.Marshal(collection)
	_, _ = w.Write(body)
	return
}
