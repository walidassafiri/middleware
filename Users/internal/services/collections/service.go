package collections

import (
	"database/sql"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/collections"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

func GetAllUsers() ([]models.UserPublic, error) {
	var err error
	// calling repository
	collections, err := repository.GetAllUsers()
	// managing errors
	if err != nil {
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return collections, nil
}

func GetUserById(id uuid.UUID) (*models.UserPublic, error) {
	collection, err := repository.GetUserById(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return nil, &models.CustomError{
				Message: "user not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return nil, &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return collection, err
}
func SetUser(name string, username string) (string, error) {
	//var err error
	uuid, err := repository.SetUser(name, username)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return "", &models.CustomError{
				Message: "user not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error adding user : %s", err.Error())
		return "", &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}
	// managing errors

	return uuid, nil
}

func DeleteUserById(id uuid.UUID) error {
	err := repository.DeleteUserById(id)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return &models.CustomError{
				Message: "user not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error deleting user : %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}
	return err
}
func UpdateUser(id uuid.UUID, name string, username string) error {
	err := repository.UpdateUser(id, name, username)
	if err != nil {
		if err.Error() == sql.ErrNoRows.Error() {
			return &models.CustomError{
				Message: "user not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error updating user : %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}
	return err
}
