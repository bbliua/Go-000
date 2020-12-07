package service

import (
	"context"

	"github.com/google/wire"
	"week.02/internal/dao"
)

var Provider = wire.NewSet(New)

type Service struct {
	dao dao.Dao
}

type BomResp struct {
	Id       int64
	FileName string
}

func New(d dao.Dao) (s *Service, err error) {
	s = &Service{
		dao: d,
	}
	return
}

func (s *Service) FetchBom(ctx context.Context, id int64) (reply *BomResp, err error) {
	bom, err := s.dao.Bom(ctx, id)
	reply = &BomResp{}
	if err != nil {
		return
	}
	reply.Id = bom.ID
	reply.FileName = bom.FileName
	return
}
