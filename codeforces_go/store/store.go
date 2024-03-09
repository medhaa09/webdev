package store

import (
	"codeforces/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	Collection1 *mongo.Collection
	Collection2 *mongo.Collection
}

const uri = "mongodb+srv://medha:drumDRO67%23%24@cluster0.qj0tdiv.mongodb.net/"

func (m *MongoStore) OpenConnectionWithMongoDB() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("codeforces").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}
	m.Collection1 = client.Database("codeforces").Collection("cfdata")
	m.Collection2 = client.Database("codeforces").Collection("users")
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

func (m *MongoStore) StoreRecentActionsInTheDatabase(actions []models.RecentActions) error {
	var toInsertInterface []interface{}
	for _, action := range actions {
		toInsertInterface = append(toInsertInterface, action)
	}
	fmt.Println("trying to insert document to mongodb")
	_, err1 := m.Collection1.InsertMany(context.TODO(), toInsertInterface)
	if err1 != nil {
		fmt.Println("error inserting documents: \n", err1)
		return err1
	}
	fmt.Println("Insertion successful")
	return nil
}
func (m *MongoStore) StoreUserData(user models.User) error {
	fmt.Println("Trying to insert user data into MongoDB")
	_, err := m.Collection2.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println("Error inserting user data:", err)
		return err
	}
	fmt.Println("Insertion of user data successful")
	return nil
}

// Fetch all unique blog IDs from the data
func FetchUniqueBlogIDs(recentActions []models.RecentActions) []int {
	uniqueIDs := make(map[int]bool)
	for _, action := range recentActions {
		uniqueIDs[action.Blog.Id] = true
	}

	ids := make([]int, 0, len(uniqueIDs))
	for id := range uniqueIDs {
		ids = append(ids, id)
	}
	return ids
}
func (m *MongoStore) QueryRecentActions(blogID int) ([]models.RecentActions, error) {
	filter := bson.M{"blog.id": blogID}
	cursor, err := m.Collection1.Find(context.TODO(), filter)
	if err != nil {
		fmt.Println("Error occured while querying recentActions collection: ", err)
	}
	var results []models.RecentActions
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		fmt.Println("Error occured while iterating the cursor: ", err)
	}
	return results, err
}
func (m *MongoStore) UserLogin(username string, password string) bool {

	var foundUser models.User
	err := m.Collection2.FindOne(context.TODO(), bson.M{
		"username": username,
		"password": password,
	}).Decode(&foundUser)

	if err != nil {
		fmt.Println("wrong credentials: ", err)
		return false
	}
	return true
}
