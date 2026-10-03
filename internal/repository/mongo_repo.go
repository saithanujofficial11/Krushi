package repository

import (
	"context"
	"errors"
	"fmt"
	"krushi-server/internal/models"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrFarmerNotFound = errors.New("farmer not found")
)

type MongoFarmerRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoFarmerRepository(uri, dbName, collName string) (*MongoFarmerRepository, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongodb: %w", err)
	}

	// Ping the database to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping mongodb: %w", err)
	}

	collection := client.Database(dbName).Collection(collName)

	return &MongoFarmerRepository{
		client:     client,
		collection: collection,
	}, nil
}

func (r *MongoFarmerRepository) Save(farmer *models.Farmer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if farmer.FarmerID == "" {
		farmer.FarmerID = uuid.New().String()
	}

	_, err := r.collection.InsertOne(ctx, farmer)
	if err != nil {
		return fmt.Errorf("failed to insert farmer: %w", err)
	}

	return nil
}

func (r *MongoFarmerRepository) GetByID(id string) (*models.Farmer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var farmer models.Farmer
		filter := bson.M{"farmer_id": id}

	err := r.collection.FindOne(ctx, filter).Decode(&farmer)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrFarmerNotFound
		}
		return nil, fmt.Errorf("failed to find farmer: %w", err)
	}

	return &farmer, nil
}

func (r *MongoFarmerRepository) ListAll() ([]*models.Farmer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to list farmers: %w", err)
	}
	defer cursor.Close(ctx)

	var farmers []*models.Farmer
	if err := cursor.All(ctx, &farmers); err != nil {
		return nil, fmt.Errorf("failed to decode farmers: %w", err)
	}

	return farmers, nil
}

func (r *MongoFarmerRepository) GetByPhone(phone string) (*models.Farmer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var farmer models.Farmer
	filter := bson.M{"contactInfo.primaryPhone": phone}

	err := r.collection.FindOne(ctx, filter).Decode(&farmer)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrFarmerNotFound
		}
		return nil, fmt.Errorf("failed to find farmer by phone: %w", err)
	}

	return &farmer, nil
}

func (r *MongoFarmerRepository) Close() error {
	return r.client.Disconnect(context.Background())
}
