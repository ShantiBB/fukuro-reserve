package service

type Repository interface {
	HotelRepository
}

type Service struct {
	repo Repository
}

func New(repo Repository) *Service {
	return &Service{repo}
}
