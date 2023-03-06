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
func UpdateUserInfo(user *model.UserInfo) {
	DB.Save(user)
}
func UpdateRelation(user1, user2 string, status int) error {
	var relation model.UserRelation

	ft := DB.Find(&relation, map[string]interface{}{
		"user1": user1,
		"user2": user2,
	})
	if ft.Error != nil {
		return ft.Error
	}
	if relation.User1 == "" {
		return DB.Create(&model.UserRelation{
			User1:  user1,
			User2:  user2,
			Status: int8(status),
		}).Error
	} else {
		relation.Status = int8(status)
		return DB.Save(relation).Error
	}
}
