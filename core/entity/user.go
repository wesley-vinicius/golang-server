package entity

type User struct {
	uuid     string
	name     string
	username string
}

func NewUser(uuid, name, username string) User {
	return User{
		uuid:     uuid,
		name:     name,
		username: username,
	}
}

func (u User) Uuid() string {
	return u.uuid
}

func (u User) Name() string {
	return u.name
}

func (u User) Username() string {
	return u.username
}
