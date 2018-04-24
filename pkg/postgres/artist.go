package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
	"github.com/techmexdev/musicreviews"
)

// NewArtistStorage returns a ArtistStorage postgres implementation.
func NewArtistStorage(dsn string) *ArtistStorage {
	return &ArtistStorage{sqlx.MustConnect("postgres", dsn)}
}

// ArtistStorage implements musicreviews.ArtistStorage.
type ArtistStorage struct {
	*sqlx.DB
}

// LoadAll returns all stored artists.
func (db *ArtistStorage) LoadAll() ([]musicreviews.Artist, error) {
	var aa []musicreviews.Artist

	err := db.Select(&aa, "SELECT name FROM artist")
	if err != nil {
		return []musicreviews.Artist{}, err
	}

	return aa, nil
}

// Load returns all stored artists
// with that name.
func (db *ArtistStorage) Load(name string) (musicreviews.Artist, error) {
	var a musicreviews.Artist

	err := db.Get(&a, "SELECT name, id FROM artist WHERE name=$1", name)
	if err != nil {
		return musicreviews.Artist{}, err
	}

	return a, nil
}

// Save inserts the festival in the database.
func (db *ArtistStorage) Save(a musicreviews.Artist) (uuid.UUID, error) {
	id, err := uuid.NewV4()
	if err != nil {
		log.Println(err)
	}

	_, err = db.Exec(`
		INSERT INTO artist(id, name) VALUES($1, $2)
		ON CONFLICT DO NOTHING`, id, a.Name)
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}
