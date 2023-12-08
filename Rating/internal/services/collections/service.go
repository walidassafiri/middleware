package collections

import (
	"database/sql"
	"errors"
	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
	"middleware/example/internal/models"
	repository "middleware/example/internal/repositories/collections"
	"net/http"
)

func GetAllRatings() ([]models.Rating, error) {
	var err error
	// calling repository
	collections, err := repository.GetAllRatings()
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

func GetRatingById(id uuid.UUID) (*models.Rating, error) {
	collection, err := repository.GetRatingById(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return nil, &models.CustomError{
				Message: "collection not found",
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
func PostRating(insert models.InsertRating) (error) {
	err := repository.PostRating(insert)
	if err != nil {
		logrus.Errorf("error while executing query : %s", err.Error())
		return  &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return err
}
func DeleteRating(id uuid.UUID) (error) {
	err := repository.DeleteRating(id)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return &models.CustomError{
				Message: "collection not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return err
}
func UpdateRating(id uuid.UUID,upmodel models.UpdateRating) (error) {
	err := repository.UpdateRating(id, upmodel)
	if err != nil {
		if errors.As(err, &sql.ErrNoRows) {
			return &models.CustomError{
				Message: "collection not found",
				Code:    http.StatusNotFound,
			}
		}
		logrus.Errorf("error retrieving collections : %s", err.Error())
		return &models.CustomError{
			Message: "Something went wrong",
			Code:    500,
		}
	}

	return err
}
