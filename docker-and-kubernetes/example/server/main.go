package main
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", DefaultHandler)

	if err := router.Run(); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}