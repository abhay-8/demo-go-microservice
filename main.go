package main

import (
	// "net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")

	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(
	// 		http.StatusOK,
	// 		"index.html",
	// 		gin.H{
	// 			"title": "Home page",
	// 		},
	// 	)
	// })

	router.GET("/", showIndexPage)

	router.GET("/article/view/:article_id", getArticle)

	router.Run()
}
