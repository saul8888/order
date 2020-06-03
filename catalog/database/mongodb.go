package database

import (
	"context"
	"fmt"
	"time"

	"github.com/orderforme/catalog/config"
	"github.com/orderforme/catalog/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const (
	mongoDefaultTimeOut   = 5 * time.Second
	defaultCollectionName = "catalog"
)

type MongoDB interface {
	ConnectDB() error
	DisconnectDB() error
	GetByID(ID string) (model.Catalog, error)
	GetAll(params *model.GetLimit) ([]model.Catalog, error)
	GetCantTotal() (int, error)
	CreateNew(params interface{}) error
	Update(ID string, params interface{}) (model.Catalog, error)
	Delete(ID string) error
	ValidateID(table string, ID primitive.ObjectID) error
	AddMarcas(table string, ID primitive.ObjectID) error
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
func (repo *Mongodb) GetByID(ID string) (model.Catalog, error) {
	collection := repo.db.Collection(defaultCollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeOut)
	defer cancel()

	catalog := model.Catalog{}
	objectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		return catalog, err
	}

	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&catalog)
	if err != nil {
		return catalog, err
	}

	return catalog, nil
}

// GetAll method
func (repo *Mongodb) GetAll(params *model.GetLimit) ([]model.Catalog, error) {
	collection := repo.db.Collection(defaultCollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeOut)
	defer cancel()

	catalogs := []model.Catalog{}
	options := options.Find()
	options.SetSkip(int64(params.Offset))
	options.SetLimit(int64(params.Limit))

	cursor, err := collection.Find(ctx, bson.M{}, options)
	if err != nil {
		return catalogs, err
	}

	if err = cursor.All(ctx, &catalogs); err != nil {
		return catalogs, err
	}

	return catalogs, nil
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
func (repo *Mongodb) Update(ID string, params interface{}) (model.Catalog, error) {

	collection := repo.db.Collection(defaultCollectionName)
	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeOut)
	defer cancel()

	updated := model.Catalog{}

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

//AddMarcas method
func (repo *Mongodb) AddMarcas(table string, ID primitive.ObjectID) error {
	collection := repo.db.Collection(table)
	ctx, cancel := context.WithTimeout(context.Background(), mongoDefaultTimeOut)
	defer cancel()
	a := model.Variations{
		Name:     "san carlos",
		Unit:     "bottle",
		Price:    8,
		Currency: "$",
		Stock: model.OptionStock{
			Description: "loss",
			InStock:     34,
			AlertStock:  5,
		},
	}
	fmt.Println(a)
	//updatedCustomer := model.Customer{}
	filter := bson.M{"_id": ID}
	update := bson.M{
		"$push": bson.M{
			//"$addToSet": bson.M{
			//"$each": {"marcas.name": "san carlos", "marcas.price": 8},
			//"marcas.name": "san carlos", "marcas.price": 8,
			//"marcas": bson.M{"name": "san carlos", "price": 8},
			//"$each": {Name: "san carlos", Price: 8},
			//"$each": {"name": "san carlos", "price": 8},
			//"$each": [{"name":"san carlos","price":8},{"name":"san carlos","price":8}],
			//"$each": [{Name:"san carlos",Price:8},{Name:"san carlos",Price:8}],

		},
	}

	//opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	//collection.FindOneAndUpdate(ctx, filter, update, opts)
	collection.UpdateOne(ctx, filter, update)
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

/*
func (repo *MongoCustomerRepo) Update() (model.Customer, error)
	updatedCustomer := model.Customer{}
	objectID, err := primitive.ObjectIDFromHex(customerID)
	if err != nil {
		return model.Customer{}, err
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{
		"$set": bson.M{
			"name":        name,
			"email":       email,
			"embeddings":  embeddings,
			"phoneNumber": phoneNumber,
			"addresses":   addresses,
			"tags":        tags,
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	err = customersCollection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedCustomer)
	return updatedCustomer, err
}
*/
