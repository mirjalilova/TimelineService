package service

import (
	"context"

	pb "github.com/mirjalilova/TimelineService/genproto/timeline"
	st "github.com/mirjalilova/TimelineService/storage/mongo"
)

type MillistonesService struct {
	storage st.Storage
	pb.UnimplementedMillistonesServiceServer
}

func NewMillistonesService(storage *st.Storage) *MillistonesService {
	return &MillistonesService{
		storage: *storage,
	}
}

func (s *MillistonesService) Create(ctx context.Context, req *pb.MillistonesCreate) (*pb.Void, error) {
    return s.storage.Millistones().Create(req)
}

func (s *MillistonesService) Update(ctx context.Context, req *pb.MillistonesUpdate) (*pb.Void, error) {
    return s.storage.Millistones().Update(req)
}

func (s *MillistonesService) Delete(ctx context.Context, req *pb.GetById) (*pb.Void, error) {
    return s.storage.Millistones().Delete(req)
}

func (s *MillistonesService) Get(ctx context.Context, req *pb.GetById) (*pb.Millistones, error) {
    return s.storage.Millistones().Get(req)
}

func (s *MillistonesService) GetAll(ctx context.Context, req *pb.GetAllReq) (*pb.GetAllRes, error) {
    return s.storage.Millistones().GetAll(req)
}

func (s *MillistonesService) GetByDateMillistones(ctx context.Context, req *pb.GetByDate) (*pb.GetAllRes, error) {
    return s.storage.Millistones().GetByDateMillistones(req)
}