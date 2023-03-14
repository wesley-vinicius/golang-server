package entity

type UserReader interface {
	SearchByNameOrUserName(query *string) ([]User, error)
}

type UserWriter interface {
	Create(user *User) (User, error)
}

type UserRepository interface {
	UserWriter
	UserReader
}
