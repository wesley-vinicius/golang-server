package repository

import (
	"github.com/wesley-vinicius/golang-server/core/entity"
)

type UserRepositoryMemory struct {
	userList []entity.User
}

func (r *UserRepositoryMemory) SearchByNameOrUserName(query *string) ([]entity.User, error) {
	return r.userList, nil
}

func (r *UserRepositoryMemory) Create(user *entity.User) (entity.User, error) {
	r.userList = append(r.userList, *user)

	return *user, nil
}
