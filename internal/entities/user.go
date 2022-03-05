package entities

import "time"

type User struct {
	// baseProperties
	id        string    `json:id`
	email     string    `json:email`
	password  string    `json:password`
	createdAt time.Time `json:createdAt`
	updatedAt time.Time `json:updatedAt`
}

func (u User) GetID() string {
	return u.id
}
func (u User) GetEmail() string {
	return u.email
}

func (u User) PasswordCompare(password string) bool {
	//TODO added
	return password == u.password
}
