package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

//IndexController ...
type IndexController struct{}

// Open index
func (ctrl IndexController) All(c *gin.Context) {
  c.HTML(http.StatusOK, "index.html", gin.H{
      })
}