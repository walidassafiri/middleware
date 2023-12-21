package collections

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"middleware/example/internal/models"
	"middleware/example/internal/services/collections"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

// UpdateUser
// @Tags         updateUsers
// @Summary      update a user.
// @Description  update a user.
// @Success      200
// @Failure      500             "Something went wrong"
// @Router       /user/{id}  [put]
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// calling service
	ctx := r.Context()
	collectionId, _ := ctx.Value("userId").(uuid.UUID)
	reqBody, _ := ioutil.ReadAll(r.Body)
	var post models.InsertUser
	err := json.Unmarshal(reqBody, &post)
	if err != nil {
		fmt.Println("Erreur lors du décodage JSON:", err)
		return
	}

	erreur := collections.UpdateUser(collectionId, post.Name, post.Username)
	if erreur != nil {
		logrus.Errorf("error : %s", erreur.Error())
		customError, isCustom := erreur.(*models.CustomError)
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
	id := models.UserName{Username: post.Username}
	body, _ := json.Marshal(id)
	w.Write(body)
	return
}
