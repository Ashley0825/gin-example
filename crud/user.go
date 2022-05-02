package crud

import "example/model"

func GetUser(userID string) model.UserInDB {
	for _, u := range model.Users {
		if u.ID == userID {
			return u
		}
	}
	return model.UserInDB{}
}
