package main

import (
	"codeforces/Auth"
	"codeforces/data"
	"codeforces/models"
	"codeforces/store"
	"codeforces/worker"
	"fmt"
	"net/http"
	"time"

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

	// Configure CORS middleware
	config := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(config))

	protected := router.Group("/")
	protected.Use(cors.New(config))
	protected.Use(Auth.TokenAuthMiddleware())

	// Setup CORS middleware
	// This applies the default CORS policies
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

	protected.GET("/activity/recent-actions-grouped", func(c *gin.Context) {
		cfClient := data.CodeforcesClient{Client: http.DefaultClient, Mongo: mongoStore}
		recentActionsData, err := cfClient.RecentActions(50)
		if err != nil {
			fmt.Printf("Error fetching recent actions: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve recent actions data grouped"})
			return
		}
		uniqueBlogIDs := store.FetchUniqueBlogIDs(recentActionsData)
		groupedCommentsMap := make(map[int]models.GroupedComments) // Corrected: Map to single GroupedComments per blog ID

		for _, blogID := range uniqueBlogIDs {
			actions, err := mongoStore.QueryRecentActions(blogID)
			//fmt.Printf("%#v\n", actions)
			if err != nil {
				fmt.Printf("Error querying recent actions for blog ID %d: %v\n", blogID, err)
				continue
			}
			// Initialize the structure only if it doesn't exist
			if _, exists := groupedCommentsMap[blogID]; !exists {
				groupedCommentsMap[blogID] = models.GroupedComments{
					BlogID:    blogID,
					BlogTitle: "", // Initialize empty; will be set in the loop below
					Comments:  []string{}}
			}
			for _, action := range actions {
				Grouped := groupedCommentsMap[blogID]                                // Copy existing structure out of the map
				Grouped.BlogTitle = action.Blog.Title                                // Set title (redundant if done multiple times but ensures it's set)
				Grouped.Comments = append(Grouped.Comments, action.Comments.Comment) // Append comments
				groupedCommentsMap[blogID] = Grouped                                 // Put the modified structure back into the map
			}
		}
		// Convert map to slice for JSON serialization because in maps the order of iteration is not fixed so we convert it into a slice adn then send it as JSON
		groupedComments := make([]models.GroupedComments, 0, len(groupedCommentsMap)) // making a slice groupedComments
		for _, group := range groupedCommentsMap {
			groupedComments = append(groupedComments, group)
		}
		c.JSON(http.StatusOK, gin.H{"groupedComments": groupedComments})
	})
	router.POST("/user/login", func(c *gin.Context) {
		var loginCredentials models.User
		err := c.ShouldBindJSON(&loginCredentials)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		isAuthenticated := mongoStore.UserLogin(loginCredentials.Username, loginCredentials.Password)
		if isAuthenticated {
			c.JSON(http.StatusOK, gin.H{"message": "successful login"})
			signedToken, signedRefreshToken, err := Auth.GenerateAllTokens(loginCredentials.Username)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"token": signedToken, "refreshToken": signedRefreshToken})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Credentials"})
		}
	})
	// Listen and serve on the specified port
	port := ":8080"
	fmt.Printf("Server running on port %s\n", port)
	if err := router.Run(port); err != nil {
		fmt.Printf("Failed to start server: %s\n", err)
	}
}
