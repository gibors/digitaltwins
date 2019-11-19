package repository

import (
	dev "caidc_auto_devicetwins/domain/model"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	ClientAut         *mongo.Client
	ClientLocal       *mongo.Client
	DMCollectionAut   *mongo.Collection
	DMCollectionLocal *mongo.Collection
}

func GetMongoClient(connectionString string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *MongoClient) GetDeviceInserted(serialNumber string) (dev.Device, error) {
	filter := bson.D{{"serialNumber", serialNumber}}
	var device = dev.Device{}
	err := c.DMCollectionAut.FindOne(context.TODO(), filter).Decode(&device)
	if err != nil {
		return device, err
	}
	return device, nil
}

func (c *MongoClient) InsertDevice(newDevice dev.Device) (bool, error) {

	insertResult, err := c.DMCollectionLocal.InsertOne(context.TODO(), newDevice)
	if err != nil {

		return false, err
	}
	log.Println(insertResult)
	return true, nil
}
