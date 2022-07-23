package invoiceitem

import "time"

// Model of invoiceitem
type Model struct {
	ID              uint
	InvoiceheaderID uint
	ProductID       uint
	CreateAt        time.Time
	UpdateAt        time.Time
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
