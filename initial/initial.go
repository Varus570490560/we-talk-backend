package initial

import (
	"github.com/lqs/sqlingo"
	"we-talk-backend/account"
)

type Service struct {
	Database       sqlingo.Database
	AccountService *account.AccountService
}

func Initial() (*Service, error) {
	service := new(Service)
	db, err := sqlingo.Open("mysql", "root:password@/we_talk")
	if err != nil {
		return nil, err
	}
	service.Database = db
	service.AccountService = account.NewAccountService(db)
	return service, nil
}
