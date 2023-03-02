package mysql

import "forum/model"

func CreateUser(info *model.UserInfo) error {
	res := DB.Create(info)
	return res.Error
}
