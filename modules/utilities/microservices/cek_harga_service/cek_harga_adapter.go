package cekhargaservice

import (
	models "jojonomic/modules/utilities/models/harga"
	hargaRepository "jojonomic/modules/utilities/repositories/harga"
)

type Service interface {
	CekHarga() (models.Harga, error)
}

type service struct {
	hargaRepository hargaRepository.Repository
}

func NewService(hargaRepository hargaRepository.Repository) *service {
	return &service{hargaRepository}
}
