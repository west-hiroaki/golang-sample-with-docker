package web

import (
	"strconv"

	// ローカルパッケージパス
	"app/enum"
	"app/model/todo"

	"github.com/gin-gonic/gin"
)

func redirectRoot(c *gin.Context) {
	c.Redirect(302, "/todo/")
}

// TodoRoot : ルート HTML 表示
func TodoRoot(c *gin.Context) {
	todos := todo.All()
	c.HTML(200, "index.tmpl", gin.H{
		"todos": todos,
	})
}

// TodoPost : データ新規登録
func TodoPost(c *gin.Context) {
	text := c.PostForm("text")

	statusInt, _ := strconv.Atoi(c.PostForm("status"))
	status := enum.Status(statusInt)

	todo.Insert(text, status)

	redirectRoot(c)
}

// TodoDetail : 詳細データ表示
func TodoDetail(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		redirectRoot(c)
	}
	todo := todo.FindByID(id)
	c.HTML(200, "detail.tmpl", gin.H{"todo": todo})
}

// TodoUpdate : データ更新
func TodoUpdate(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		redirectRoot(c)
	}

	text := c.PostForm("text")

	statusInt, _ := strconv.Atoi(c.PostForm("status"))
	status := enum.Status(statusInt)

	todo.Update(id, text, status)

	redirectRoot(c)
}

// TodoDeleteCheck : データ削除チェック
func TodoDeleteCheck(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		redirectRoot(c)
	}
	todo := todo.FindByID(id)
	c.HTML(200, "delete.tmpl", gin.H{"todo": todo})
}

// TodoDelete : データ削除
func TodoDelete(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		redirectRoot(c)
	}
	todo.Delete(id)
	redirectRoot(c)
}
