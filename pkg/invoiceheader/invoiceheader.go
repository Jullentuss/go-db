package invoiceheader

import "time"

// Model of invoice-header
type Model struct {
	ID       uint
	Client   string
	CreateAt time.Time
	UpdateAt time.Time
}

type Storage interface {
	Migrate() error
}

type Service struct {
	storage Storage
}

// NewService Constructor
func NewService(s Storage) *Service {
	return &Service{s}
}

func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
