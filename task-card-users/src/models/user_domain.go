package models

import (
	"crypto/md5"
	"encoding/hex"
)

type UserDomainInterface interface {
	GetUsername() string
	GetUserpass() string
	GetEmail() string
	GetRoleId() int
	EncryptPassword()
}

func NewUserDomain(username, userpass, email string, roleId int) UserDomainInterface {
	return &userDomain{
		username: username,
		userpass: userpass,
		email:    email,
		roleId:   roleId,
	}
}

type userDomain struct {
	username string
	userpass string
	email    string
	roleId   int
}

func (u *userDomain) GetUsername() string {
	return u.username
}
func (u *userDomain) GetUserpass() string {
	return u.userpass
}
func (u *userDomain) GetEmail() string {
	return u.email
}
func (u *userDomain) GetRoleId() int {
	return u.roleId
}

func (u *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()

	hash.Write([]byte(u.userpass))
	u.userpass = hex.EncodeToString(hash.Sum(nil))
}
