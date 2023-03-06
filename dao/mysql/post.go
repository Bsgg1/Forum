package mysql

import "forum/model"

func NewPost(post *model.Post) error {
	return DB.Create(post).Error
}
func ListPost(userid uint) ([]string, error) {
	var posts []string
	err := DB.Table(`post`).Select("content").Where(`user_id = ?`, userid).Find(&posts).Error
	return posts, err
}
