package controller

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tckthecreator/clean_arch_go/model"
)

// needs update
var db *sql.DB

func createPost(c *gin.Context) {
	var newPost model.Post
	if err := c.BindJSON(&newPost); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	_, err := db.ExecContext(context.Background(),
		"INSERT INTO posts (title, body, owner_id, created_at, updated_at, deleted_at) VALUES ($1, $2, $3, $4, $5, $6)",
		newPost.Title, newPost.Body, newPost.OwnerID, newPost.CreatedAt, newPost.UpdatedAt, newPost.DeletedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, err.Error())
}

func getPostsByOwner(c *gin.Context) {
	ownerID := c.Param("ownerID")

	rows, err := db.QueryContext(context.Background(),
		"SELECT title, body, owner_id, created_at, updated_at, deleted_at FROM posts WHERE owner_id = $1", ownerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	defer rows.Close()

	var posts []model.Post
	for rows.Next() {
		var post model.Post
		err := rows.Scan(&post.Title, &post.Body, &post.OwnerID, &post.CreatedAt, &post.UpdatedAt, &post.DeletedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		posts = append(posts, post)
	}

	c.JSON(http.StatusOK, posts)
}

func getPostByID(c *gin.Context) {
	id := c.Param("id")
	var post model.Post

	err := db.QueryRowContext(context.Background(),
		"SELECT title, body, owner_id, created_at, updated_at, deleted_at FROM posts WHERE id = $1", id).
		Scan(&post.Title, &post.Body, &post.OwnerID, &post.CreatedAt, &post.UpdatedAt, &post.DeletedAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, post)
}

func updateFields(fields []string) string {
	updates := make([]string, len(fields))
	for i, field := range fields {
		updates[i] = field + "=$" + fmt.Sprint(i+2)
	}
	return fmt.Sprintf("%s", updates)
}

func updatePostTitleOrBody(c *gin.Context) {
	id := c.Param("id")
	var updateData map[string]string
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	title, titleExists := updateData["title"]
	body, bodyExists := updateData["body"]

	if !titleExists && !bodyExists {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No valid update data provided"})
		return
	}

	var updatedFields []string
	if titleExists {
		updatedFields = append(updatedFields, "title")
	}
	if bodyExists {
		updatedFields = append(updatedFields, "body")
	}

	_, err := db.ExecContext(context.Background(),
		fmt.Sprintf("UPDATE posts SET %s WHERE id = $1", updateFields(updatedFields)),
		id, title, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(200, gin.H{"message": "Post updated successfully"})
}

func deletePost(c *gin.Context) {
	id := c.Param("id")
	_, err := db.ExecContext(context.Background(),
		"UPDATE posts SET deleted_at = NOW() WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, err.Error())
}

func postsEndPoints(r *gin.Engine) {
	go r.POST("/post", createPost)
	go r.GET("/posts/:ownerID", getPostsByOwner)
	go r.GET("/posts/:id", getPostByID)
	go r.PATCH("/posts/:id", updatePostTitleOrBody)
	go r.DELETE("/posts/:id", deletePost)
}
