package kafka

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/mirjalilova/TimelineService/genproto/timeline"
	"github.com/mirjalilova/TimelineService/service"
)

func MillistonesCreateHandler(m *service.MillistonesService) func(message []byte) {
	return func(message []byte) {
		var millistone pb.MillistonesCreate
		if err := json.Unmarshal(message, &millistone); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		_, err := m.Create(context.Background(), &millistone)
		if err != nil {
			log.Printf("Cannot create millistone via Kafka: %v", err)
			return
		}
		log.Printf("Created millistone")
	}
}

func MillistonesUpdateHandler(m *service.MillistonesService) func(message []byte) {
	return func(message []byte) {
		var millistones pb.MillistonesUpdate
		if err := json.Unmarshal(message, &millistones); err != nil {
			log.Printf("Cannot unmarshal JSON: %v", err)
			return
		}

		_, err := m.Update(context.Background(), &millistones)
		if err != nil {
			log.Printf("Cannot update millistones via Kafka: %v", err)
			return
		}
		log.Printf("Updated millistones")
	}
}

