package postgres

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoomRtp struct {
	ID    primitive.ObjectID
	Time  time.Time `bson:"date"`
	Bet   uint64
	Score uint64
}

type GameRtp struct {
	ID    primitive.ObjectID
	Time  time.Time `bson:"date"`
	Bet   uint64
	Score uint64
}

