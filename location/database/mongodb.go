package database

import (
	"context"
	"time"

	"github.com/orderforme/location/config"
	"github.com/orderforme/location/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const (
	mongoDefaultTimeOut   = 5 * time.Second
	defaultCollectionName = "location"
)

type MongoDB interface {
	ConnectDB() error
	DisconnectDB() error
	GetByID(ID string) (model.Location, error)
	GetAll(params *model.GetLimit) ([]model.Location, error)
	GetCantTotal() (int, error)
	CreateNew(params interface{}) error
	Update(ID string, params interface{}) (model.Location, error)
	Delete(ID string) error
	ValidateID(table string, ID primitive.ObjectID) error
	Search(params interface{}) (*mongo.Cursor, context.Context, error)
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
func (repo *Mongodb) GetByID(ID string) (model.Location, error) {
	collection := repo.db.Collection(defaultCollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeOut)
	defer cancel()

	location := model.Location{}
	objectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return location, err
	}

	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&location)
	if err != nil {
		return location, err
	}

	return location, nil
}

// GetAll method
func (repo *Mongodb) GetAll(params *model.GetLimit) ([]model.Location, error) {
	collection := repo.db.Collection(defaultCollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeOut)
	defer cancel()

	locations := []model.Location{}
	options := options.Find()
	options.SetSkip(int64(params.Offset))
	options.SetLimit(int64(params.Limit))

	cursor, err := collection.Find(ctx, bson.M{}, options)
	if err != nil {
		return locations, err
	}

	if err = cursor.All(ctx, &locations); err != nil {
		return locations, err
	}

	return locations, nil
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

// CreateNew method
func (repo *Mongodb) CreateNew(params interface{}) error {
	//connect collection
	collection := repo.db.Collection(defaultCollectionName)

	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeOut)
	defer cancel()

	_, err := collection.InsertOne(ctx, params)
	if err != nil {
		return err
	}

	return nil
}

// Update method
func (repo *Mongodb) Update(ID string, params interface{}) (model.Location, error) {

	collection := repo.db.Collection(defaultCollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeOut)
	defer cancel()

	updated := model.Location{}

	objectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return updated, err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": params}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	err = collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updated)

	return updated, err
}

// Delete method
func (repo *Mongodb) Delete(ID string) error {
	collection := repo.db.Collection(defaultCollectionName)

	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeOut)
	defer cancel()

	objectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objectID})

	return err
}

// ValidateID method
func (repo *Mongodb) ValidateID(table string, ID primitive.ObjectID) error {

	collection := repo.db.Collection(table)
	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeOut)
	defer cancel()

	err := collection.FindOne(ctx, bson.M{"_id": ID}).Err()
	if err != nil {
		return err
	}

	return nil
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
