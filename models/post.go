package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

var db *gorm.DB

func init() {
  conn, err := gorm.Open("mysql", "root:password@/gin_sample?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
  	panic(err)
  }
  db = conn
}

// 投稿情報の構造体
type Post struct {
	ID     int         `gorm:"AUTO_INCREMENT;primary_key"`
    Header string
    Body   string
    Author string
    create_time string
}

// Repository
type PostRepository struct {
}

// new PostRepository
func NewPostRepository() PostRepository {
    return PostRepository{}
}

func (m PostRepository) GetByPostID(id int) *Post {
	var post Post
	return &post
}

// 記事を全検索
func (m PostRepository) GetAllPost() []Post {
	var post []Post
	db.Select("").Find(&post)
	fmt.Printf("post: %v", post)
	return post
}

// 全記事の件名
func (m PostRepository) GetAllHeader() []Post {
	var post []Post
	return post
}

// 記事を投稿する
func (m PostRepository) CreatePost(header string, body string, author string) interface{} {
	return nil
}

// 記事の件名と本文を更新する
func (m PostRepository) UpdatePost(id int, header string, body string) interface{} {
	return nil
}

// 記事を削除する
func (m PostRepository) DeletePost(id int) interface{} {
	return nil
}