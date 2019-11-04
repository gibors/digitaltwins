package repository

// import (
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	"os"
// 	"context"
// )

// type MongoClient struct {
// 	Client                     *mongo.Client
// 	DMCollection *mongo.Collection
// }

// func getMongoClient() (*mongo.Client, error){
// 	DMSTRING = os.Getenv(DbConnectionString_DM)
// 	clientOptions := options.Client().ApplyURI(DMSTRING)
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err !=  nil {
// 		return (nil, error.Error(err))
// 	}
// 	return client
// }

// func (c *MongoClient) GetDeviceInserted(serialNumber string) map[string]interface{} {

// 	device : = c.DMCollection.find({"serialNumber": serialNumber})
// 	log.Println(device)
// 	return nil
// }
