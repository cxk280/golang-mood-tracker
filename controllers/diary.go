package controllers

import (
  "strconv"
  "net/http"
  "fmt"

  "golang-mood-tracker/forms"
  "golang-mood-tracker/models"

  "github.com/gin-gonic/gin"
  "github.com/gin-gonic/gin/binding"
)

//DiaryController ...
type DiaryController struct{}

var diaryModel = new(models.DiaryModel)

// //Create ...
func (ctrl DiaryController) Create(c *gin.Context) {

  feelingValue := c.PostForm("feeling");
  fmt.Println("feelingValue")
  fmt.Println(feelingValue)

  //Sign in via the browser with redis-server running or in Postman directly in order for Create to increment properly
  userID := getUserID(c)

  if userID == 0 {
    c.HTML(http.StatusOK, "error.html", gin.H{
      "errorMessage": "Please login first.",
    })
    c.Abort()
    return
  }

  var diaryForm forms.DiaryForm

  //Make sure to use lower-case keys in the form data
  if err := c.ShouldBindWith(&diaryForm, binding.Form); err != nil {
    c.HTML(http.StatusOK, "error.html", gin.H{
      "errorMessage": "Invalid form.",
    })
    c.Abort()
    return
  }

  diaryID, err := diaryModel.Create(userID, diaryForm)

  if diaryID > 0 && err != nil {
    c.HTML(http.StatusOK, "error.html", gin.H{
      "errorMessage": "Diary could not be created.",
    })
    c.Abort()
    return
  }

  c.HTML(http.StatusOK, "error.html", gin.H{
    "errorMessage": "Diary created.",
  })
}

//All ...
func (ctrl DiaryController) All(c *gin.Context) {
  userID := getUserID(c)

  if userID == 0 {
    c.HTML(http.StatusOK, "error.html", gin.H{
      "errorMessage": "Please login first.",
    })
    c.Abort()
    return
  }

  data, err := diaryModel.All(userID)

  fmt.Println("data in diary.All")
  fmt.Println(data)

  if err != nil {
    c.HTML(http.StatusOK, "error.html", gin.H{
      "errorMessage": "Could not get the diaries.",
    })
    c.Abort()
    return
  }

  c.JSON(200, gin.H{"data": data})
}

//One ...
func (ctrl DiaryController) One(c *gin.Context) {
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

    data, err := diaryModel.One(userID, id)
    if err != nil {
      c.HTML(http.StatusOK, "error.html", gin.H{
        "errorMessage": "Diary not found.",
      })
      c.Abort()
      return
    }
    c.JSON(200, gin.H{"data": data})
  } else {
    c.HTML(http.StatusOK, "error.html", gin.H{
      "errorMessage": "Invalid parameter.",
    })
  }
}

//Update ...
func (ctrl DiaryController) Update(c *gin.Context) {
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

    var diaryForm forms.DiaryForm

    if c.BindJSON(&diaryForm) != nil {
      c.HTML(http.StatusOK, "error.html", gin.H{
        "errorMessage": "Invalid parameters.",
      })
      c.Abort()
      return
    }

    err := diaryModel.Update(userID, id, diaryForm)
    if err != nil {
      c.HTML(http.StatusOK, "error.html", gin.H{
        "errorMessage": "Diary could not be updated.",
      })
      c.Abort()
      return
    }
    c.HTML(http.StatusOK, "error.html", gin.H{
      "errorMessage": "Diary updated.",
    })
  } else {
    c.HTML(http.StatusOK, "error.html", gin.H{
      "errorMessage": "Invalid parameter.",
    })
  }
}

//Delete ...
func (ctrl DiaryController) Delete(c *gin.Context) {
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

    err := diaryModel.Delete(userID, id)
    if err != nil {
      c.HTML(http.StatusOK, "error.html", gin.H{
        "errorMessage": "Diary could not be deleted.",
      })
      c.Abort()
      return
    }
    c.HTML(http.StatusOK, "error.html", gin.H{
      "errorMessage": "Diary deleted.",
    })
  } else {
    c.HTML(http.StatusOK, "error.html", gin.H{
      "errorMessage": "Invalid parameter.",
    })
  }
}
