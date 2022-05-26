package data

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUserData(t *testing.T) {
	userStore := userStore{Users: []User{}}
	user, err := userStore.Add(User{FirstName: "John", LastName: "Doe"})

	assert.Equal(t, User{FirstName: "John", LastName: "Doe", ID: "1"}, user)
	assert.Nil(t, err)
}

func TestAllData(t *testing.T) {
	userStore := userStore{Users: []User{}}
	userStore.Add(User{FirstName: "John", LastName: "Doe"})
	users := userStore.All()

	assert.Equal(t, userStore.Users, users)
}

func TestFindUser(t *testing.T) {
	userStore := userStore{Users: []User{}}
	userStore.Add(User{FirstName: "John", LastName: "Doe"})
	user, err := userStore.Find("1")

	assert.Equal(t, User{FirstName: "John", LastName: "Doe", ID: "1"}, user)
	assert.Nil(t, err)
}

func TestRemoveKnownUser(t *testing.T) {
	userStore := userStore{Users: []User{}}
	userStore.Add(User{FirstName: "John1", LastName: "Doe1"})
	userStore.Add(User{FirstName: "John2", LastName: "Doe2"})
	user, err := userStore.Remove("2")

	assert.Equal(t, User{FirstName: "John2", LastName: "Doe2", ID: "2"}, user)
	assert.Nil(t, err)
}

func TestRemoveUnKnownUser(t *testing.T) {
	userStore := userStore{Users: []User{}}
	userStore.Add(User{FirstName: "John1", LastName: "Doe1"})
	userStore.Add(User{FirstName: "John2", LastName: "Doe2"})
	user, err := userStore.Remove("20")

	assert.Equal(t, User{FirstName: "", LastName: "", ID: ""}, user)
	assert.NotNil(t, err)
}

func TestFindUnknownUser(t *testing.T) {
	userStore := userStore{Users: []User{}}
	_, err := userStore.Find("1")
	assert.NotNil(t, err)
}
