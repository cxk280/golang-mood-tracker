package controllers

import (
  "strconv"
  "net/http"
  "fmt"

  "up-and-down-mood-tracker/forms"
  "up-and-down-mood-tracker/models"

  "github.com/gin-gonic/gin"
  "github.com/gin-gonic/gin/binding"
)

// Define the AnalyticsController struct
type AnalyticsController struct{}

// Initialize a new analytics model
var analyticsModel = new(models.AnalyticsModel)

// Create new analytics
func (ctrl AnalyticsController) Create(c *gin.Context) {
  userID := getUserID(c)

  if userID == 0 {
    c.HTML(http.StatusOK, "error.html", gin.H{
      "errorMessage": "Please login first.",
    })
    c.Abort()
    return
  }

  var analyticsForm forms.AnalyticsForm

  // Make sure to use lower-case keys in the form data
  if err := c.ShouldBindWith(&analyticsForm, binding.Form); err != nil {
    c.HTML(http.StatusOK, "error.html", gin.H{
      "errorMessage": "Invalid form.",
    })
    c.Abort()
    return
  }

  analyticsID, err := analyticsModel.Create(userID, analyticsForm)

  if analyticsID > 0 && err != nil {
    c.HTML(http.StatusOK, "error.html", gin.H{
      "errorMessage": "Analytics could not be created.",
    })
    c.Abort()
    return
  }

  c.HTML(http.StatusOK, "error.html", gin.H{
      "errorMessage": "Analytics created.",
    })
}

// Get all analytics
func (ctrl AnalyticsController) All(c *gin.Context) {

  userID := getUserID(c)

  fmt.Println("userID: ",userID)

  if userID == 0 {
    c.HTML(http.StatusOK, "error.html", gin.H{
      "errorMessage": "Please login first.",
    })
    c.Abort()
    return
  }

  data, err := analyticsModel.All(userID)

  fmt.Println(data)

  if err != nil {
    c.HTML(http.StatusOK, "error.html", gin.H{
      "errorMessage": "Could not get the analytics.",
    })
    c.Abort()
    return
  }

  c.HTML(http.StatusOK, "analytics.html", gin.H{
  })
}

// Get one analytics
func (ctrl AnalyticsController) One(c *gin.Context) {
  userID := getUserID(c)

  if userID == 0 {
    c.HTML(http.StatusOK, "error.html", gin.H{
      "errorMessage": "Please login first.",
    })
    c.Abort()
    return
  }

  id := c.Param("id")

  if id, err := strconv.ParseInt(id, 10, 64); err == nil {

    data, err := analyticsModel.One(userID, id)
    if err != nil {
      c.HTML(http.StatusOK, "error.html", gin.H{
        "errorMessage": "Analytics not found.",
      })
      c.Abort()
      return
    }
    c.JSON(406, gin.H{"data": data})
  } else {
    c.HTML(http.StatusOK, "error.html", gin.H{
      "errorMessage": "Invalid parameter.",
    })
  }
}

// Update one analytics
func (ctrl AnalyticsController) Update(c *gin.Context) {
  userID := getUserID(c)

  if userID == 0 {
    c.HTML(http.StatusOK, "error.html", gin.H{
      "errorMessage": "Please login first.",
    })
    c.Abort()
    return
  }

  id := c.Param("id")
  if id, err := strconv.ParseInt(id, 10, 64); err == nil {

    var analyticsForm forms.AnalyticsForm

    if c.BindJSON(&analyticsForm) != nil {
      c.HTML(http.StatusOK, "error.html", gin.H{
        "errorMessage": "Invalid parameters.",
      })
      c.Abort()
      return
    }

    err := analyticsModel.Update(userID, id, analyticsForm)
    if err != nil {
      c.HTML(http.StatusOK, "error.html", gin.H{
        "errorMessage": "Analytics could not be updated.",
      })
      c.Abort()
      return
    }
    c.HTML(http.StatusOK, "error.html", gin.H{
      "errorMessage": "Analytics updated.",
    })
  } else {
    c.HTML(http.StatusOK, "error.html", gin.H{
      "errorMessage": "Invalid parameter.",
    })
  }
}

// Delete one analytics
func (ctrl AnalyticsController) Delete(c *gin.Context) {
  userID := getUserID(c)

  if userID == 0 {
    c.HTML(http.StatusOK, "error.html", gin.H{
      "errorMessage": "Please login first.",
    })
    c.Abort()
    return
  }

  id := c.Param("id")
  if id, err := strconv.ParseInt(id, 10, 64); err == nil {

    err := analyticsModel.Delete(userID, id)
    if err != nil {
      c.HTML(http.StatusOK, "error.html", gin.H{
        "errorMessage": "Analytics could not be deleted.",
      })
      c.Abort()
      return
    }
    c.HTML(http.StatusOK, "error.html", gin.H{
      "errorMessage": "Analytics deleted.",
    })
  } else {
    c.HTML(http.StatusOK, "error.html", gin.H{
      "errorMessage": "Invalid parameter.",
    })
  }
}
