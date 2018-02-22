package controllers

import (

  "fmt"
  "log"
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

  if err := c.ShouldBindWith(&signinForm, binding.Form); err != nil {
    fmt.Println("***************************")
    fmt.Println("This is inside if")
    log.Println("err: ",err)
    log.Println("signinForm: ",signinForm)
    fmt.Println("***************************")
	  c.JSON(406, gin.H{"message": "Invalid signin form", "form": signinForm})
	  c.Abort()
	  return
  }

  user, err := userModel.Signin(signinForm)

  fmt.Println("***************************")
  fmt.Println("This is outside if")
  log.Println("err: ",err)
  log.Println("signinForm: ",signinForm)
  fmt.Println("***************************")

	if err == nil {
		session := sessions.Default(c)
		session.Set("user_id", user.ID)
		session.Set("user_email", user.Email)
		session.Set("user_name", user.Name)
		session.Save()

		c.JSON(200, gin.H{"message": "User signed in", "user": user})
	} else {
		c.JSON(406, gin.H{"message": "Invalid signin details", "error": err.Error()})
	}

		// emailValue := c.PostForm("email");
		// passwordValue := c.PostForm("password");

		// c.JSON(200, gin.H{
		// 	"status":  "posted to login",
		// 	"message": "whoo",
		// 	"email": emailValue,
		// 	"password": passwordValue})

}

//Signup ...
func (ctrl UserController) Signup(c *gin.Context) {

	var signupForm forms.SignupForm

	if err := c.ShouldBindWith(&signupForm, binding.Form); err != nil {
    fmt.Println("***************************")
    fmt.Println("This is inside if")
    log.Println("err: ",err)
    log.Println("signinForm: ",signupForm)
    fmt.Println("***************************")
	  c.JSON(406, gin.H{"message": "Invalid signup form", "form": signupForm})
	  c.Abort()
	  return
  }

	user, err := userModel.Signup(signupForm)

	if err != nil {
		c.JSON(406, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	if user.ID > 0 {

		//This session section is causing the runtime error with "invalid memory address or nil pointer dereference"
		//Maybe instead of using session I should just to login page and have them log in?
		// session := sessions.Default(c)
		// session.Set("user_id", user.ID)
		// session.Set("user_email", user.Email)
		// session.Set("user_name", user.Name)
		// session.Save()

		// c.JSON(200, gin.H{"message": "Success signup", "user": user})

		c.Redirect(http.StatusMovedPermanently, "/")

	} else {
		c.JSON(406, gin.H{"message": "Could not signup this user", "error": err.Error()})
	}

}

//Signout ...
func (ctrl UserController) Signout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(200, gin.H{"message": "Signed out..."})
}
