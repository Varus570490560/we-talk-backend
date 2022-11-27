package account

import (
	"fmt"
	"github.com/lqs/sqlingo"
	we_talk_dsl "we-talk-backend/generated/sqlingo"
)

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
	Register(nickname string, password string) (userId int64, status RegisterReturnStatus)
	Login(nickname string, password string) (status LoginReturnStatus)
	Logout(nickname string)
}

type AccountService struct {
	db sqlingo.Database
}

func (a AccountService) Register(nickname string, password string) (userId int64, status RegisterReturnStatus) {
	_, err := a.db.
		InsertInto(we_talk_dsl.User).
		Fields(we_talk_dsl.User.NickName, we_talk_dsl.User.Password, we_talk_dsl.User.Avatar).
		Values(nickname, password, 1).
		Execute()
	if err != nil {
		fmt.Println(err)
		return 0, RegisterReturnStatusRepeat
	}
	var userModel we_talk_dsl.UserModel
	_, err = a.db.
		SelectFrom(we_talk_dsl.User).
		Where(we_talk_dsl.User.NickName.
			Equals(nickname)).
		FetchFirst(&userModel)
	if err != nil {
		fmt.Println(err)
		return 0, RegisterReturnStatusFailed
	}
	return userModel.Id, RegisterReturnStatusSuccess
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
	service.db = db
	return service
}
