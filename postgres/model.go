package postgres

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// only record each day RTP , and compute total RTP on each day 00:00 and store it to redis
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
