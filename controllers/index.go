package controllers

import (
  "net/http"

  "golang-mood-tracker/models"

  "github.com/gin-gonic/gin"
)

//IndexController ...
type IndexController struct{}

// Initialize a new index model
var indexModel = new(models.IndexModel)


// Open index
func (ctrl IndexController) All(c *gin.Context) {
  c.HTML(http.StatusOK, "index.html", gin.H{
      })
}