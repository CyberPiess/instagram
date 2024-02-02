package instagram

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	database "github.com/CyberPiess/instagram/internal/app/instagram/database"
	_ "github.com/lib/pq"
)

func TestCreate(t *testing.T) {
	db, err := database.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Create(w, r, db)
	}))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", resp.StatusCode)
	}

}
