package account

import "github.com/lqs/sqlingo"

type RegisterReturnStatus int32
type LoginReturnStatus int32

const (
	RegisterReturnStatusInvalid RegisterReturnStatus = 0
	RegisterReturnStatusSuccess RegisterReturnStatus = 1
	RegisterReturnStatusRepeat  RegisterReturnStatus = 2
	RegisterReturnStatusFailed  RegisterReturnStatus = 3
)
const (
	LoginReturnStatusInvalid       LoginReturnStatus = 0
	LoginReturnStatusSuccess       LoginReturnStatus = 1
	LoginReturnStatusWrongPassword LoginReturnStatus = 2
	LoginReturnStatusFailed        LoginReturnStatus = 3
)

type Service interface {
	Register(nickname string, password string) (userId int32, status RegisterReturnStatus)
	Login(nickname string, password string) (status LoginReturnStatus)
	Logout(nickname string)
}

type AccountService struct {
	db sqlingo.Database
}

func (a AccountService) Register(nickname string, password string) (userId int32, status RegisterReturnStatus) {

	return 0, 0
}

func (a AccountService) Login(nickname string, password string) (status LoginReturnStatus) {

	return 0
}

func (a AccountService) Logout(nickname string) {

	return
}

func NewAccountService(db sqlingo.Database) *AccountService {
	var service *AccountService
	service = new(AccountService)
	return service
}
