package postgres

import (
	"testing"

	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq" // init postgres driver
	"github.com/techmexdev/musicreviews"
	"github.com/techmexdev/musicreviews/pkg/postgres"
)

func TestStoreArtist(t *testing.T) {
	dsn := "postgres://musicreviews_test:themusic123@localhost/musicreviews_test?sslmode=disable"
	postgres.MigrateDB(dsn)
	artStore := postgres.NewArtistStorage(dsn)

	aa := []musicreviews.Artist{
		musicreviews.Artist{Name: "The Beatles"},
		musicreviews.Artist{Name: "Lana Del Rey"},
		musicreviews.Artist{Name: "Kanye West"},
	}

	for _, a := range aa {
		_, err := artStore.Save(a)
		if err != nil {
			t.Fatal(err)
		}
	}

	storedA, err := artStore.LoadAll()
	if err != nil {
		t.Fatal(err)
	}

	if len(storedA) != 3 {
		t.Fatal("have " + string(len(storedA)) + "stored artists, want 3")
	}

	for i := range storedA {
		if storedA[i].Name != aa[i].Name {
			if err != nil {
				t.Fatalf("error loading festivals. have %#v, want %#v", storedA[i], aa[i])
			}
		}
	}
}
