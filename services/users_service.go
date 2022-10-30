package services

import (
	"github.com/Bash-Clevin/bookstore_users_api/domain/users"
	"github.com/Bash-Clevin/bookstore_users_api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
