package collections

import (
	"fmt"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"

	"github.com/gofrs/uuid"
)

func GetAllUsers() ([]models.UserPublic, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT id, name, username FROM users")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	collections := []models.UserPublic{}
	for rows.Next() {
		var data models.UserPublic
		err = rows.Scan(&data.Id, &data.Name, &data.Username)
		if err != nil {
			return nil, err
		}
		collections = append(collections, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return collections, err
}

func GetUserById(id uuid.UUID) (*models.UserPublic, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT id, name, username FROM users WHERE id=?", id.String())
	helpers.CloseDB(db)

	var collection models.UserPublic
	err = row.Scan(&collection.Id, &collection.Name, &collection.Username)
	if err != nil {
		return nil, err
	}
	return &collection, err
}
func SetUser(name string, username string) (string, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return "", err
	}
	uuid, err := uuid.NewV4()

	result, err := db.Exec("INSERT INTO users (id, name, username) VALUES (?, ?, ?)", uuid.String(), name, username)
	if err != nil {
		fmt.Println("Erreur lors de l'insertion dans la base de données:", err)
		return "", err
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Println("Nombre de lignes affectées:", rowsAffected)
	helpers.CloseDB(db)

	return uuid.String(), err
}
func DeleteUserById(userId uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM users WHERE id=?", userId)
	if err != nil {
		fmt.Println("Erreur lors de la suppression dans la base de données:", err)
	}

	helpers.CloseDB(db)

	return err
}
func UpdateUser(userId uuid.UUID, name string, username string) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE users SET name=? , username=? WHERE id=?", name, username, userId)
	if err != nil {
		fmt.Println("Erreur lors de la suppression dans la base de données:", err)
	}

	helpers.CloseDB(db)

	return err
}
