package usecases

import (
	"github.com/Rayato159/go-clean-arch-v2/cockroach/entities"
	"github.com/Rayato159/go-clean-arch-v2/cockroach/models"
)

type CockroachUsecase interface {
	CockroachDataProcessing(in *models.AddCockroachData) error
	CockroachSearchTransactions(in *models.IdCockroachData) (entities.Cockroach, error)
}
