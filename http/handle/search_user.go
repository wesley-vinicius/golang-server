package handle

import (
	"encoding/json"
	"github.com/wesley-vinicius/golang-server/core/service"
	"github.com/wesley-vinicius/golang-server/infra/repository"
	"net/http"
)

func SearchUserHandle(w http.ResponseWriter, e *http.Request) {
	w.Header().Set("Content-type", "json")

	query := e.URL.Query().Get("query")

	repo := repository.NewRepository()
	s := service.NewSearchUser(repo)
	r, err := s.Search(&query)

	if err != nil {
		println(err.Error())
	}

	_ = json.NewEncoder(w).Encode(r)
}
