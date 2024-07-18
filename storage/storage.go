package storage

import pb "github.com/mirjalilova/TimelineService/genproto/timeline"

type StorageI interface {
	Millistones() MillistonesI
}

type MillistonesI interface {
	Create(*pb.MillistonesCreate) (*pb.Void, error)
	Update(*pb.MillistonesUpdate) (*pb.Void, error)
	Delete(*pb.GetById) (*pb.Void, error)
	Get(*pb.GetById) (*pb.Millistones, error)
	GetAll(*pb.GetAllReq) (*pb.GetAllRes, error)
	GetByDateMillistones(*pb.GetByDate) (*pb.GetAllRes, error)
}