package middleware

import (
	"errors"
	"go-api/api"
	"go-api/internal/tools"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var ErrUnauthorised = errors.New("invalid user or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := r.URL.Query().Get("username")
		token := r.Header.Get("Authorization")
		if username == "" || token == "" {
			log.Error(ErrUnauthorised)
			api.RequestErrorHandler(w, ErrUnauthorised)
			return
		}
		var db *tools.DBInterface
		var err error
		db, err = tools.NewDB()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}
		loginDetails := (*db).GetUserLoginDetails(username)

		if loginDetails == nil || token != (*loginDetails).AuthToken {
			log.Error(ErrUnauthorised)
			api.RequestErrorHandler(w, ErrUnauthorised)
			return
		}
		next.ServeHTTP(w, r)
	})
}
