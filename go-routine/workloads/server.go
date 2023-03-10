package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()
	server.MaxMultipartMemory = 8 << 20 // 8MB
	server.GET("/", func(ctx *gin.Context) {
		response, err := http.Get("https://restcountries.com/v3.1/all")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
		}
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Error": err.Error(),
			})
		}
		file, err := os.Create("hoax.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		_, err = file.WriteString(string(body))
		ctx.JSON(http.StatusOK, gin.H{
			"Data": string(body),
		})
	})

	server.Run(":9090")

}
