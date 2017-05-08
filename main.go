package main

import (
    "controllers"
    "github.com/gin-gonic/gin"
    "reflect"
    "strconv"
    "net/http"
    "fmt"
)

// main ...
func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/*") // 事前にテンプレートをロード

    // トップページ
    router.GET("/index", func(c *gin.Context) {
        // データを処理する
        ctrl := controllers.NewPost()
        result := ctrl.GetAllPost()
        // テンプレートを使って、値を置き換えてHTMLレスポンスを応答
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "posts": result,
        })
    })

    // 掲示板の詳細表示
    router.GET("/post/get/:id", func(c *gin.Context) {
        // Pramを処理する
        n := c.Param("id")
        id, _ := strconv.Atoi(n)
        // データを処理する
        ctrl := controllers.NewPost()
        result := ctrl.GetId(id)
        // テンプレートを使って、値を置き換えてHTMLレスポンスを応答
        c.HTML(http.StatusOK, "post.tmpl", gin.H{
            "post": result,
        })
    })

    // 記事の投稿画面を表示
    router.GET("/post/new", func(c *gin.Context) {
        // テンプレートを使って、値を置き換えてHTMLレスポンスを応答
        c.HTML(http.StatusOK, "new.tmpl", gin.H{})
    })

    // 記事を投稿
    router.POST("/post/create", func(c *gin.Context) {
        header := c.PostForm("header")
        body := c.PostForm("body")
        author := c.PostForm("author")
        // 記事を投稿
        ctrl := controllers.NewPost()
        ok := ctrl.CreatePost(header, body, author)
        if ok == 0 {
            fmt.Printf("投稿成功\n")
        }
        // トップページを表示
        top := ctrl.GetAllPost()
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "posts": top,
        })
    })

    // 掲示板の編集
    router.GET("/post/edit/:id", func(c *gin.Context) {
        // Pramを処理する
        n := c.Param("id")
        id, _ := strconv.Atoi(n)
        // データを処理する
        ctrl := controllers.NewPost()
        result := ctrl.GetId(id)
        // テンプレートを使って、値を置き換えてHTMLレスポンスを応答
        c.HTML(http.StatusOK, "edit.tmpl", gin.H{
            "post": result,
        })
    })

    // 掲示板の更新
    router.POST("/post/update/:id", func(c *gin.Context) {
        // QueryStringを取得する
        n := c.Param("id")
        id, _ := strconv.Atoi(n)
        header := c.PostForm("header")
        body := c.PostForm("body")
        //fmt.Printf("main.go Value id: %v, header %v, body %v", id, header, body)
        // 件名と本文を更新する
        // TODO:後で書く
        ctrl := controllers.NewPost()
        ok := ctrl.UpdatePost(id, header, body)
        if ok == nil{
            fmt.Printf("エラーです\n")
        }
        // トップページを表示
        top := ctrl.GetAllPost()
        // テンプレートを使って、値を置き換えてHTMLレスポンスを応答
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "posts": top,
        })
    })

    // 記事の削除
    router.POST("/post/delete/:id", func(c *gin.Context) {
        // QueryStringを取得する
        n := c.Param("id")
        id, _ := strconv.Atoi(n)
        ctrl := controllers.NewPost()

        // 記事の削除
        ok := ctrl.DeletePost(id)
        if ok == 0{
            fmt.Printf("削除成功\n")
        }
        // トップページを表示
        top := ctrl.GetAllPost()
        // テンプレートを使って、値を置き換えてHTMLレスポンスを応答
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "posts": top,
        })
    })

    router.GET("/api/get/:id", func(c *gin.Context) {
        // Pramを処理する
        n := c.Param("id")
        id, err := strconv.Atoi(n)
        if err != nil {
            c.JSON(400, err)
            return
        }
        if id <= 0 {
            c.JSON(400, gin.H{"status": "id should be bigger than 0"})
            return
        }
        // データを処理する
        ctrl := controllers.NewPost()
        result := ctrl.GetId(id)
        if result == nil || reflect.ValueOf(result).IsNil() {
            c.JSON(404, gin.H{})
            return
        }
        c.JSON(200, result)
    })
    router.GET("/api/all", func(c *gin.Context) {
        // データを処理する
        ctrl := controllers.NewPost()
        result := ctrl.GetAllPost()
        if result == nil || reflect.ValueOf(result).IsNil() {
            c.JSON(404, gin.H{})
            return
        }
        c.JSON(200, result)
    })

    router.Run(":8080")
}