package postgres

import (
	"gorm.io/gorm"
	"surlit/internal/logic/interfaces"
)

type HealthCheck struct {
	db *gorm.DB
}

var _ interfaces.HealthCheck = (*HealthCheck)(nil)

func NewHealthCheck(db *gorm.DB) *HealthCheck {
	return &HealthCheck{
		db: db,
	}
}

func (h HealthCheck) Name() string {
	return "database connection"
}

func (h HealthCheck) Exec() (details any, err error) {
	db, err := h.db.DB()
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return nil, nil
}
