package database

import (
	"testing"

	"github.com/kthanasio/golang/37-APIS/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		panic(err)
	}
	user, err := entity.NewUser("kleber", "mail@test.com", "123456")
	userDB := NewUser(db)
	err = userDB.Create(user)
	assert.Nil(t, err)

	var userFound *entity.User
	// userFound, err = userDB.FindByEmail("mail@test.com")
	err = db.Find(&userFound, "id = ?", user.ID).Error
	assert.NotNil(t, userFound)
	assert.Nil(t, err)
	assert.NotEmpty(t, userFound.ID)
	assert.NotEmpty(t, userFound.Name)
	assert.Equal(t, user.Name, userFound.Name)
	assert.NotEmpty(t, userFound.Email)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotEmpty(t, userFound.Password)

}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		panic(err)
	}
	user, err := entity.NewUser("kleber2", "mail2@test.com", "123456")
	userDB := NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	var userFound *entity.User
	userFound, err = userDB.FindByEmail("mail2@test.com")
	assert.NotNil(t, userFound)
	assert.Nil(t, err)
	assert.NotEmpty(t, userFound.ID)
	assert.NotEmpty(t, userFound.Name)
	assert.Equal(t, user.Name, userFound.Name)
	assert.NotEmpty(t, userFound.Email)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotEmpty(t, userFound.Password)
}

func TestFindByEmailNotFound(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		panic(err)
	}
	userDB := NewUser(db)

	var userNotFound *entity.User
	userNotFound, err = userDB.FindByEmail("mailNotFound@test.com")
	assert.Nil(t, userNotFound)
	assert.EqualError(t, err, "record not found")

	var emptyMailUser *entity.User
	emptyMailUser, err = userDB.FindByEmail("")
	assert.Nil(t, emptyMailUser)
	assert.EqualError(t, err, "record not found")

}
