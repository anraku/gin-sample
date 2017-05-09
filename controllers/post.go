package controllers
 
import (
    "../models"
)
 
// Post Model
type Post struct {
}
 
// Post Modelを返す
func NewPost() Post {
    return Post{}
}
 
// idに合致する記事の情報を返す
func (c Post) GetId(n int) interface{} {
    repo := models.NewPostRepository()
    post := repo.GetByPostID(n)
    return post
}

// 全記事の情報
func (c Post) GetAllPost() interface{} {
	repo := models.NewPostRepository()
    posts := repo.GetAllPost()
    return posts
}

// 全記事の件名
func (c Post) GetAllHeader() interface{} {
	repo := models.NewPostRepository()
    posts := repo.GetAllHeader()
    return posts
}

// 記事を投稿
func (c Post) CreatePost(header string, body string, author string) bool {
	repo := models.NewPostRepository()
	ok := repo.CreatePost(header, body, author)
	return ok
}

// 記事を更新
func (c Post) UpdatePost(id int, header string, body string) interface{} {
	repo := models.NewPostRepository()
	ok := repo.UpdatePost(id, header, body)
	return ok
}

// 記事を削除
func (c Post) DeletePost(n int) interface{} {
    repo := models.NewPostRepository()
    result := repo.DeletePost(n)
    return result
}
