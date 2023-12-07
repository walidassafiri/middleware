package collections

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"middleware/example/internal/models"
	"middleware/example/internal/services/collections"
	"net/http"

	"github.com/sirupsen/logrus"
)

// PostUser
// @Tags         postUser
// @Summary      Post user.
// @Description  Adding a user.
// @Success      200
// @Failure      500             "Something went wrong"
// @Router       /user [post]
func PostUsers(w http.ResponseWriter, r *http.Request) {
	// calling service
	reqBody, _ := ioutil.ReadAll(r.Body)
	var post models.InsertUser
	err := json.Unmarshal(reqBody, &post)
	if err != nil {
		fmt.Println("Erreur lors du décodage JSON:", err)
		return
	}

	erreur := collections.SetUser(post.Name, post.Mail)
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
	return
}
