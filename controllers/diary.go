package controllers

import (
  "strconv"
  "net/http"

  "golang-mood-tracker/forms"
  "golang-mood-tracker/models"

  "github.com/gin-gonic/gin"
)

//DiaryController ...
type DiaryController struct{}

var diaryModel = new(models.DiaryModel)

// //Create ...
func (ctrl DiaryController) Create(c *gin.Context) {
  userID := getUserID(c)

  if userID == 0 {
    c.JSON(403, gin.H{"message": "Please login first"})
    c.Abort()
    return
  }

  var diaryForm forms.DiaryForm

  //If using bindJSON, form must be submitted as raw JSON (type application/json) in order for c.BindJSON below to work
  //Example:
  //{
  //   "Title": "MyTitle",
  //   "Content": "MyContent"
  // }

  // if c.BindJSON(&diaryForm) != nil {
    // c.JSON(406, gin.H{"message": "Invalid form", "form": diaryForm})
    // c.Abort()
    // return
  // }

  titleValue := c.PostForm("Title");
  contentValue := c.PostForm("Content");

  if (titleValue == "") || (contentValue == "") {
    c.JSON(406, gin.H{"message": "Invalid form", "Title": titleValue, "Content": contentValue})
    c.Abort()
    return
  }

  diaryID, err := diaryModel.Create(userID, diaryForm)

  if diaryID > 0 && err != nil {
    c.JSON(406, gin.H{"message": "diary could not be created", "error": err.Error()})
    c.Abort()
    return
  }

  c.JSON(200, gin.H{"message": "diary created", "id": diaryID})
}

//All ...
func (ctrl DiaryController) All(c *gin.Context) {
  userID := getUserID(c)

  if userID == 0 {
    c.JSON(403, gin.H{"message": "Please login first"})
    c.Abort()
    return
  }

  data, err := diaryModel.All(userID)

  if err != nil {
    c.JSON(406, gin.H{"Message": "Could not get the diaries", "error": err.Error()})
    c.Abort()
    return
  }

  c.HTML(http.StatusOK, "diary.html", gin.H{"data": data})
}

//One ...
func (ctrl DiaryController) One(c *gin.Context) {
  userID := getUserID(c)

  if userID == 0 {
    c.JSON(403, gin.H{"message": "Please login first"})
    c.Abort()
    return
  }

  id := c.Param("id")

  if id, err := strconv.ParseInt(id, 10, 64); err == nil {

    data, err := diaryModel.One(userID, id)
    if err != nil {
      c.JSON(404, gin.H{"Message": "diary not found", "error": err.Error()})
      c.Abort()
      return
    }
    c.JSON(200, gin.H{"data": data})
  } else {
    c.JSON(404, gin.H{"Message": "Invalid parameter"})
  }
}

//Update ...
func (ctrl DiaryController) Update(c *gin.Context) {
  userID := getUserID(c)

  if userID == 0 {
    c.JSON(403, gin.H{"message": "Please login first"})
    c.Abort()
    return
  }

  id := c.Param("id")
  if id, err := strconv.ParseInt(id, 10, 64); err == nil {

    var diaryForm forms.DiaryForm

    if c.BindJSON(&diaryForm) != nil {
      c.JSON(406, gin.H{"message": "Invalid parameters", "form": diaryForm})
      c.Abort()
      return
    }

    err := diaryModel.Update(userID, id, diaryForm)
    if err != nil {
      c.JSON(406, gin.H{"Message": "diary could not be updated", "error": err.Error()})
      c.Abort()
      return
    }
    c.JSON(200, gin.H{"message": "diary updated"})
  } else {
    c.JSON(404, gin.H{"Message": "Invalid parameter", "error": err.Error()})
  }
}

//Delete ...
func (ctrl DiaryController) Delete(c *gin.Context) {
  userID := getUserID(c)

  if userID == 0 {
    c.JSON(403, gin.H{"message": "Please login first"})
    c.Abort()
    return
  }

  id := c.Param("id")
  if id, err := strconv.ParseInt(id, 10, 64); err == nil {

    err := diaryModel.Delete(userID, id)
    if err != nil {
      c.JSON(406, gin.H{"Message": "diary could not be deleted", "error": err.Error()})
      c.Abort()
      return
    }
    c.JSON(200, gin.H{"message": "diary deleted"})
  } else {
    c.JSON(404, gin.H{"Message": "Invalid parameter"})
  }
}
