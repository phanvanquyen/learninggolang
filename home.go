package  main

import (
	"github.com/gin-gonic/gin"
	ctr "learninggolang/controllers"
	"log"
	"os"
	"path/filepath"
)

func homePage(c *gin.Context)  {
	log.Println("OK path")
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(dir)
	//c.String(http.StatusOK, "Hello")
	c.HTML(500,"index.html",nil)
}
func  main()  {

	log.Println("Start server  http  port 8080")
	serve :=gin.Default()
	serve.LoadHTMLGlob("templates/*.html")
	serve.GET("/",homePage)
	serve.GET("/about",ctr.About)
	serve.GET("/comic",ctr.Comic)
	serve.Run()
}