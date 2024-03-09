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
		fmt.Println("Error occurred while calling the API: ", err)
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		fmt.Printf("Received non-OK HTTP status code: %d\n", response.StatusCode)
		data, _ := io.ReadAll(response.Body)
		fmt.Println("Response Body:", string(data))
		return nil, fmt.Errorf("received non-OK HTTP status code: %d", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error occurred while reading response body: ", err)
		return nil, err
	}

	var result struct {
		Status        string                 `json:"status"`
		RecentActions []models.RecentActions `json:"result"`
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("Error occurred while decoding JSON: ", err)
		fmt.Println("Response body: ", string(data))
		return nil, err
	}

	return result.RecentActions, nil
}
