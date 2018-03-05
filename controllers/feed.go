package controllers

import (
  // "strconv"
  "net/http"

  // "golang-mood-tracker/forms"
  "golang-mood-tracker/models"

  "github.com/gin-gonic/gin"
)

//FeedController ...
type FeedController struct{}

var feedModel = new(models.FeedModel)

Create ...
func (ctrl FeedController) Create(c *gin.Context) {
  userID := getUserID(c)

  if userID == 0 {
    c.JSON(403, gin.H{"message": "Please login first"})
    c.Abort()
    return
  }

  var feedForm forms.FeedForm

  //Make sure to use lower-case keys in the form data
  if err := c.ShouldBindWith(&feedForm, binding.Form); err != nil {
    c.JSON(406, gin.H{"message": "Invalid form", "form": feedForm})
    c.Abort()
    return
  }

  feedID, err := feedModel.Create(userID, feedForm)

  if feedID > 0 && err != nil {
    c.JSON(406, gin.H{"message": "feed could not be created", "error": err.Error()})
    c.Abort()
    return
  }

  c.JSON(200, gin.H{"message": "feed created", "id": feedID})
}

//All ...
func (ctrl FeedController) All(c *gin.Context) {
  userID := getUserID(c)

  if userID == 0 {
    c.JSON(403, gin.H{"message": "Please login first"})
    c.Abort()
    return
  }

  data, err := feedModel.All(userID)

  if err != nil {
    c.JSON(406, gin.H{"Message": "Could not get the feeds", "error": err.Error()})
    c.Abort()
    return
  }

  c.HTML(http.StatusOK, "analytics.html", gin.H{
    })
}

//One ...
func (ctrl FeedController) One(c *gin.Context) {
  userID := getUserID(c)

  if userID == 0 {
    c.JSON(403, gin.H{"message": "Please login first"})
    c.Abort()
    return
  }

  id := c.Param("id")

  if id, err := strconv.ParseInt(id, 10, 64); err == nil {

    data, err := feedModel.One(userID, id)
    if err != nil {
      c.JSON(404, gin.H{"Message": "feed not found", "error": err.Error()})
      c.Abort()
      return
    }
    c.JSON(200, gin.H{"data": data})
  } else {
    c.JSON(404, gin.H{"Message": "Invalid parameter"})
  }
}

//Update ...
func (ctrl FeedController) Update(c *gin.Context) {
  userID := getUserID(c)

  if userID == 0 {
    c.JSON(403, gin.H{"message": "Please login first"})
    c.Abort()
    return
  }

  id := c.Param("id")
  if id, err := strconv.ParseInt(id, 10, 64); err == nil {

    var feedForm forms.FeedForm

    if c.BindJSON(&feedForm) != nil {
      c.JSON(406, gin.H{"message": "Invalid parameters", "form": feedForm})
      c.Abort()
      return
    }

    err := feedModel.Update(userID, id, feedForm)
    if err != nil {
      c.JSON(406, gin.H{"Message": "feed could not be updated", "error": err.Error()})
      c.Abort()
      return
    }
    c.JSON(200, gin.H{"message": "feed updated"})
  } else {
    c.JSON(404, gin.H{"Message": "Invalid parameter", "error": err.Error()})
  }
}

//Delete ...
func (ctrl FeedController) Delete(c *gin.Context) {
  userID := getUserID(c)

  if userID == 0 {
    c.JSON(403, gin.H{"message": "Please login first"})
    c.Abort()
    return
  }

  id := c.Param("id")
  if id, err := strconv.ParseInt(id, 10, 64); err == nil {

    err := feedModel.Delete(userID, id)
    if err != nil {
      c.JSON(406, gin.H{"Message": "feed could not be deleted", "error": err.Error()})
      c.Abort()
      return
    }
    c.JSON(200, gin.H{"message": "feed deleted"})
  } else {
    c.JSON(404, gin.H{"Message": "Invalid parameter"})
  }
}
