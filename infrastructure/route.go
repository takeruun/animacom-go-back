package infrastructure

import (
	controllers "app/interfaces/controllers"
	middleware "app/middleware"

	"github.com/gin-gonic/gin"
)

type Routing struct {
	DB   *DB
	Gin  *gin.Engine
	Port string
}

func NewRouting(db *DB) *Routing {
	c := NewConfig()

	r := &Routing{
		DB:   db,
		Gin:  gin.Default(),
		Port: c.Routing.Port,
	}
	r.setRouting()
	return r
}

func (r *Routing) setRouting() {
	usersController := controllers.NewUsersController(r.DB)
	postsController := controllers.NewPostsController(r.DB)
	v1 := r.Gin.Group("v1")
	v1.Use(middleware.RecordLogAndTime)
	{
		users := v1.Group("users")
		{
			posts := users.Group("posts")
			posts.Use(middleware.LoginCheckMiddleware())
			{
				posts.GET("", func(c *gin.Context) { postsController.Get(c) })
			}

			users.GET("/:id", func(c *gin.Context) { usersController.Get(c) })
			users.POST("/auth/registrations", func(c *gin.Context) { usersController.Create(c) })
			users.POST("/auth/session", func(c *gin.Context) { usersController.Login(c) })
		}
	}
}

func (r *Routing) SetMiddleware() {
	r.Gin.Use(gin.Recovery(), gin.Logger())
}

func (r *Routing) Run() {
	r.Gin.Run(r.Port)
}
