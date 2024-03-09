package worker

import (
	"codeforces/data"
	"codeforces/store"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func PerformWork(m *store.MongoStore) {
	var wg sync.WaitGroup
	for {
		// Opens connection with the database
		m.OpenConnectionWithMongoDB()

		// Creates a new CodeforcesClient
		cfClient := &data.CodeforcesClient{
			Client: http.DefaultClient,
			Mongo:  m}
		wg.Add(1)
		go func() {
			defer wg.Done()
			recentActionsData, err := cfClient.RecentActions(50)
			if err != nil {
				//c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve recent actions data"})
				fmt.Println("Failed to retrieve recent actions data")
				return
			}

			m.StoreRecentActionsInTheDatabase(recentActionsData)
		}()
		wg.Wait()
		// Sleep for 5 minutes
		time.Sleep(5 * time.Minute)
	}
}
