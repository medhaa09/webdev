package main

import (
	"codeforces/data"
	"codeforces/models"
	"codeforces/store"
	"codeforces/worker"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors" // Use gin-contrib/cors, not rs/cors
	"github.com/gin-gonic/gin"
)

func main() {
	mongoStore := &store.MongoStore{}
	mongoStore.OpenConnectionWithMongoDB()

	// Launches PerformWork as a goroutine
	go worker.PerformWork(mongoStore)

	// Server (using gin framework)
	router := gin.Default()

	// Setup CORS middleware
	router.Use(cors.Default()) // This applies the default CORS policies
	router.POST("/user/signup", func(c *gin.Context) {
		var newUser models.User
		if err := c.ShouldBindJSON(&newUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := mongoStore.StoreUserData(newUser); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store user data"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Signup successful"})
	})
	// Route to show all the recent actions data
	router.GET("/activity/recent-actions", func(c *gin.Context) {
		cfClient := data.CodeforcesClient{Client: http.DefaultClient, Mongo: mongoStore}
		recentActionsData, err := cfClient.RecentActions(50)
		if err != nil {
			fmt.Printf("Error fetching recent actions: %v\n", err) // Log server-side for debugging
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve recent actions data"})

		}
		c.JSON(http.StatusOK, gin.H{"recentActions": recentActionsData})

	})

	// Listen and serve on the specified port
	port := ":8080"
	fmt.Printf("Server running on port %s\n", port)
	if err := router.Run(port); err != nil {
		fmt.Printf("Failed to start server: %s\n", err)
	}
}
