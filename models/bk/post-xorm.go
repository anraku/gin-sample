package models

import (
    "github.com/go-xorm/xorm"
    _ "github.com/go-sql-driver/mysql"
    "fmt"
)

var engine *xorm.Engine
// init ...
func init() {
    var err error
    engine, err = xorm.NewEngine("mysql", "root:password@/gin_sample")
    if err != nil {
        panic(err)
    }
}

// Post is
type Post struct {
    ID       int 		`json:"id" xorm:"'id'"`
    Header string		`json:"header" xorm:"'header'"`
    Body   string       `json:"body" xorm:"'body'"`
    Author string       `json:"author" xorm:"'author'"`
    Create_time string  `json:"create_time" xorm:"'create_time'"`
}

// NewPost ...
func NewPost(id int, header string, body string, author string, create_time string) Post {
    return Post{
        ID:       id,
        Header: header,
        Body: body,
        Author: author,
        Create_time: create_time,
    }
}

// UserRepository is
type PostRepository struct {
}

// NewUserRepository ...
func NewPostRepository() PostRepository {
    return PostRepository{}
}

// GetByID ...
func (m PostRepository) GetByPostID(id int) *Post {
    var post = Post{ID:id}
    has, _:= engine.Get(&post)
    if has {
        fmt.Println(post)
        return &post
    }
    return nil
}

// 記事を全検索
func (m PostRepository) GetAllPost() []Post {
    var posts []Post
    if err := engine.Find(&posts); err != nil {
        return nil
    }
    return posts
}

// 全記事の件名
func (m PostRepository) GetAllHeader() []Post {
    var headers []Post
    if err := engine.Cols("header").Find(&headers); err != nil {
        return nil
    }
    fmt.Println(headers)
    return headers
}

// 記事を投稿する
func (m PostRepository) CreatePost(header string, body string, author string) interface{} {
    var post = Post{Header: header, Body: body, Author: author}
    fmt.Println(post)
    affected, err := engine.Insert(&post)
    if err != nil {
        return nil
    }
    return affected
}

// 記事の件名と本文を更新する
func (m PostRepository) UpdatePost(id int, header string, body string) interface{} {
    var post = Post{ID:id, Header: header, Body: body}
    //fmt.Printf("Value id: %v, header %v, body %v", id, header, body)
    affected, err := engine.Update(&post, &Post{ID:id})
    if err != nil {
        return nil
    }
    return affected
}

// 記事を削除する
func (m PostRepository) DeletePost(id int) interface{} {
    post := Post{ID: id}
    ok, err := engine.ID(id).Delete(&post)
    if err != nil {
        return err
    }
    return ok
}
