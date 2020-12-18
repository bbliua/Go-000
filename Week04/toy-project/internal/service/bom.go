package service

import (
	"context"
	v1 "toy/api/bom/v1"
	"toy/internal/biz"

	"log"
)


type BomService struct {
	v1.UnimplementedBomServer
	ouc *biz.BomUseCase
}

func NewBomService(ouc *biz.BomUseCase) v1.BomServer  {
	return &BomService{ouc:ouc}
}

func (s *BomService) CreateBom(ctx context.Context, req *v1.CreateBomRequest) (*v1.CreateBomReply, error)  {
	b := new(biz.Bom)

	b.FileName = req.FileName
	s.ouc.Add(b)
	log.Println(req)
	return &v1.CreateBomReply{
		Message: "succ",
	}, nil
}