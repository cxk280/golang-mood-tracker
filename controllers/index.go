package controllers

import (
  // "strconv"
  "net/http"

  // "golang-mood-tracker/forms"
  "golang-mood-tracker/models"

  "github.com/gin-gonic/gin"
)

//IndexController ...
type IndexController struct{}

var indexModel = new(models.IndexModel)

// //Create ...
// func (ctrl IndexController) Create(c *gin.Context) {
//   userID := getUserID(c)

//   if userID == 0 {
//     c.JSON(403, gin.H{"message": "Please login first"})
//     c.Abort()
//     return
//   }

//   var indexForm forms.IndexForm

//   if c.BindJSON(&indexForm) != nil {
//     c.JSON(406, gin.H{"message": "Invalid form", "form": indexForm})
//     c.Abort()
//     return
//   }

//   indexID, err := indexModel.Create(userID, indexForm)

//   if indexID > 0 && err != nil {
//     c.JSON(406, gin.H{"message": "index could not be created", "error": err.Error()})
//     c.Abort()
//     return
//   }

//   c.JSON(200, gin.H{"message": "index created", "id": indexID})
// }

//All ...
func (ctrl IndexController) All(c *gin.Context) {
  // userID := getUserID(c)

  // if userID == 0 {
  //   c.JSON(403, gin.H{"message": "Please login first"})
  //   c.Abort()
  //   return
  // }

  // data, err := indexModel.All(userID)

  // if err != nil {
  //   c.JSON(406, gin.H{"Message": "Could not get the indexs", "error": err.Error()})
  //   c.Abort()
  //   return
  // }

  // c.JSON(200, gin.H{"data": data})

  c.HTML(http.StatusOK, "index.html", gin.H{
      })
}

// //One ...
// func (ctrl IndexController) One(c *gin.Context) {
//   userID := getUserID(c)

//   if userID == 0 {
//     c.JSON(403, gin.H{"message": "Please login first"})
//     c.Abort()
//     return
//   }

//   id := c.Param("id")

//   if id, err := strconv.ParseInt(id, 10, 64); err == nil {

//     data, err := indexModel.One(userID, id)
//     if err != nil {
//       c.JSON(404, gin.H{"Message": "index not found", "error": err.Error()})
//       c.Abort()
//       return
//     }
//     c.JSON(200, gin.H{"data": data})
//   } else {
//     c.JSON(404, gin.H{"Message": "Invalid parameter"})
//   }
// }

// //Update ...
// func (ctrl IndexController) Update(c *gin.Context) {
//   userID := getUserID(c)

//   if userID == 0 {
//     c.JSON(403, gin.H{"message": "Please login first"})
//     c.Abort()
//     return
//   }

//   id := c.Param("id")
//   if id, err := strconv.ParseInt(id, 10, 64); err == nil {

//     var indexForm forms.IndexForm

//     if c.BindJSON(&indexForm) != nil {
//       c.JSON(406, gin.H{"message": "Invalid parameters", "form": indexForm})
//       c.Abort()
//       return
//     }

//     err := indexModel.Update(userID, id, indexForm)
//     if err != nil {
//       c.JSON(406, gin.H{"Message": "index could not be updated", "error": err.Error()})
//       c.Abort()
//       return
//     }
//     c.JSON(200, gin.H{"message": "index updated"})
//   } else {
//     c.JSON(404, gin.H{"Message": "Invalid parameter", "error": err.Error()})
//   }
// }

// //Delete ...
// func (ctrl IndexController) Delete(c *gin.Context) {
//   userID := getUserID(c)

//   if userID == 0 {
//     c.JSON(403, gin.H{"message": "Please login first"})
//     c.Abort()
//     return
//   }

//   id := c.Param("id")
//   if id, err := strconv.ParseInt(id, 10, 64); err == nil {

//     err := indexModel.Delete(userID, id)
//     if err != nil {
//       c.JSON(406, gin.H{"Message": "index could not be deleted", "error": err.Error()})
//       c.Abort()
//       return
//     }
//     c.JSON(200, gin.H{"message": "index deleted"})
//   } else {
//     c.JSON(404, gin.H{"Message": "Invalid parameter"})
//   }
// }
