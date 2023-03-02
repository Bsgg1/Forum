package mysql

import (
	"forum/model"
)

func CreateUser(info *model.UserInfo) error {
	res := DB.Create(info)
	return res.Error
}
func FindByName(username string) *model.UserInfo {
	var user model.UserInfo
	DB.First(&user, map[string]interface{}{
		"user_name": username,
	})
	return &user
}
