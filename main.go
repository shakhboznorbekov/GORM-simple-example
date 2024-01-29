package main

import (
	"github.com/gin-gonic/gin"
	"github.com/GORM-simple-example/controllers"
	"github.com/GORM-simple-example/initializers"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main(){
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.GetPostsAll)
	r.GET("/posts/:id", controllers.GetPostByID)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)



	r.Run()
}