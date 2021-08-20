package main

import (
	"Details/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	fmt.Println("hello")
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/book_detail", func(c *gin.Context) {

			name := c.GetHeader("name")
			book := utils.GetBookByName(name)
			c.JSON(http.StatusOK, gin.H{"status": "success", "book": &book})

		})
		api.GET("/book_list", func(c *gin.Context) {

		books := utils.GetBooks()

		c.JSON(http.StatusOK, gin.H{"status": "success", "books": &books})

		})

	}
	_ = router.Run(":8082")


}
