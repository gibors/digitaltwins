package repository

import (
	dev "caidc_auto_devicetwins/domain/model"
	"caidc_auto_devicetwins/domain/utils"
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

const DEVICETOREGISTER = "mongodb://ChuladaUser:Honeywe!!Up!nThe$ky786@devcaidcdb-01.westus.cloudapp.azure.com:24403/automationdb?readPreference=primary&maxIdleTimeMS=60000"

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

func (c *MongoClient) GetDeviceInserteOnPrem(serialNumber string) (dev.Device, error) {
	filter := bson.D{{"serialNumber", serialNumber}}
	var device = dev.Device{}
	err := c.DMCollectionLocal.FindOne(context.TODO(), filter).Decode(&device)
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

func (c *MongoClient) GetDevicesFromDeviceToRegisterCollection() dev.DeviceToRegister {
	client, _ := GetMongoClient(DEVICETOREGISTER)
	deviceToRegister := client.Database("automationdb").Collection("DeviceToRegister")

	filter := bson.D{{"Status", 0}}
	var device = dev.DeviceToRegister{}
	err := deviceToRegister.FindOne(context.TODO(), filter).Decode(&device)
	utils.FailOnError(err, "Failed when getting device to register ")
	client.Disconnect(context.TODO())
	return device
}

func (c *MongoClient) UpdateDeviceToRegister(deviceTR dev.DeviceToRegister, device dev.Device) bool {
	client, _ := GetMongoClient(DEVICETOREGISTER)
	deviceToRegister := client.Database("automationdb").Collection("DeviceToRegister")

	filter := bson.D{{"Status", 0}, {"DeviceId", deviceTR.DeviceID}}
	update := bson.M{
		"$set": bson.M{
			"SystemGuid": device.SystemGUID,
			"Status":     2,
		},
	}
	// update := bson.M{"$set": bson.M{"SystemGuid": device.SystemGUID}}
	result, err := deviceToRegister.UpdateOne(context.TODO(), filter, update)
	log.Printf("Result: %s", result.UpsertedID)
	utils.FailOnError(err, "Failed when getting device to register ")
	client.Disconnect(context.TODO())
	return true
}
