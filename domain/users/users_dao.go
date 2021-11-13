package users

import (
	"fmt"
	errors "github.com/lawrenceMichaelMargaja/bookstore_users_api/utils"
)

var(
	userDB = make(map[int64]*User)
)

func something() {
	user := User{}

	if err := user.Get(); err != nil {
		fmt.Println(err)
		return
	}
}

func (user *User) Get() *errors.RestErr {
	result := userDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found.", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
 	return nil
}

func (user * User) Save() *errors.RestErr {
	current := userDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("Email %s is already registered.", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("User %d already exists.", user.Id))
	}
	userDB[user.Id] = user
	return nil
}