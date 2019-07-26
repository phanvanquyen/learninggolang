package  main

import (
	"github.com/gin-gonic/gin"
	cnf "learninggolang/config"
	ctr "learninggolang/controllers"
	app "learninggolang/libs"
	middleware "learninggolang/middeware"
	"log"
)

func homePage(c *gin.Context)  {
	//c.String(http.StatusOK, "Hello")
	 db := app.GetDb();
	log.Println(db);
	arr := []string{"Ana","Davis","Coco","Thomas"}
	c.HTML(500,"index.html",gin.H{"listname":arr,"title":" World"})
}
func  main()  {
	log.Println(cnf.DB_USERNAME)
	log.Println("Start server  http  port 8080")
	serve :=gin.Default()
	serve.LoadHTMLGlob("templates/*.html")
	serve.Use(middleware.IsLogined())
	serve.GET("/",homePage)
	serve.GET("/about",ctr.About)
	serve.GET("/comic",ctr.Comic)
	app.Views(main);
	serve.Run()
}