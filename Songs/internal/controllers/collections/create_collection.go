package collections

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	"middleware/example/internal/repositories/collections"
	"net/http"
)

// CreateSong
// @Tags         collections
// @Summary      Create a new song.
// @Description  Create a new song.
// @Param        song         	body      models.Song   true  "Song object to be created"
// @Success      201            {object}  models.Song
// @Failure      400            "Invalid request body"
// @Failure      500            "Something went wrong"
// @Router       /collections [post]
func CreateSong(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the song details
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newSong models.Song

	err := json.Unmarshal(reqBody, &newSong)
	if err != nil {
		fmt.Println("Erreur lors du décodage JSON:", err)
		return
	}

	err = collections.CreateSong(newSong)

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