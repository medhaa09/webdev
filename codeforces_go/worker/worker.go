package worker

import (
	"codeforces/data"
	"codeforces/models"
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

		maxTimeStamp, err := m.GetMaxTimeStamp()
		if err != nil {
			fmt.Printf("Error fetching max timestamp: %v. Defaulting to 0.\n", err)
			maxTimeStamp = 0 // Assuming 0 is a safe default if no actions are stored yet
		}

		// Creates a new CodeforcesClient
		cfClient := &data.CodeforcesClient{
			Client: http.DefaultClient,
			Mongo:  m}
		wg.Add(1)
		go func() {
			defer wg.Done()
			recentActionsData, err := cfClient.RecentActions(50)
			if err != nil {
				fmt.Println("Failed to retrieve recent actions data")
				return
			}

			// Filter actions newer than the last stored timestamp
			var newActions []models.RecentActions
			for _, action := range recentActionsData {
				if action.Time > maxTimeStamp {
					newActions = append(newActions, action)
				}
			}

			if len(newActions) > 0 {
				err = m.StoreRecentActionsInTheDatabase(newActions)
				if err != nil {
					fmt.Println("Error storing recent actions data:", err)
				} else {
					fmt.Println("New recent actions data with new time stored successfully.")
				}
			} else {
				fmt.Println("No new actions to store.")
			}
		}()
		wg.Wait()
		// Sleep for 5 minutes
		time.Sleep(5 * time.Minute)
	}
}
