package handle

import (
	"encoding/json"
	"github.com/software-engineer-challenge/core/entity"
	"github.com/software-engineer-challenge/core/service"
	"github.com/software-engineer-challenge/infra/repository"
	"net/http"
)

func SearchUserHandle(w http.ResponseWriter, e *http.Request) {
	w.Header().Set("Content-type", "json")

	query := e.URL.Query().Get("query")
	repo := createUsers()
	s := service.NewSearchUser(&repo)
	r, _ := s.Search(&query)

	_ = json.NewEncoder(w).Encode(r)
}

func createUsers() repository.UserRepositoryMemory {
	u := entity.NewUser("a672df0d-2927-43b0-bd33-50382ca86a20", "wesley vinicius", "viweesleyy")
	u2 := entity.NewUser("a672df0d-2927-43b0-bd33-50382ca86a20", "wesley vinicius", "viiweesleyy")
	repo := repository.UserRepositoryMemory{}
	repo.AddUser(&u)
	repo.AddUser(&u2)
	return repo
}
