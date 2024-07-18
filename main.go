package main

import (
	"log"
	"net"

	"golang.org/x/exp/slog"

	cf "github.com/mirjalilova/TimelineService/config"
	"github.com/mirjalilova/TimelineService/service"

	"github.com/mirjalilova/TimelineService/kafka"

	pb "github.com/mirjalilova/TimelineService/genproto/timeline"
	"github.com/mirjalilova/TimelineService/storage/mongo"

	"google.golang.org/grpc"
)

func main() {
	config := cf.Load()
	db, err := mongo.ConnectMongo()
	if err != nil {
		slog.Error("can't connect to db: %v", err)
		return
	}

	millistonesService := service.NewMillistonesService(db)

	brokers := []string{"localhost:9092"}

	kcm := kafka.NewKafkaConsumerManager()

	if err := kcm.RegisterConsumer(brokers, "create-millistones", "eval", kafka.MillistonesCreateHandler(millistonesService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'create-millistones' already exists")
		} else {
			log.Fatalf("Error registering consumer: %v", err)
		}
	}

	if err := kcm.RegisterConsumer(brokers, "update-millistones", "eval", kafka.MillistonesCreateHandler(millistonesService)); err != nil {
		if err == kafka.ErrConsumerAlreadyExists {
			log.Printf("Consumer for topic 'update-millistones' already exists")
		} else {
			log.Fatalf("Error registering consumer: %v", err)
		}
	}

	listener, err := net.Listen("tcp", config.TIMELINE_PORT)
	if err != nil {
		slog.Error("can't listen: %v", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterMillistonesServiceServer(s, service.NewMillistonesService(db))

	slog.Info("server started port", config.TIMELINE_PORT)
	if err := s.Serve(listener); err != nil {
		slog.Error("can't serve: %v", err)
		return
	}
}
