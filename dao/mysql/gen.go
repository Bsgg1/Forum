package mysql

import "forum/model"

func GenModel() {
	DB.AutoMigrate(&model.UserInfo{})
	DB.AutoMigrate(&model.UserRelation{})
	DB.AutoMigrate(&model.CommentRelation{})
	DB.AutoMigrate(&model.PostRelation{})
	DB.AutoMigrate(&model.Comment{})
	DB.AutoMigrate(&model.Post{})
	DB.AutoMigrate(&model.PostLike{})
}
