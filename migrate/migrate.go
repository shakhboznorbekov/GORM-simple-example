package main

import (
	"github.com/GORM-simple-example/initializers"
	"github.com/GORM-simple-example/models"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main(){
	 // Migrate the schema
	 initializers.DB.AutoMigrate(&models.Post{})
}