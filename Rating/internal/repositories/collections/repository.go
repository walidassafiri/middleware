package collections

import (
	"github.com/gofrs/uuid"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"
)

func GetAllRatings() ([]models.Rating, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM ratings")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	collections := []models.Rating{}
	for rows.Next() {
		var data models.Rating
		err = rows.Scan(&data.Id,&data.Score,&data.IdUser,&data.IdSong,&data.Content)
		if err != nil {
			return nil, err
		}
		collections = append(collections, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return collections, err
}

func GetRatingById(id uuid.UUID) (*models.Rating, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM ratings WHERE id=?", id.String())
	helpers.CloseDB(db)

	var collection models.Rating
	err = row.Scan(&collection.Id,&collection.Score,&collection.IdUser,&collection.IdSong,&collection.Content)
	if err != nil {
		return nil, err
	}
	return &collection, err
}
func PostRating(insert models.InsertRating) (error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	newUUID, err := uuid.NewV4()
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO ratings(id,score,idUser,idSong,content)VALUES(?,?,?,?,?);",newUUID.String(), &insert.Score,&insert.IdUser,&insert.IdSong,&insert.Content)
	helpers.CloseDB(db)
	if err != nil {
        return err
    }

	return err
}
