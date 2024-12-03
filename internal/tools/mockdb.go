package tools

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"sudipto": {
		AuthToken: "123sudo",
		Username:  "sudipto",
	},
	"harshada": {
		AuthToken: "456harsha",
		Username:  "harshada",
	},
}

var mockDoBDetails = map[string]DoBDetails{
	"sudipto": {
		Date:     "08-12-2000",
		Username: "sudipto",
	},
	"harshada": {
		Date:     "23-11-1997",
		Username: "harshada",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDB) GetUserDoB(username string) *DoBDetails {
	var coinData = DoBDetails{}
	coinData, ok := mockDoBDetails[username]
	if !ok {
		return nil
	}
	return &coinData
}

func (d *mockDB) SetupDB() error {
	return nil //mock call
}
