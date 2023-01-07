package usecase

type Usecase struct {
	repositories UserRepository
}

func NewLibraryUsecase(repo UserRepository) *Usecase {
	return &Usecase{
		repositories: repo,
	}
}
