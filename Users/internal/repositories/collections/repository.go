package collections

import (
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"

	"github.com/gofrs/uuid"
)

func GetAllUsers() ([]models.Collection, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM collection")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	collections := []models.Collection{}
	for rows.Next() {
		var data models.Collection
		err = rows.Scan(&data.Id, &data.Content)
		if err != nil {
			return nil, err
		}
		collections = append(collections, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return collections, err
}

func GetUserById(id uuid.UUID) (*models.Collection, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM collections WHERE id=?", id.String())
	helpers.CloseDB(db)

	var collection models.Collection
	err = row.Scan(&collection.Id, &collection.Content)
	if err != nil {
		return nil, err
	}
	return &collection, err
}
