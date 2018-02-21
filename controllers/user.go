package controllers

import (

  "fmt"
  "log"
  "regexp"
  "unicode"

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

func verifyPassword(s string) (sevenOrMore, number, upper, special bool) {
    letters := 0
    for _, s := range s {
        switch {
        case unicode.IsNumber(s):
            number = true
        case unicode.IsUpper(s):
            upper = true
            letters++
        case unicode.IsPunct(s) || unicode.IsSymbol(s):
            special = true
        case unicode.IsLetter(s) || s == ' ':
            letters++
        default:
        		fmt.Println("password field is not an password")
            return false, false, false, false
        }
    }
    sevenOrMore = letters >= 7
    if sevenOrMore == true {
    	fmt.Println("password field is an accepted password")
    } else {
    	fmt.Println("password field has too few letters")
    }
    return
	}

//Signin ...
func (ctrl UserController) Signin(c *gin.Context) {

	var signinForm forms.SigninForm
  user, err := userModel.Signin(signinForm)

  // The if conditional below validates if the email is an email or not
  if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, c.PostForm("email")); !m {
    fmt.Println("email field is not an email")
  }	else {
    fmt.Println("email field is an email")
  }

  verifyPassword(c.PostForm("password"))

  if err := c.ShouldBindWith(&signinForm, binding.Form); err != nil {
    fmt.Println("***************************")
    log.Println("err: ",err)
    fmt.Println("***************************")
	  c.JSON(406, gin.H{"message": "Invalid signin form", "form": signinForm})
	  c.Abort()
	  return
  }


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

	if c.BindJSON(&signupForm) != nil {
		c.JSON(406, gin.H{"message": "Invalid form boo", "form": signupForm})
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
		session := sessions.Default(c)
		session.Set("user_id", user.ID)
		session.Set("user_email", user.Email)
		session.Set("user_name", user.Name)
		session.Save()
		c.JSON(200, gin.H{"message": "Success signup", "user": user})
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
