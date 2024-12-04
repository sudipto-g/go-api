package tools

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
