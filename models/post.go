package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var db *gorm.DB

func init() {
  conn, err := gorm.Open("mysql", "root:password@/gin_sample?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
  	panic(err)
  }
  db = conn
  //DB Migrate
  if !db.HasTable("posts") {
  	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Post{})
  }
}

// 投稿情報の構造体
type Post struct {
	ID     int            `gorm:"AUTO_INCREMENT;primary_key"`
    Header string         `gorm:"not null;size:255"`
    Body   string         `gorm:"not null;size:13000"`
    Author string         `gorm:"not null;size:30"`
    CreatedAt time.Time   `gorm:"not null"`
}

// Repository
type PostRepository struct {
}

// Repositoryを返す
func NewPostRepository() PostRepository {
    return PostRepository{}
}

// idに合致する記事を取得
func (m PostRepository) GetByPostID(id int) *Post {
	var post Post
	db.Where(Post{ID: id}).Find(&post)
	return &post
}

// 記事を全検索
func (m PostRepository) GetAllPost() []Post {
	var post []Post
	db.Select("*").Find(&post)
	return post
}

// 記事を投稿する
func (m PostRepository) CreatePost(header string, body string, author string) bool {
	post := Post {
		Header: header,
		Body:   body,
	}
	db.Create(&post)
	return true
}

// 記事の件名と本文を更新する
func (m PostRepository) UpdatePost(id int, header string, body string) interface{} {
	post := Post{ID: id}
	db.Model(&post).Updates(map[string]interface{}{
		"Header":header, 
		"Body":  body,
		})
	return true
}

// 記事を削除する
func (m PostRepository) DeletePost(id int) interface{} {
	post := Post{ID: id}
	db.Delete(post)
	return nil
}