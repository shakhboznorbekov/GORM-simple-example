package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/GORM-simple-example/initializers"
	"github.com/GORM-simple-example/models"
)

func PostsCreate(c *gin.Context){

	//Get data off req body
	var body struct{
		Body string
		Title string
	}
	c.Bind(&body)
	//Create post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil{
		c.Status(400)
		return
	}
	//Return it
	c.JSON(200, gin.H{
		"post": post,
	})
} 

func GetPostsAll(c *gin.Context){
	//Get the post
	var posts []models.Post
	initializers.DB.Find(&posts)

	//Respond with them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPostByID(c *gin.Context){
	//Get id off url
	id := c.Param("id")
	//Get the post
	var post models.Post 
	initializers.DB.Find(&post, id)

	//Respond with them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsUpdate(c *gin.Context){
	//Get id off url
	id := c.Param("id")

	//Get data off req body
	var body struct{
		Body  string
		Title string
	}

	c.Bind(&body)

	//Find the post were updating
	var post models.Post 
	initializers.DB.First(&post, id)

	//Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,  
		Body: body.Body,
	})
	//Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsDelete(c *gin.Context){
	//Get if off the url
	id := c.Param("id")

	//Delete the posts
	initializers.DB.Delete(&models.Post{}, id)

	//Respond it
	c.Status(200)
}