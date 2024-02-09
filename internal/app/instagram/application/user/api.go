package application

import (
	"net/http"
)

type UserHandler interface {
	Create(w http.ResponseWriter, r *http.Request)
}
