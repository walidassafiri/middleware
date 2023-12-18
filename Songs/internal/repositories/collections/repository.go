package collections

import (
	"middleware/example/internal/helpers"
	"middleware/example/internal/models"

	"github.com/gofrs/uuid"
)

func GetAllSongs() ([]models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM song")
	helpers.CloseDB(db)
	if err != nil {
		return nil, err
	}

	// parsing datas in object slice
	songs := []models.Song{}
	for rows.Next() {
		var data models.Song
		err = rows.Scan(&data.Id, &data.Artist, &data.Title, &data.Album, &data.Content)
		if err != nil {
			return nil, err
		}
		songs = append(songs, data)
	}
	// don't forget to close rows
	_ = rows.Close()

	return songs, err
}

func GetSongById(id uuid.UUID) (*models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return nil, err
	}
	row := db.QueryRow("SELECT * FROM song WHERE id=?", id.String())
	helpers.CloseDB(db)

	var song models.Song
	err = row.Scan(&song.Id, &song.Artist, &song.Title, &song.Album, &song.Content)
	if err != nil {
		return nil, err
	}
	return &song, err
}

func CreateSong(song models.Song) (error) {
    db, err := helpers.OpenDB()
    if err != nil {
        return err
    }

	newUUID, err := uuid.NewV4()
	if err != nil {
		return err
	}

    _, err = db.Exec("INSERT INTO song (id, artist, title, album, content) VALUES (?, ?, ?, ?, ?)",
	newUUID.String(), song.Artist, song.Title, song.Album, song.Content)
    helpers.CloseDB(db)
    if err != nil {
        return err
    }

    return nil
}

func UpdateSong(id uuid.UUID, song models.Song) (models.Song, error) {
	db, err := helpers.OpenDB()
	if err != nil {
		return models.Song{}, err
	}

	_, err = db.Exec("UPDATE song SET artist=?, title=?, album=?, content=? WHERE id=?",
		song.Artist, song.Title, song.Album, song.Content, id.String())
	helpers.CloseDB(db)
	if err != nil {
		return models.Song{}, err
	}

	return song, nil
}

func DeleteSong(id uuid.UUID) error {
	db, err := helpers.OpenDB()
	if err != nil {
		return err
	}

	_, err = db.Exec("DELETE FROM song WHERE id=?", id.String())
	helpers.CloseDB(db)
	if err != nil {
		return err
	}

	return nil
}
