package repository

import (
	"database/sql"
	"majoo-test-case/entity/user"
)

type UserRepository interface {
	UserLogin(u *user.UserLoginModel) (*user.UserLoginDataModel, error)
	GetUserIDByMerchant(merchantId int) (int, error)
	GetUserIDByOutlet(outletId int) (int, error)
}

type mysqlUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) UserRepository {
	return &mysqlUserRepository{
		db: db,
	}
}

func (m *mysqlUserRepository) UserLogin(u *user.UserLoginModel) (*user.UserLoginDataModel, error) {
	userObj := new(user.UserLoginDataModel)
	row := m.db.QueryRow("SELECT id, password FROM Users WHERE user_name=?", u.Username)
	err := row.Scan(&userObj.ID, &userObj.Password)
	if err != nil {
		return userObj, err
	}
	return userObj, nil

}

func (m *mysqlUserRepository) GetUserIDByMerchant(merchantId int) (int, error) {
	row := m.db.QueryRow("SELECT user_id FROM Merchants WHERE id=?", merchantId)
	var userId int
	err := row.Scan(&userId)
	if err != nil {
		return 0, err
	}

	return userId, nil

}

func (m *mysqlUserRepository) GetUserIDByOutlet(outletId int) (int, error) {
	row := m.db.QueryRow("SELECT Merchants.user_id FROM Outlets JOIN Merchants ON Merchants.id=Outlets.merchant_id WHERE Outlets.id=?", outletId)
	var userId int
	err := row.Scan(&userId)
	if err != nil {
		return 0, err
	}

	return userId, nil

}
