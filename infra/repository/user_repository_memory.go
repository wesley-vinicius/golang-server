package repository

import (
	"github.com/software-engineer-challenge/core/entity"
)

type UserRepositoryMemory struct {
	userList []entity.User
}

func (r *UserRepositoryMemory) SearchByNameOrUserName(query *string) ([]entity.User, error) {
	println(*query)
	return r.userList, nil
}

func (r *UserRepositoryMemory) AddUser(user *entity.User) {
	r.userList = append(r.userList, *user)
}
