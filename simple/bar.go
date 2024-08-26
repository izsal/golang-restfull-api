package simple

type BarRepository struct {
}

func NewBarRepository() *BarRepository {
	return &BarRepository{}
}

type BarService struct {
	*BarRepository
}

func NewBarService() *BarService {
	return &BarService{BarRepository: NewBarRepository()}
}
