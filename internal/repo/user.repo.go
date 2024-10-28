package repo

import (
	"github.com/dinos/go-ecommerce-be-api/global"
	"github.com/dinos/go-ecommerce-be-api/internal/database"
)

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

type userRepo struct {
	sqlc *database.Queries
}

func (up *userRepo) GetUserByEmail(email string) bool {
	//SELECT * FROM user WHERE email = '??' ORDER BY email
	//row := global.Mdb.Table(TableNameGoCrmUser).Where("usr_email = ?", email).First(&model.GoCrmUser{}).RowsAffected
	//return row != NumberNull
	user, err := up.sqlc.GetUserByEmailSQLC(ctx, email)
	if err != nil {
		return false
	}
	return user.UsrID != 0
}

func NewUserRepo() IUserRepo {
	return &userRepo{
		sqlc: database.New(global.Mdbc),
	}
}
