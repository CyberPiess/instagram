package application

import (
	"database/sql"
	"net/http"
)

type UserHandler interface {
	Create(w http.ResponseWriter, r *http.Request, db *sql.DB)
}
