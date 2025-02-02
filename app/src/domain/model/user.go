package model

type User struct {
	id     int
	name   string
	apiKey string
}

func NewUser(id int, name, apiKey string) *User {
	return &User{
		id:     id,
		name:   name,
		apiKey: apiKey,
	}
}

func (u *User) Id() int {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) ApiKey() string {
	return u.apiKey
}
