package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)
func d(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"dats": "hello world"})    
}

func main() {
	r := gin.Default()
	r.GET("/", d)
	r.Run()
}
