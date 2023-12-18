package collections

import (
	"database/sql"
	"errors"
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
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: " not found",
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
func SetUser(name string, mail string, password string) (string, error) {
	//var err error
	uuid, err := repository.SetUser(name, mail, password)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return "", &models.CustomError{
				Message: " not found",
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
		if errors.As(err, &sql.ErrNoRows) {
			return &models.CustomError{
				Message: " not found",
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
func UpdateUser(id uuid.UUID, name string, mail string, password string) error {
	err := repository.UpdateUser(id, name, mail, password)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return &models.CustomError{
				Message: " not found",
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
