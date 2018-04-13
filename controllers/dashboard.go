package controllers

import (
  "net/http"
  "fmt"

  "golang-mood-tracker/models"

  "github.com/gin-gonic/gin"
)

// Define the DashboardController struct
type DashboardController struct{}

// Initialize a new dashboard model
var dashboardModel = new(models.DashboardModel)

// Open dashboard
func (ctrl DashboardController) All(c *gin.Context) {
  userID := getUserID(c)

  if userID == 0 {
    c.HTML(http.StatusOK, "error.html", gin.H{
      "errorMessage": "Please login first.",
    })
    c.Abort()
    return
  }

  c.HTML(http.StatusOK, "dashboard.html", gin.H{
    })
}
