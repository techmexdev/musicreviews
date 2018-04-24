package postgres

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/satori/go.uuid"
	"github.com/techmexdev/musicreviews"
)

// NewAlbumReviewStorage returns a AlbumReviewStorage postgres implementation.
func NewAlbumReviewStorage(dsn string) *AlbumReviewStorage {
	return &AlbumReviewStorage{sqlx.MustConnect("postgres", dsn)}
}

// AlbumReviewStorage implements musicreviews.AlbumReviewStorage.
type AlbumReviewStorage struct {
	*sqlx.DB
}

// LoadAll returns all stored album_reviews.
func (db *AlbumReviewStorage) LoadAll() ([]musicreviews.AlbumReview, error) {
	var ars []musicreviews.AlbumReview

	err := db.Select(&ars, `
		SELECT album_review.name, album_review.rating, artist.name
		FROM album_review, artist
		WHERE album_review.artist_id = artist.id`)
	if err != nil {
		return []musicreviews.AlbumReview{}, err
	}

	return ars, nil
}

// Load returns all stored festivals with that name.
func (db *AlbumReviewStorage) Load(name string) (musicreviews.AlbumReview, error) {
	var ar musicreviews.AlbumReview

	err := db.Select(&ar, "SELECT name FROM album_review WHERE name ="+name)
	if err != nil {
		return musicreviews.AlbumReview{}, err
	}

	return ar, nil
}

// Save inserts the festival in the databdbe.
func (db *AlbumReviewStorage) Save(ar musicreviews.AlbumReview) (uuid.UUID, error) {
	var artistID uuid.UUID
	var err error

	artStore := ArtistStorage{db.DB}

	artist, err := artStore.Load(ar.Artist.Name)
	if err != nil {
		log.Println(err)
		artistID, err = artStore.Save(ar.Artist)
		if err != nil {
			return uuid.UUID{}, err
		}
	} else {
		artistID = artist.ID
	}

	id, err := uuid.NewV4()
	if err != nil {
		log.Println(err)
	}

	_, err = db.Exec(
		`INSERT INTO album_review(id, artist_id, name, rating) VALUES($1, $2, $3, $4)
		ON CONFLICT(name) DO UPDATE
		SET rating = EXCLUDED.rating`,
		id, artistID, ar.Album, ar.Rating,
	)
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}

// FromArtist retrieves an artist's albums.
func (db *AlbumReviewStorage) FromArtist(artistName string) ([]musicreviews.AlbumReview, error) {
	var aa []musicreviews.AlbumReview

	q := `SELECT name FROM album_review WHERE id IN (
  	SELECT album_review_id FROM album_artist WHERE artist_id=(
  		SELECT id FROM artist WHERE name="$1"))`

	err := db.Select(&aa, q, artistName)
	if err != nil {
		return []musicreviews.AlbumReview{}, err
	}

	return aa, nil
}
