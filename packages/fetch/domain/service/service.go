package service

import (
	"github.com/bariasabda/monorepo/packages/fetch/config"
	"github.com/bariasabda/monorepo/packages/fetch/domain/repository"
)

type service struct {
	cfg  config.Config
	repo repository.RepositoryInterface
}

type ServiceInterface interface {
	VerifyToken(reqToken string) (*User, error)
	Aggregator(reqToken string) (*[]AggregatorResponse, error)
	CurrencyConverter(reqToken string) (*[]CurrencyConverterResponse, error)
}

func NewService(config config.Config, repo repository.RepositoryInterface) *service {
	return &service{
		cfg:  config,
		repo: repo,
	}
}
