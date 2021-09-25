package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ensureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedIn, _ := c.Get("is_logged_in")
		if loggedIn == false || loggedIn == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func ensureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		loggedIn, _ := c.Get("is_logged_in")
		if loggedIn == true {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func setUserStatus() gin.HandlerFunc {
	return func(c *gin.Context) {
		if token, err := c.Cookie("token"); err == nil || token != "" {
			c.Set("is_logged_in", true)
		} else {
			c.Set("is_logged_in", false)
		}
	}
}
