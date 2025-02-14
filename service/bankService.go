package service

import (
	model "main/model"
	_bankRepository "main/repository"
)

type BankService interface {
	GetBankData() ([]model.Bank, error)
}

type BankServiceImpl struct {
	Repository *_bankRepository.BankRepository
}

func NewBankServiceImpl(_repository *_bankRepository.BankRepository) *BankServiceImpl {
	return &BankServiceImpl{Repository: _repository}
}

func (s *BankServiceImpl) GetBankData() ([]model.Bank, error) {
	result, err := s.Repository.GetBankData()
	if err != nil {
		return nil, err
	}
	return result, nil
}
