package musicreviews

import uuid "github.com/satori/go.uuid"

// AlbumReview is an album review.
type AlbumReview struct {
	ID    uuid.UUID
	Album string
	Artist
	// Rating is out of 10
	Rating int
}

// AlbumStorage is an interface for
// storing/saving an album.
type AlbumReviewStorage interface {
	LoadAll() ([]AlbumReview, error)
	Load(string) (AlbumReview, error)
	Save(AlbumReview) (uuid.UUID, error)
}

// Artist is a music artist.
type Artist struct {
	ID   uuid.UUID
	Name string
}

// ArtistStorage is an interface for
// storing/saving an artist.
type ArtistStorage interface {
	LoadAll() ([]Artist, error)
	Load(string) (Artist, error)
	Save(Artist) (uuid.UUID, error)
}
