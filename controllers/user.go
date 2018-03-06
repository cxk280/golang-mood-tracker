package controllers

import (

  "fmt"
  "net/http"

	"golang-mood-tracker/forms"
	"golang-mood-tracker/models"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
  "github.com/gin-gonic/gin/binding"
)

//UserController ...
type UserController struct{}

var userModel = new(models.UserModel)

//getUserID ...
func getUserID(c *gin.Context) int64 {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID != nil {
		return models.ConvertToInt64(userID)
	}
	return 0
}

//getSessionUserInfo ...
func getSessionUserInfo(c *gin.Context) (userSessionInfo models.UserSessionInfo) {
	session := sessions.Default(c)
	userID := session.Get("user_id")
	if userID != nil {
		userSessionInfo.ID = models.ConvertToInt64(userID)
		userSessionInfo.Name = session.Get("user_name").(string)
		userSessionInfo.Email = session.Get("user_email").(string)
	}
	return userSessionInfo
}

//Signin ...
func (ctrl UserController) Signin(c *gin.Context) {

	var signinForm forms.SigninForm

	//Line below is blank
	fmt.Println(signinForm)

  if err := c.ShouldBindWith(&signinForm, binding.Form); err != nil {
	  c.HTML(http.StatusOK, "error.html", gin.H{
	    "errorMessage": "Invalid signin form.",
	  })
	  c.Abort()
	  return
  }

  //Line below contains correct data
  fmt.Println(signinForm)

  user, err := userModel.Signin(signinForm)



	if err == nil {

		//redis-server must be running for sessions to work
		session := sessions.Default(c)
		session.Set("user_id", user.ID)
		session.Set("user_email", user.Email)
		session.Set("user_name", user.Name)
		session.Save()

		c.Redirect(http.StatusMovedPermanently, "/dashboard")
	} else {
		c.HTML(http.StatusOK, "error.html", gin.H{
	    "errorMessage": "Invalid signin details.",
	  })
	}

}

//Signup ...
func (ctrl UserController) Signup(c *gin.Context) {

	var signupForm forms.SignupForm

	if err := c.ShouldBindWith(&signupForm, binding.Form); err != nil {
	  c.HTML(http.StatusOK, "error.html", gin.H{
	    "errorMessage": "Invalid signup form.",
	  })
	  c.Abort()
	  return
  }

	user, err := userModel.Signup(signupForm)

	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
	    "errorMessage": "Invalid form.",
	  })
		c.Abort()
		return
	}

	if user.ID > 0 {

		c.Redirect(http.StatusMovedPermanently, "/")

	} else {
		c.HTML(http.StatusOK, "error.html", gin.H{
    "errorMessage": "Could not sign up this user.",
  })
	}

}

//Signout ...
func (ctrl UserController) Signout(c *gin.Context) {

	userID := getUserID(c)

	fmt.Println(userID)

  if userID == 0 {
    c.HTML(http.StatusOK, "error.html", gin.H{
      "errorMessage": "No user is logged in currently, so you can't sign out...",
    })
    c.Abort()
    return
  }


	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.HTML(http.StatusOK, "error.html", gin.H{
    "errorMessage": "Signed out.",
  })
}
