package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/Massad/gin-boilerplate/db"
	"golang-mood-tracker/controllers"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

//CORSMiddleware ...
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func main() {
	r := gin.Default()

	store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("gin-boilerplate-session", store))

	r.Use(CORSMiddleware())

	db.Init()

	v1 := r.Group("/v1")
	{


		// USE THESE CONTROLLERS FOR USER AUTH

		/*** START USER ***/
		user := new(controllers.UserController)

		v1.POST("/user/signin", user.Signin)
		v1.POST("/user/signup", user.Signup)
		v1.GET("/user/signout", user.Signout)

		/*** START Article ***/
		article := new(controllers.ArticleController)

		v1.POST("/article", article.Create)
		v1.GET("/articles", article.All)
		v1.GET("/article/:id", article.One)
		v1.PUT("/article/:id", article.Update)
		v1.DELETE("/article/:id", article.Delete)
	}

	user := new(controllers.UserController)

	r.LoadHTMLGlob("./public/html/*.html")

	r.Static("/public", "./public")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ginBoilerplateVersion": "v0.03",
			"goVersion":             runtime.Version(),
		})
	})

	r.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", gin.H{
		})
	})

	r.GET("/analytics", func(c *gin.Context) {
		c.HTML(http.StatusOK, "analytics.html", gin.H{
		})
	})

	r.GET("/dashboard", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard.html", gin.H{
		})
	})

	r.GET("/diary", func(c *gin.Context) {
		c.HTML(http.StatusOK, "diary.html", gin.H{
		})
	})

	r.GET("/feed", func(c *gin.Context) {
		c.HTML(http.StatusOK, "feed.html", gin.H{
		})
	})

	// r.POST("/login", func(c *gin.Context) {
		// emailValue := c.PostForm("email");
		// passwordValue := c.PostForm("password");

		// c.JSON(200, gin.H{
		// 	"status":  "posted to login",
		// 	"message": "whoo",
		// 	"email": emailValue,
		// 	"password": passwordValue})
	// })

	r.POST("/login", user.Signin)

	r.POST("/signup", func(c *gin.Context) {
		emailValue := c.PostForm("email");
		passwordValue := c.PostForm("password");

		c.JSON(200, gin.H{
			"status":  "posted to signup",
			"message": "whoo",
			"email": emailValue,
			"password": passwordValue})
	})

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	r.Run(":9000")
}
