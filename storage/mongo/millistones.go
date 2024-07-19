package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	pb "github.com/mirjalilova/TimelineService/genproto/timeline"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MillistonesRepo struct {
	Millistones *mongo.Collection
}

func NewMillistonesRepo(db *mongo.Database) *MillistonesRepo {
	return &MillistonesRepo{
		Millistones: db.Collection("millistones"),
	}
}

func (m *MillistonesRepo) Create(req *pb.MillistonesCreate) (*pb.Void, error) {
	res := &pb.Void{}

	millistones := &pb.Millistone{
		Id:          uuid.New().String(),
		Title:       req.Title,
		Date:        req.Date,
		Description: req.Description,
		UserId:      req.UserId,
		Category:    req.Category,
		CreateAt:    time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt:    time.Now().Format("2006-01-02 15:04:05"),
		DeletedAt:   0,
	}

	_, err := m.Millistones.InsertOne(context.TODO(), millistones)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *MillistonesRepo) Update(req *pb.MillistonesUpdate) (*pb.Void, error) {
	res := &pb.Void{}

	filter := bson.M{"id": req.Id}
	filter = bson.M{"deletedat": 0}
	filter = bson.M{"userid": req.UserId}

	update := bson.D{
		{"$set", bson.D{{"updated_at", time.Now().Format("2006-01-02 15:04:05")}}},
	}

	if req.Title != "" && req.Title != "string" {
		update = bson.D{
			{"$set", bson.D{{"title", req.Title}}},
		}
	}
	if req.Date != "" && req.Date != "string" {
		update = bson.D{
			{"$set", bson.D{{"date", req.Date}}},
		}
	}
	if req.Description != "" && req.Description != "string" {
		update = bson.D{
			{"$set", bson.D{{"description", req.Description}}},
		}
	}
	if req.Category != "" && req.Category != "string" {
		update = bson.D{
			{"$set", bson.D{{"category", req.Category}}},
		}
	}

	_, err := m.Millistones.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *MillistonesRepo) Delete(req *pb.GetById) (*pb.Void, error) {
	res := &pb.Void{}

	filter := bson.M{"id": req.Id}
	filter = bson.M{"deletedat": 0}
	filter = bson.M{"userid": req.UserId}

	update := bson.D{
		{"$set", bson.D{{"deletedat", time.Now().Unix()}}},
	}

	_, err := m.Millistones.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *MillistonesRepo) Get(req *pb.GetById) (*pb.Millistones, error) {
	res := &pb.Millistones{}

	filter := bson.M{
		"id":         req.Id,
		"deletedat": 0,
		"userid":    req.UserId,
	}

	err := m.Millistones.FindOne(context.TODO(), filter).Decode(res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("no millistone found with the given ID")
		}
		return nil, err
	}

	return res, nil
}

func (m *MillistonesRepo) GetAll(req *pb.GetAllReq) (*pb.GetAllRes, error) {
	res := &pb.GetAllRes{}

	filter := bson.M{"deletedat": 0}
	filter = bson.M{"userid": req.UserId}

	if req.Category != "" {
		filter["category"] = req.Category
	}
	if req.Date != "" {
		filter["date"] = req.Date
	}

	cur, err := m.Millistones.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var millistone pb.Millistones
		err := cur.Decode(&millistone)
		if err != nil {
			return nil, err
		}
		res.Millistones = append(res.Millistones, &millistone)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	res.TotalCount = int32(len(res.Millistones))

	return res, nil
}

func (m *MillistonesRepo) GetByDateMillistones(req *pb.GetByDate) (*pb.GetAllRes, error) {
	res := &pb.GetAllRes{}

	filter := bson.M{
		"$and": []bson.M{
			{"deletedat": 0},
			{"userid": req.UserId},
			{"date": bson.M{"$gte": req.FromDate, "$lte": req.ToDate}},
		},
	}

	cur, err := m.Millistones.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var millistone pb.Millistones
		err := cur.Decode(&millistone)
		if err != nil {
			return nil, err
		}
		res.Millistones = append(res.Millistones, &millistone)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return res, nil
}
