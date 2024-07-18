package mongo

import (
	"context"

	"github.com/mirjalilova/TimelineService/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	DB *mongo.Database
	MillistonesS storage.MillistonesI
}

func ConnectMongo() (*Storage,error) {
	clientOptions := options.Client(

	).ApplyURI("mongodb://mongo:27017").
	SetAuth(options.Credential{Username: "mongo", Password: "feruza1727"})

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err!= nil {
        return nil,err
    }
	err = client.Ping(context.TODO(), nil)
	if err!= nil {
        return nil,err
    }
	db := client.Database("timecapsule")

	return &Storage{
		DB: db,
        MillistonesS: NewMillistonesRepo(db),
	}, nil
}
func (s *Storage) Millistones() (storage.MillistonesI) {
	if s.MillistonesS == nil {
		s.MillistonesS = NewMillistonesRepo(s.DB)
	}
	return s.MillistonesS
}