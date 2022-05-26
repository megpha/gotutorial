package data

import (
	"errors"
	"fmt"
)

type Storage interface {
	Add(user User) (User, error)
	Find(id string) (User, error)
	Remove(id string) (User, error)
	All() []User
}

type userStore struct {
	Users []User
}

func NewStorage() Storage {
	return &userStore{}
}

func (u *userStore) Add(user User) (User, error) {

	if error := user.Validate(); error != nil {
		return User{}, errors.New("Invalid User Information")
	}

	user.ID = fmt.Sprintf("%d", len(u.Users)+1)
	u.Users = append(u.Users, user)

	return user, nil
}

func (u *userStore) All() []User {
	return u.Users
}

func (u *userStore) Find(id string) (User, error) {
	for _, user := range u.Users {
		if user.ID == id {
			return user, nil
		}
	}

	return User{}, errors.New("No User Found")
}

func (u *userStore) Remove(id string) (User, error) {
	for index, user := range u.Users {
		if user.ID == id {
			u.Users = append(u.Users[:index], u.Users[index+1:]...)
			return user, nil

		}
	}

	return User{}, errors.New("No User Found")
}

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	ID        string `json:"id"`
}

func (u User) Validate() error {
	if u.FirstName == "" {
		return errors.New("Empty First Name")
	}
	if u.LastName == "" {
		return errors.New("Empty Last Name")
	}

	return nil
}
