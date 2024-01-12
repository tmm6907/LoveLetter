package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/tmm6907/LoveLetter/models"
	"github.com/tmm6907/LoveLetter/routes/api"
)

var TEMPLATE_DIRS = [...]string{
	"templates/*.html",
	"templates/components/*.html",
	"templates/pages/*.html",
}

func ParseTemplates() (*template.Template, error) {
	files := []string{}
	for _, dir := range TEMPLATE_DIRS {
		ff, err := filepath.Glob(dir)
		if err != nil {
			return nil, err
		}
		files = append(files, ff...)
	}
	log.Println("Templates loaded:", files)
	return template.ParseFiles(files...)
}

func main() {
	server := gin.Default()
	tree, err := ParseTemplates()
	if err != nil {
		log.Fatalln(err)
	}
	server.Static("/static", "./static")
	server.SetHTMLTemplate(tree)
	port := os.Getenv("PORT")
	api.RegisterRoutes(server)
	server.GET("/", func(ctx *gin.Context) {
		title := "Will you go on a date with me?"
		body := "Test"
		options := []string{
			"Yes",
			"No",
		}
		letter, err := models.NewLetter(title, body)
		fmt.Println("Title", letter.Title)
		fmt.Println("Options", options)
		if err != nil {
			fmt.Println("Error!", err)
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": err,
			})
			return
		}
		ctx.HTML(http.StatusOK, "base", gin.H{
			"Letter": letter,
			"Yes":    options[0],
			"No":     options[1],
		})
	})
	server.Run(port)
}
