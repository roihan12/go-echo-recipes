package users

import (
	"echo-recipe/businesses/users"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) users.Repository {
	return &userRepository{
		conn: conn,
	}
}

func (ur *userRepository) Register(userDomain *users.Domain) users.Domain {
	password, _ := bcrypt.GenerateFromPassword([]byte(userDomain.Password), bcrypt.DefaultCost)

	rec := FromDomain(userDomain)

	rec.Password = string(password)

	result := ur.conn.Create(&rec)

	result.Last(&rec)

	return rec.ToDomain()
}

func (ur *userRepository) GetByEmail(email string, password string) users.Domain {
	var user User

	ur.conn.First(&user, "email = ?", email)

	if user.ID == 0 {
		fmt.Println("user not found")
		return users.Domain{}
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		fmt.Println("password failed!")
		return users.Domain{}
	}

	return user.ToDomain()
}

func (ur *userRepository) Update(userDomain *users.Domain) users.Domain {
	password, _ := bcrypt.GenerateFromPassword([]byte(userDomain.Password), bcrypt.DefaultCost)

	rec := FromDomain(userDomain)

	rec.Password = string(password)

	result := ur.conn.Updates(rec).Where(&rec, "ID = ?", userDomain.ID)

	result.Last(&rec)

	return rec.ToDomain()
}
