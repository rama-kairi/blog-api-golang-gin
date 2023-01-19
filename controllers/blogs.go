package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rama-kairi/blog-api-golang-gin/utils"
)

type Blog struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type BlogStore struct {
	Blogs []Blog
}

func NewBlogStore() *BlogStore {
	return &BlogStore{
		Blogs: []Blog{},
	}
}

// Get all blogs
func (t BlogStore) GetAllBlogs(c *gin.Context) {
	t.loadFromJson()

	// Write the json to the response
	utils.Response(c, http.StatusOK, t.Blogs, "Blogs found")
}

// Get a blog
func (t BlogStore) GetBlog(c *gin.Context) {
	t.loadFromJson()

	// Get blog id from url
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error getting blog")
		return
	}

	for _, blog := range t.Blogs {
		if blog.Id == id {
			blogJson, err := json.Marshal(blog)
			if err != nil {
				log.Fatal(err)
			}
			utils.Response(c, http.StatusOK, blogJson, "Blog found")
			return
		}
	}
	utils.Response(c, http.StatusNotFound, nil, "Blog not found")
}

// Create a blog
func (t BlogStore) CreateBlog(c *gin.Context) {
	t.loadFromJson()

	var blog Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error creating blog")
		return
	}

	blog.Id = t.newTodoId()

	// Append the blog to the slice
	t.Blogs = append(t.Blogs, blog)
	t.saveToJson()

	// Marshal the blog into json
	utils.Response(c, http.StatusCreated, blog, "Blog created successfully")
}

// Delete a blog
func (t BlogStore) DeleteBlog(c *gin.Context) {
	t.loadFromJson()

	// Get blog id from url
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error getting blog")
		return
	}

	for i, blog := range t.Blogs {
		if blog.Id == id {
			t.Blogs = append(t.Blogs[:i], t.Blogs[i+1:]...)
			t.saveToJson()
			utils.Response(c, http.StatusNoContent, nil, "Blog deleted successfully")
			return
		}
	}

	// Save the blogs to the json file
	t.saveToJson()

	// If the blog is not found, return 404
	utils.Response(c, http.StatusNotFound, nil, "Blog not found")
}

// Update a blog
func (t BlogStore) UpdateBlog(c *gin.Context) {
	t.loadFromJson()

	// Get blog id from url
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error getting blog")
		return
	}

	// Get the blog from the request body
	var blog Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		utils.Response(c, http.StatusBadRequest, nil, "Error updating blog")
		return
	}

	for i, b := range t.Blogs {
		if b.Id == id {
			blog.Id = id
			t.Blogs[i] = blog
			t.saveToJson()
			utils.Response(c, http.StatusOK, blog, "Blog updated successfully")
			return
		}
	}

	// If the blog is not found, return 404
	utils.Response(c, http.StatusNotFound, nil, "Blog not found")
}
