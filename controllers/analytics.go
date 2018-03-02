package controllers

import (
  "strconv"

  "golang-mood-tracker/forms"
  "golang-mood-tracker/models"

  "github.com/gin-gonic/gin"
)

//AnalyticsController ...
type AnalyticsController struct{}

var analyticsModel = new(models.AnalyticsModel)

//Create ...
func (ctrl AnalyticsController) Create(c *gin.Context) {
  userID := getUserID(c)
f
  if userID == 0 {
    c.JSON(403, gin.H{"message": "Please login first"})
    c.Abort()
    return
  }

  var analyticsForm forms.AnalyticsForm

  if c.BindJSON(&analyticsForm) != nil {
    c.JSON(406, gin.H{"message": "Invalid form boo", "form": analyticsForm})
    c.Abort()
    return
  }

  analyticsID, err := analyticsModel.Create(userID, analyticsForm)

  if analyticsID > 0 && err != nil {
    c.JSON(406, gin.H{"message": "analytics could not be created", "error": err.Error()})
    c.Abort()
    return
  }

  c.JSON(200, gin.H{"message": "analytics created", "id": analyticsID})
}

//All ...
func (ctrl AnalyticsController) All(c *gin.Context) {
  userID := getUserID(c)

  if userID == 0 {
    c.JSON(403, gin.H{"message": "Please login first"})
    c.Abort()
    return
  }

  data, err := analyticsModel.All(userID)

  if err != nil {
    c.JSON(406, gin.H{"Message": "Could not get the analytics", "error": err.Error()})
    c.Abort()
    return
  }

  c.JSON(200, gin.H{"data": data})
}

//One ...
func (ctrl AnalyticsController) One(c *gin.Context) {
  userID := getUserID(c)

  if userID == 0 {
    c.JSON(403, gin.H{"message": "Please login first"})
    c.Abort()
    return
  }

  id := c.Param("id")

  if id, err := strconv.ParseInt(id, 10, 64); err == nil {

    data, err := analyticsModel.One(userID, id)
    if err != nil {
      c.JSON(404, gin.H{"Message": "analytics not found", "error": err.Error()})
      c.Abort()
      return
    }
    c.JSON(200, gin.H{"data": data})
  } else {
    c.JSON(404, gin.H{"Message": "Invalid parameter"})
  }
}

//Update ...
func (ctrl AnalyticsController) Update(c *gin.Context) {
  userID := getUserID(c)

  if userID == 0 {
    c.JSON(403, gin.H{"message": "Please login first"})
    c.Abort()
    return
  }

  id := c.Param("id")
  if id, err := strconv.ParseInt(id, 10, 64); err == nil {

    var analyticsForm forms.AnalyticsForm

    if c.BindJSON(&analyticsForm) != nil {
      c.JSON(406, gin.H{"message": "Invalid parameters", "form": analyticsForm})
      c.Abort()
      return
    }

    err := analyticsModel.Update(userID, id, analyticsForm)
    if err != nil {
      c.JSON(406, gin.H{"Message": "analytics could not be updated", "error": err.Error()})
      c.Abort()
      return
    }
    c.JSON(200, gin.H{"message": "analytics updated"})
  } else {
    c.JSON(404, gin.H{"Message": "Invalid parameter", "error": err.Error()})
  }
}

//Delete ...
func (ctrl AnalyticsController) Delete(c *gin.Context) {
  userID := getUserID(c)

  if userID == 0 {
    c.JSON(403, gin.H{"message": "Please login first"})
    c.Abort()
    return
  }

  id := c.Param("id")
  if id, err := strconv.ParseInt(id, 10, 64); err == nil {

    err := analyticsModel.Delete(userID, id)
    if err != nil {
      c.JSON(406, gin.H{"Message": "analytics could not be deleted", "error": err.Error()})
      c.Abort()
      return
    }
    c.JSON(200, gin.H{"message": "analytics deleted"})
  } else {
    c.JSON(404, gin.H{"Message": "Invalid parameter"})
  }
}
