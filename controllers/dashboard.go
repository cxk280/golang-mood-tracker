package controllers

import (
  // "strconv"
  "net/http"
  "fmt"

  // "golang-mood-tracker/forms"
  "golang-mood-tracker/models"

  "github.com/gin-gonic/gin"
)

//DashboardController ...
type DashboardController struct{}

var dashboardModel = new(models.DashboardModel)

// //Create ...
// func (ctrl DashboardController) Create(c *gin.Context) {
//   userID := getUserID(c)

//   if userID == 0 {
//     c.JSON(403, gin.H{"message": "Please login first"})
//     c.Abort()
//     return
//   }

//   var dashboardForm forms.DashboardForm

//   if c.BindJSON(&dashboardForm) != nil {
//     c.JSON(406, gin.H{"message": "Invalid form", "form": dashboardForm})
//     c.Abort()
//     return
//   }

//   dashboardID, err := dashboardModel.Create(userID, dashboardForm)

//   if dashboardID > 0 && err != nil {
//     c.JSON(406, gin.H{"message": "dashboard could not be created", "error": err.Error()})
//     c.Abort()
//     return
//   }

//   c.JSON(200, gin.H{"message": "dashboard created", "id": dashboardID})
// }

//All ...
func (ctrl DashboardController) All(c *gin.Context) {
  userID := getUserID(c)

  if userID == 0 {
    c.HTML(http.StatusOK, "error.html", gin.H{
      "errorMessage": "Please login first.",
    })
    c.Abort()
    return
  }

  // data, err := dashboardModel.All(userID)

  // if err != nil {
  //   c.JSON(406, gin.H{"Message": "Could not get the dashboards", "error": err.Error()})
  //   c.Abort()
  //   return
  // }

  // c.JSON(200, gin.H{"data": data})

  fmt.Println("opening dashboard.html")

  c.HTML(http.StatusOK, "dashboard.html", gin.H{
    })
}

// //One ...
// func (ctrl DashboardController) One(c *gin.Context) {
//   userID := getUserID(c)

//   if userID == 0 {
//     c.JSON(403, gin.H{"message": "Please login first"})
//     c.Abort()
//     return
//   }

//   id := c.Param("id")

//   if id, err := strconv.ParseInt(id, 10, 64); err == nil {

//     data, err := dashboardModel.One(userID, id)
//     if err != nil {
//       c.JSON(404, gin.H{"Message": "dashboard not found", "error": err.Error()})
//       c.Abort()
//       return
//     }
//     c.JSON(200, gin.H{"data": data})
//   } else {
//     c.JSON(404, gin.H{"Message": "Invalid parameter"})
//   }
// }

// //Update ...
// func (ctrl DashboardController) Update(c *gin.Context) {
//   userID := getUserID(c)

//   if userID == 0 {
//     c.JSON(403, gin.H{"message": "Please login first"})
//     c.Abort()
//     return
//   }

//   id := c.Param("id")
//   if id, err := strconv.ParseInt(id, 10, 64); err == nil {

//     var dashboardForm forms.DashboardForm

//     if c.BindJSON(&dashboardForm) != nil {
//       c.JSON(406, gin.H{"message": "Invalid parameters", "form": dashboardForm})
//       c.Abort()
//       return
//     }

//     err := dashboardModel.Update(userID, id, dashboardForm)
//     if err != nil {
//       c.JSON(406, gin.H{"Message": "dashboard could not be updated", "error": err.Error()})
//       c.Abort()
//       return
//     }
//     c.JSON(200, gin.H{"message": "dashboard updated"})
//   } else {
//     c.JSON(404, gin.H{"Message": "Invalid parameter", "error": err.Error()})
//   }
// }

// //Delete ...
// func (ctrl DashboardController) Delete(c *gin.Context) {
//   userID := getUserID(c)

//   if userID == 0 {
//     c.JSON(403, gin.H{"message": "Please login first"})
//     c.Abort()
//     return
//   }

//   id := c.Param("id")
//   if id, err := strconv.ParseInt(id, 10, 64); err == nil {

//     err := dashboardModel.Delete(userID, id)
//     if err != nil {
//       c.JSON(406, gin.H{"Message": "dashboard could not be deleted", "error": err.Error()})
//       c.Abort()
//       return
//     }
//     c.JSON(200, gin.H{"message": "dashboard deleted"})
//   } else {
//     c.JSON(404, gin.H{"Message": "Invalid parameter"})
//   }
// }
