package usecase

type Usecase struct {
	repositories BookRepository
}

func NewBookUsecase(repo BookRepository) *Usecase {
	return &Usecase{
		repositories: repo,
	}
}
