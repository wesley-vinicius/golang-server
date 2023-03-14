package entity

type User struct {
	Uuid     string
	Name     string
	Username string
}

func NewUser(uuid, name, username string) User {
	return User{
		Uuid:     uuid,
		Name:     name,
		Username: username,
	}
}
