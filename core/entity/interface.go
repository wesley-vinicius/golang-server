package entity

type UserReader interface {
	SearchByNameOrUserName(query *string) ([]User, error)
}

type UserWriter interface {
	AddUser(user *User)
}

type UserRepository interface {
	UserWriter
	UserReader
}
