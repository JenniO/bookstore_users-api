package users

import (
	"encoding/json"
	"log"
)

type (
	PublicUser struct {
		Id          int64  `json:"id"`
		DateCreated string `json:"date_created"`
		Status      string `json:"status"`
	}
	PrivateUser struct {
		Id          int64  `json:"id"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		Email       string `json:"email"`
		DateCreated string `json:"date_created"`
		Status      string `json:"status"`
	}
)

func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for i, user := range users {
		result[i] = user.Marshall(isPublic)
	}
	return result
}

func (user *User) Marshall(isPublic bool) interface{} {
	// One way to do it
	if isPublic {
		return PublicUser{
			Id:          user.Id,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}
	// other way, if the json tags are the same
	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	if err := json.Unmarshal(userJson, &privateUser); err != nil {
		log.Println("fuck")
	}
	return privateUser
}
