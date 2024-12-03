package handlers

import (
	"encoding/json"
	"go-api/api"
	"go-api/internal/tools"
	"net/http"

	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetUserDoB(w http.ResponseWriter, r *http.Request) {
	var params = api.DoBParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var db *tools.DBInterface
	db, err = tools.NewDB()
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	tokenDetails := (*db).GetUserDoB(params.Username)
	if tokenDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var response = api.DoBResponse{
		Date: (*tokenDetails).Date,
		Code: http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}
