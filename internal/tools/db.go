package tools

import (
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct {
	AuthToken string
	Username  string
}
type DoBDetails struct {
	Date     string
	Username string
}

type DBInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserDoB(username string) *DoBDetails
	SetupDB() error
}

func NewDB() (*DBInterface, error) {
	var db DBInterface = &mockDB{}
	var err error = db.SetupDB()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &db, nil
}
