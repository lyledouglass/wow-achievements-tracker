package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	err := os.Chdir("../templates")
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	r.LoadHTMLFiles("apiData.html")
	r.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "apiData.html", gin.H{
			"status": "success",
		})
	})
	r.Run()
}
