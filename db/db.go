package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/MyRetail/common"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var EMPTY_DATA = errors.New("EMPTY_VALUE")

var db *mongo.Database
var prodColl *mongo.Collection
var isConnectedtoDB bool

const (
	DB_ADDR            = "mongodb://localhost:27017"
	DB_NAME            = "MYRETAIL"
	PRODUCT_COLLECTION = "PRODUCTS"
)

func init() {

	client, e := mongo.Connect(context.TODO(), options.Client().ApplyURI(DB_ADDR))

	if e != nil {
		fmt.Println("not able to connect to db", e)
	} else if e = client.Ping(context.TODO(), nil); e != nil {
		fmt.Println("not able to reach  db", e)
	} else {
		isConnectedtoDB = true
		db = client.Database(DB_NAME)
		prodColl = db.Collection(PRODUCT_COLLECTION)
	}
}

func UpdateProduct(pro common.Product) error {
	if !isConnectedtoDB {
		return errors.New("NOT_CONNECTED_TO_DB")
	}
	p := new(common.Product)
	filter := bson.M{"ID": pro.ID}
	err := prodColl.FindOne(context.TODO(), filter).Decode(p)
	if err != nil && err == mongo.ErrNoDocuments {
		_, err = prodColl.InsertOne(context.TODO(), pro.Getbson())
		return err
	} else if err != nil {
		return err
	}
	_, err = prodColl.UpdateOne(context.TODO(), filter, bson.M{"$set": pro.Getbson()})
	fmt.Println("updated product", pro.Name, err)
	return err
}

func GetProduct(id string) (*common.Product, error) {
	if !isConnectedtoDB {
		return nil, errors.New("NOT_CONNECTED_TO_DB")
	}
	pro := new(common.Product)
	filter := bson.M{"ID": id}
	err := prodColl.FindOne(context.TODO(), filter).Decode(pro)
	if err != nil {
		return nil, err
	}
	return pro, nil
}
