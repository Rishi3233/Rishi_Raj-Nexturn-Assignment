package handlers

import (
	"assignment_blog_system/config"
	"assignment_blog_system/models"
	
	"net/http"
	
	"github.com/gin-gonic/gin"
)

// Handler to create a blog post
func CreateBlog(c *gin.Context) {
	var blog models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	stmt, err := config.DB.Prepare("INSERT INTO blogs(title, content, author) VALUES(?, ?, ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error preparing the query"})
		return
	}

	_, err = stmt.Exec(blog.Title, blog.Content, blog.Author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error inserting into the database"})
		return
	}

	c.JSON(http.StatusCreated, blog)
}

// Handler to get a blog post by ID
func GetBlog(c *gin.Context) {
	id := c.Param("id")
	var blog models.Blog

	row := config.DB.QueryRow("SELECT id, title, content, author, timestamp FROM blogs WHERE id = ?", id)
	err := row.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	c.JSON(http.StatusOK, blog)
}

// Handler to get all blogs
func GetBlogs(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, title, content, author, timestamp FROM blogs ORDER BY timestamp DESC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching blogs"})
		return
	}
	defer rows.Close()

	var blogs []models.Blog
	for rows.Next() {
		var blog models.Blog
		err := rows.Scan(&blog.ID, &blog.Title, &blog.Content, &blog.Author, &blog.Timestamp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning blog data"})
			return
		}
		blogs = append(blogs, blog)
	}

	c.JSON(http.StatusOK, blogs)
}

// Handler to update a blog post by ID
func UpdateBlog(c *gin.Context) {
	id := c.Param("id")
	var blog models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	stmt, err := config.DB.Prepare("UPDATE blogs SET title = ?, content = ?, author = ? WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error preparing update query"})
		return
	}
	_, err = stmt.Exec(blog.Title, blog.Content, blog.Author, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating blog"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Blog updated"})
}

// Handler to delete a blog post by ID
func DeleteBlog(c *gin.Context) {
	id := c.Param("id")

	stmt, err := config.DB.Prepare("DELETE FROM blogs WHERE id = ?")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error preparing delete query"})
		return
	}
	_, err = stmt.Exec(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting blog"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "Blog deleted"})
}
