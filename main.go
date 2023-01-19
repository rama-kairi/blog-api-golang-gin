package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rama-kairi/blog-api-golang-gin/controllers"
)

func GetBlogs(c *gin.Context) {
	c.JSON(200, map[string]any{
		"message": "Hello World",
		"status":  200,
	})
}

func main() {
	r := gin.Default()
	newBlogC := controllers.NewBlogStore()

	r.GET("/blog", newBlogC.GetAllBlogs)
	r.GET("/blog/:id", newBlogC.GetBlog)
	r.POST("/blog", newBlogC.CreateBlog)
	r.DELETE("/blog/:id", newBlogC.DeleteBlog)
	r.PATCH("/blog/:id", newBlogC.UpdateBlog)

	r.Run(":8080")
}
