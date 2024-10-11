package repo

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (ul *UserRepo) GetInfoUser() string {
	return "Tip js"
}
