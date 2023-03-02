package model

import "gorm.io/gorm"

type Post struct {
	*gorm.Model
	UserId       uint
	Content      string
	LikeCount    int64
	CommentCount int64
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
