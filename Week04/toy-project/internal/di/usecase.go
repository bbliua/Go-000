package di

import (
	"toy/internal/biz"
)

type UseCase struct {
	*biz.BomUseCase
}

func NewUseCase(buc *biz.BomUseCase) (*UseCase, error) {
	uc := &UseCase{
		buc,
	}
	return uc, nil
}