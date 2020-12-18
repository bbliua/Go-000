package biz

type Bom struct {
	Id int32
	FileName string
}

type BomRepo interface {
	SaveBom(bom *Bom) error
}

func NewBomUseCase(repo BomRepo) *BomUseCase {
	return &BomUseCase{repo: repo}
}

type BomUseCase struct {
	repo BomRepo
}

func (uc *BomUseCase) Add(b *Bom)  {
	// Todo 验证等逻辑
	uc.repo.SaveBom(b)
}