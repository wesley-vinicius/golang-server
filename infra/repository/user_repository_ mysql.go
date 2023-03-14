package repository

import (
	"database/sql"
	"github.com/software-engineer-challenge/core/entity"
	"github.com/software-engineer-challenge/infra/database"
)

type UserRepositoryMysql struct {
	db *sql.DB
}

func NewRepository() *UserRepositoryMysql {
	return &UserRepositoryMysql{
		db: database.Connect(),
	}
}

func (r *UserRepositoryMysql) SearchByNameOrUserName(query *string) ([]entity.User, error) {
	stmt, err := r.db.Prepare(`SELECT uuid, name, user_name FROM users where name LIKE ? OR user_name LIKE ?`)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query("%"+*query+"%", "%"+*query+"%")
	if err != nil {
		return nil, err
	}
	var users []entity.User

	for rows.Next() {
		u := entity.User{}
		err := rows.Scan(&u.Uuid, &u.Name, &u.Username)
		if err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	return users, nil
}

func (r *UserRepositoryMysql) Create(user *entity.User) (entity.User, error) {
	stmt, err := r.db.Prepare(`INSERT INTO users (uuid, name, user_name) values (?, ?, ?)`)
	if err != nil {
		return *user, err
	}

	_, err = stmt.Exec(user.Uuid, user.Name, user.Username)
	if err != nil {
		return entity.User{}, err
	}

	return *user, nil
}
