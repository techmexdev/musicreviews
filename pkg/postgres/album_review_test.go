package postgres

import (
	"testing"

	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"github.com/techmexdev/musicreviews"
	"github.com/techmexdev/musicreviews/pkg/postgres"
)

func TestStoreAlbumReview(t *testing.T) {
	dsn := "postgres://musicreviews_test:themusic123@localhost/musicreviews_test?sslmode=disable"
	postgres.MigrateDB(dsn)
	arStore := postgres.NewAlbumReviewStorage(dsn)

	ars := []musicreviews.AlbumReview{
		{Album: "Revolver", Artist: musicreviews.Artist{Name: "The Beatles"}, Rating: 10},
		{Album: "Born To Die", Artist: musicreviews.Artist{Name: "Lana Del Rey"}, Rating: 1},
		{Album: "Yeezus", Artist: musicreviews.Artist{Name: "Kanye West"}, Rating: 5},
	}
	for _, ar := range ars {
		_, err := arStore.Save(ar)
		if err != nil {
			t.Fatal("failed saving:", err)
		}
	}

	storedARs, err := arStore.LoadAll()
	if err != nil {
		t.Fatal("failes loading all", err)
	}

	if len(storedARs) != 3 {
		t.Fatal("have " + string(len(storedARs)) + "stored album reviews, want 3")
	}

	for i := range storedARs {
		if storedARs[i].Album != ars[i].Album || storedARs[i].Artist.Name != ars[i].Artist.Name ||
			storedARs[i].Rating != ars[i].Rating {
			if err != nil {
				t.Fatalf("error loading festivals. have %#v, want %#v", storedARs[i], ars[i])
			}
		}

	}
}
