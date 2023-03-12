package service

import (
	"github.com/software-engineer-challenge/core/dto"
	"github.com/software-engineer-challenge/core/entity"
)

type SearchUser struct {
	repo entity.UserRepository
}

func NewSearchUser(repo entity.UserRepository) SearchUser {
	return SearchUser{repo: repo}
}

func (s *SearchUser) Search(query *string) ([]dto.User, error) {
	users, err := s.repo.SearchByNameOrUserName(query)
	if err != nil {
		return nil, err
	}

	var response []dto.User
	for _, u := range users {
		response = append(response, dto.User{
			Uuid:     u.Uuid(),
			Name:     u.Name(),
			Username: u.Username(),
		})
	}

	return response, nil
}
