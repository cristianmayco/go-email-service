package domain

import "time"

type AbstractModel struct {
	ID        string    `validate:"required"`
	CreatedAt time.Time `validate:"required"`
	UpdatedAt time.Time `validate:"required"`
}

func (m *AbstractModel) CreateAbstractModel() {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
}
