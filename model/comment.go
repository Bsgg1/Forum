package model

import "gorm.io/gorm"

type Comment struct {
	*gorm.Model
	UserId    uint
	Content   string
	PostId    uint
	LikeCount uint
}
type CommentRelation struct {
	*gorm.Model
	CommentId1 uint
	CommentId2 uint
}

func (Comment) TableName() string {
	return "comment"
}
func (CommentRelation) TableName() string {
	return "comment_relation"
}
