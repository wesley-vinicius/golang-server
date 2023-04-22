package service

import (
	"github.com/wesley-vinicius/golang-server/core/dto"
	"github.com/wesley-vinicius/golang-server/core/entity"
)

type UserService struct {
	repo entity.UserRepository
}

func NewSearchUser(repo entity.UserRepository) UserService {
	return UserService{repo: repo}
}

func (s *UserService) Search(query *string) ([]dto.User, error) {
	users, err := s.repo.SearchByNameOrUserName(query)
	if err != nil {
		return nil, err
	}

	var response []dto.User
	for _, u := range users {
		response = append(response, dto.User{
			Uuid:     u.Uuid,
			Name:     u.Name,
			Username: u.Username,
		})
	}

	return response, nil
}
