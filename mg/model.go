package mg

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// only record each day RTP , and compute total RTP on each day 00:00 and store it to redis
type RoomDailyRtp struct {
	ID     primitive.ObjectID
	GameID primitive.ObjectID
	Date   time.Time `bson:"date"`
	Bet    uint64
	Score  uint64
}

type RoomWeeklyRtp struct {
	ID     primitive.ObjectID
	GameID primitive.ObjectID
	Type   string
	Bet    uint64
	Score  uint64
}

type RoomMonthlyRtp struct {
	ID     primitive.ObjectID
	GameID primitive.ObjectID
	Type   string
	Bet    uint64
	Score  uint64
}

type RoomTotalRtp struct {
	ID     primitive.ObjectID
	GameID primitive.ObjectID
	Type   string
	Bet    uint64
	Score  uint64
}

type GameDailyRtp struct {
	ID    primitive.ObjectID
	Bet   uint64
	Score uint64
}

type GameWeeklyRtp struct {
	ID    primitive.ObjectID
	Type  string
	Bet   uint64
	Score uint64
}

type GameMonthlyRtp struct {
	ID    primitive.ObjectID
	Type  string
	Bet   uint64
	Score uint64
}

type GameTotalRtp struct {
	ID    primitive.ObjectID
	Type  string
	Bet   uint64
	Score uint64
}

func CreateCollection() {
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("Mongo connect error:", err)
	}
	defer client.Disconnect(ctx)

	// Ping the database to ensure connection is established
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("Mongo ping error:", err)
	}

	db := client.Database("your_db_name")
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "date", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	name, err := db.Collection("rtp").Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Fatalf("Failed to create daily RTP index: %v", err)
	}
	log.Printf("Created daily RTP index: %s\n", name)
}
