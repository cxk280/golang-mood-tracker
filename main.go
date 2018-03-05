package main

import (
	"fmt"
	"net/http"
	// "runtime"

	"golang-mood-tracker/db"
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

	user 				:= new(controllers.UserController)
	analytics 	:= new(controllers.AnalyticsController)
	dashboard 	:= new(controllers.DashboardController)
	diary 			:= new(controllers.DiaryController)
	feed 				:= new(controllers.FeedController)
	index 			:= new(controllers.IndexController)

	r.LoadHTMLGlob("./public/html/*.html")

	r.Static("/public", "./public")

	// r.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "index.html", gin.H{
	// 		"ginBoilerplateVersion": "v0.03",
	// 		"goVersion":             runtime.Version(),
	// 	})
	// })

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

	r.POST("/feed", feed.Create)
	r.GET("/feed", feed.All)
	r.GET("/feed/:id", feed.One)
	r.PUT("/feed/:id", feed.Update)
	r.DELETE("/feed/:id", feed.Delete)

	r.POST("/login", user.Signin)

	r.POST("/signup", user.Signup)

	r.POST("/signout", user.Signout)

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{})
	})

	r.Run(":9000")
}
