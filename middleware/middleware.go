package middleware

import (
	auth "app/usecase/auth"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type H struct {
	Message string `json:"message"`
}

func RecordLogAndTime(c *gin.Context) {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err.Error())
	}
	oldTime := time.Now()
	c.Next()
	logger.Info("incoming request",
		zap.String("path", c.Request.URL.Path),
		zap.Int("status", c.Writer.Status()),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
}

func LoginCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Access-Token")

		_, err := auth.ValidToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, &H{Message: "please login"})
			c.Abort()
		} else {
			c.Next()
		}
	}
}
