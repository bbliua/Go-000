package di

import (
	"toy/internal/biz"
)

type Repo struct {
	*biz.BomUseCase
}

func NewRepo(uc *biz.BomUseCase) (*Repo, error) {
	rep := &Repo{
		uc,
	}
	return rep, nil
}