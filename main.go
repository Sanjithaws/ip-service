package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func getIP(c *gin.Context) {
	ip := c.ClientIP()
	if realIP := c.GetHeader("X-Real-IP"); realIP != "" {
		ip = realIP
	}
	if fwd := c.GetHeader("X-Forwarded-For"); fwd != "" {
		ip = strings.Split(strings.TrimSpace(fwd), ",")[0]
	}
	c.JSON(http.StatusOK, gin.H{"ip": strings.TrimSpace(ip)})
}

func health(c *gin.Context) { c.String(http.StatusOK, "OK") }

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/ipconfig", getIP)
	r.GET("/health", health)
	r.GET("/ready", health)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
