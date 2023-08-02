package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	data := User{
		Name:     "Username",
		Email:    "user@mail.com",
		Password: "123456",
	}
	u, err := NewUser(data.Name, data.Email, data.Password)
	assert.Nil(t, err)
	assert.NotNil(t, u)
	assert.NotEmpty(t, u.ID)
	assert.NotEmpty(t, u.Name)
	assert.NotEmpty(t, u.Email)
	assert.NotEmpty(t, u.Password)
	assert.Equal(t, data.Name, u.Name)
	assert.Equal(t, data.Email, u.Email)
}

func Test_ValidatePassword(t *testing.T) {
	data := User{
		Name:     "Username",
		Email:    "user@mail.com",
		Password: "123456",
	}
	u, err := NewUser(data.Name, data.Email, data.Password)
	assert.Nil(t, err)
	assert.True(t, u.ValidatePassword(data.Password))
	assert.False(t, u.ValidatePassword(data.Password+"error"))
	assert.NotEqual(t, data.Password, u.Password)

	data.Password = "01234567890123456789012345678901234567890123456789012345678901234567890123456789"
	u, err = NewUser(data.Name, data.Email, data.Password)
	assert.Nil(t, u)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "bcrypt: password length exceeds 72 bytes")

}
