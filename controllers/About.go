package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func About(c *gin.Context){

	c.String(http.StatusOK, "About Page")
}
func  Comic(c *gin.Context)  {
	c.String(http.StatusOK, "Comic Page")
}