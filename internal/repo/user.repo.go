package repo

//type UserRepo struct {
//}
//
//func NewUserRepo() *UserRepo {
//	return &UserRepo{}
//}
//
//func (ul *UserRepo) GetInfoUser() string {
//	return "Tip js"
//}

type IUserRepo interface {
	GetUserByEmail(email string) bool
}

type userRepo struct{}

func (ur *userRepo) GetUserByEmail(email string) bool {
	//TODO implement me
	return true
}

func NewUserRepo() IUserRepo {
	return &userRepo{}
}
