package product

import (
	"time"

	"github.com/pkg/errors"
)

var (
	ErrIDNotFound = errors.New("El producto no contiene un ID")
)

//Model of product
type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Models []*Model

type Storage interface {
	Migrate() error
	// 	Create(*Model) error
	// 	Update(*Model) error
	// 	GetAll() (Models, error)
	// 	GetByID(uint) (*Model, error)
	// 	Delete(uint) error
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

// func (s *Service) Create(m *Model) error {
// 	m.CreatedAt = time.Now()
// 	return s.storage.Create(m)
// }

// func (m *Model) String() string {
// 	return fmt.Sprintf("%02d | %-20s | %-20s | %5d | %10s | %10s\n",
// 		m.ID,
// 		m.Name,
// 		m.Observations,
// 		m.Price,
// 		m.CreatedAt.Format("2006-01-02"),
// 		m.UpdatedAt.Format("2006-01-02"),
// 	)
// }

// func (s *Service) GetAll() (Models, error) {
// 	return s.storage.GetAll()
// }

// func (s *Service) GetByID(id uint) (*Model, error) {
// 	return s.storage.GetByID(id)
// }

// func (s *Service) Update(m *Model) error {
// 	if m.ID == 0 {
// 		return ErrIDNotFound
// 	}
// 	m.UpdatedAt = time.Now()

// 	return s.storage.Update(m)
// }

// func (s *Service) Delete(id uint) error {
// 	return s.storage.Delete(id)
// }
