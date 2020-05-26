package database

import (
	"context"
	"fmt"
	"time"

	"github.com/orderforme/merchant/config"
	"github.com/orderforme/merchant/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const (
	mongoDefaultTimeOut   = 5 * time.Second
	defaultCollectionName = "merchant"
)

type MongoDB interface {
	ConnectDB() error
	DisconnectDB() error
	GetByID(merchantID string) (*mongo.SingleResult, error)
	GetAll(params *model.GetLimit) (*mongo.Cursor, context.Context, error)
	GetCantTotal() (int, error)
	Search(params interface{}) (*mongo.Cursor, context.Context, error)
	CreateNew(params interface{}) (string, error)
	Update(merchantID string, params interface{}) (model.Merchant, error)
	Delete(merchantID string) error
}

type Mongodb struct {
	client  *mongo.Client
	context context.Context
	db      *mongo.Database
}

// Connect method
func (repo *Mongodb) ConnectDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeOut)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.AppConfig.DatabaseURI))
	defer cancel()
	if err != nil {
		return err
	}

	database := client.Database(config.AppConfig.DatabaseName)

	repo.client = client
	repo.context = ctx
	repo.db = database

	return nil
}

// Disconnect method
func (repo *Mongodb) DisconnectDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeOut)
	defer cancel()
	err := repo.client.Disconnect(ctx)
	return err
}

// GetByID method
func (repo *Mongodb) GetByID(merchantID string) (*mongo.SingleResult, error) {
	collection := repo.db.Collection(defaultCollectionName)

	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeOut)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(merchantID)
	if err != nil {
		return nil, err
	}

	find := collection.FindOne(ctx, bson.M{"_id": objectID})

	return find, nil
}

// GetAll method
func (repo *Mongodb) GetAll(params *model.GetLimit) (*mongo.Cursor, context.Context, error) {
	collection := repo.db.Collection(defaultCollectionName)

	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeOut)
	defer cancel()

	options := options.Find()
	options.SetSkip(int64(params.Offset))
	options.SetLimit(int64(params.Limit))

	cursor, err := collection.Find(ctx, bson.M{}, options)
	if err != nil {
		return nil, nil, err
	}

	return cursor, ctx, nil
}

//GetCantTotal method
func (repo *Mongodb) GetCantTotal() (int, error) {
	//connect collection
	collection := repo.db.Collection(defaultCollectionName)

	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeOut)
	defer cancel()

	total, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return 0, err
	}

	return int(total), nil
}

//Search method
func (repo *Mongodb) Search(params interface{}) (*mongo.Cursor, context.Context, error) {
	collection := repo.db.Collection(defaultCollectionName)

	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeOut)
	defer cancel()
	cursor, err := collection.Find(ctx, params)
	if err != nil {
		return nil, nil, err
	}

	return cursor, ctx, nil
}

// CreateNew method
func (repo *Mongodb) CreateNew(params interface{}) (string, error) {
	//connect collection
	collection := repo.db.Collection(defaultCollectionName)

	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeOut)
	defer cancel()

	result, err := collection.InsertOne(ctx, params)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", result.InsertedID), nil
}

// Update method
func (repo *Mongodb) Update(merchantID string, params interface{}) (model.Merchant, error) {

	collection := repo.db.Collection(defaultCollectionName)

	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeOut)
	defer cancel()

	updatedMerchant := model.Merchant{}

	objectID, err := primitive.ObjectIDFromHex(merchantID)
	if err != nil {
		return model.Merchant{}, err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": params}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	err = collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedMerchant)

	return updatedMerchant, err
}

// Delete method
func (repo *Mongodb) Delete(merchantID string) error {
	collection := repo.db.Collection(defaultCollectionName)

	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeOut)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(merchantID)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objectID})

	return err
}
