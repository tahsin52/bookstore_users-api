package users

import (
	"fmt"
	"github.com/tahsin52/bookstore_users-api/utils/errors"
)

var(
	usersDB = make(map[int64]*User)
)

//func something()  {
//	user := User{}
//	if err := user.Get(); err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(user.FirstName)
//}

func (user *User) Get() *errors.RestErr  {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d Not Found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr  {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("Email %d Allready Exists", user.Email))

		}
		return errors.NewBadRequestError(fmt.Sprintf("User %d Allready Exists", user.Id))
	}
	usersDB[user.Id] = user
	return nil
}