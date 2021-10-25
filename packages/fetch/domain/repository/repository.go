package repository

import (
	"github.com/bariasabda/monorepo/packages/fetch/config"
	"github.com/bariasabda/monorepo/packages/fetch/domain/entity"
)

type repository struct {
	cfg config.Config
}

type RepositoryInterface interface {
	GetResource() (*[]entity.Data, error)
	CurrencyConverter(from string, to string) (*entity.Currency, error)
}

func NewRepository(config config.Config) *repository {
	return &repository{
		cfg: config,
	}
}
