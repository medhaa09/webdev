package data

import (
	"codeforces/models"
	"codeforces/store"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type CodeforcesClient struct {
	Client *http.Client
	Mongo  *store.MongoStore
}

func (cfClient *CodeforcesClient) RecentActions(maxCount int) ([]models.RecentActions, error) {
	apiURL := "https://codeforces.com/api/recentActions?maxCount=" + strconv.Itoa(maxCount)
	response, err := cfClient.Client.Get(apiURL)
	if err != nil {
		fmt.Println("error occured while calling of api: ", err)
		return nil, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error occured while reading response body: ", err)
		return nil, err

	}
	// Decode the []byte into a slice of models.RecentActions
	var result struct {
		Status        string                 `json:"status"`
		RecentActions []models.RecentActions `json:"result"`
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("Error occurred while decoding JSON: ", err)
		return nil, err
	}

	return result.RecentActions, nil
}

//https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Collection.InsertMany this function takes interfaces ka array as input. create an empty array and keep appending the recent actions
//in mongo struct, there is one collection of recentactions. if we keep users info then we'll add another collection for users to it.
