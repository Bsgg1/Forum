package model

import "gorm.io/gorm"

type Post struct {
	*gorm.Model
	UserId    uint   `json:"userid"`
	Content   string `json:"content"`
	LikeCount int64  `json:"likecount"`
}
type PostRelation struct {
	*gorm.Model
	UserId uint
	PostId uint
}
type PostLike struct {
	*gorm.Model
	UserId uint
	PostId uint
}

func (Post) TableName() string {
	return "post"
}
func (PostRelation) TableName() string {
	return "post_relation"
}
func (PostLike) TableName() string {
	return "post_like"
}
