package main

import (
	"fmt"
	"net/http"

	"golang-mood-tracker/db"
	"golang-mood-tracker/controllers"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Allow CORS
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

	// Initialize Gin
	r := gin.Default()

	// Store session info
	store, _ := sessions.NewRedisStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	r.Use(sessions.Sessions("gin-boilerplate-session", store))

	// Call function to allow CORS
	r.Use(CORSMiddleware())

	// Initialize the database
	db.Init()

	// Initialize new controllers
	user 				:= new(controllers.UserController)
	analytics 			:= new(controllers.AnalyticsController)
	dashboard 			:= new(controllers.DashboardController)
	diary 				:= new(controllers.DiaryController)
	index 				:= new(controllers.IndexController)

	// Load HTML files
	r.LoadHTMLGlob("./public/html/*.html")

	// Load static assets
	r.Static("/public", "./public")



	// *****
	// Define routes
	r.GET("/", index.All)

	r.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup.html", gin.H{
		})
	})

	r.POST("/analytics", analytics.Create)
	r.GET("/analytics", analytics.All)
	r.GET("/analytics/:id", analytics.One)
	r.PUT("/analytics/:id", analytics.Update)
	r.DELETE("/analytics/:id", analytics.Delete)

	r.GET("/dashboard", dashboard.All)

	r.POST("/diary", diary.Create)
	r.GET("/diary", diary.All)
	r.GET("/diary/:id", diary.One)
	r.PUT("/diary/:id", diary.Update)
	r.DELETE("/diary/:id", diary.Delete)

	r.GET("/diaryPage", func(c *gin.Context) {
		c.HTML(http.StatusOK, "diary.html", gin.H{
		})
	})

	r.GET("/dashboard_answered", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard_answered.html", gin.H{
		})
	})

	r.GET("/dashboard_unanswered", func(c *gin.Context) {
		c.HTML(http.StatusOK, "dashboard_unanswered.html", gin.H{
		})
	})

	r.POST("/login", user.Signin)

	r.POST("/signup", user.Signup)

	r.GET("/signout", user.Signout)

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})
	// *****


	// Start the server
	r.Run(":9000")
}
