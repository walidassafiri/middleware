package collections

import (
	"fmt"
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"

	"github.com/gofrs/uuid"
)

func GetAllUsers() ([]models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT id, name, mail FROM users")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	collections := []models.User{}
	for rows.Next() {
		var data models.User
		err = rows.Scan(&data.Id, &data.Name, &data.Mail)
		if err != nil {
			return nil, err
		}
		collections = append(collections, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return collections, err
}

func GetUserById(id uuid.UUID) (*models.User, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT id, name, mail FROM users WHERE id=?", id.String())
	helpers.CloseDB(db)

	var collection models.User
	err = row.Scan(&collection.Id, &collection.Name, &collection.Mail)
	if err != nil {
		return nil, err
	}
	return &collection, err
}
func SetUser(name string, mail string) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}
	uuid, err := uuid.NewV4()

	result, err := db.Exec("INSERT INTO users (id, name, mail, password) VALUES (?, ?, ?, ?)", uuid.String(), name, mail, "password")
	if err != nil {
		fmt.Println("Erreur lors de l'insertion dans la base de données:", err)
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	fmt.Println("Nombre de lignes affectées:", rowsAffected)
	helpers.CloseDB(db)

	/*var collection models.User
	err = row.Scan(&collection.Id, &collection.Name, &collection.Mail)
	if err != nil {
		return err
	}*/
	return err
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
func UpdateUser(userId uuid.UUID, name string, mail string) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE users SET name=? , mail=? WHERE id=?", name, mail, userId)
	if err != nil {
		fmt.Println("Erreur lors de la suppression dans la base de données:", err)
	}

	helpers.CloseDB(db)

	return err
}
