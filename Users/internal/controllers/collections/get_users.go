package collections

import (
	"encoding/json"
	"middleware/example/internal/models"
	"middleware/example/internal/services/collections"
	"net/http"

	"github.com/sirupsen/logrus"
)

// GetUsers
// @Tags         users
// @Summary      Get users.
// @Description  Get users.
// @Success      200            {array}  models.User
// @Failure      500             "Something went wrong"
// @Router       /user [get]
func GetUsers(w http.ResponseWriter, _ *http.Request) {
	// calling service
	collections, err := collections.GetAllUsers()
	if err != nil {
		// logging error
		logrus.Errorf("error : %s", err.Error())
		customError, isCustom := err.(*models.CustomError)
		if isCustom {
			// writing http code in header
			w.WriteHeader(customError.Code)
			// writing error message in body
			body, _ := json.Marshal(customError)
			_, _ = w.Write(body)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	//	w.Header().Set("Content-Type", "application/vnd.yourcustomtype+json")
	body, _ := json.Marshal(collections)
	_, _ = w.Write(body)
	return
}
