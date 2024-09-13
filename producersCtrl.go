package controllers

import (
	interfaces2 "github.com/Jarozin/interfaces2"
	models "github.com/Jarozin/models"
)

type ProducersCtrl struct {
	ProducersService interfaces2.IRepoProducers
}

func NewProducersCtrl(service interfaces2.IRepoProducers) *ProducersCtrl {
	return &ProducersCtrl{ProducersService: service}
}

func (ctrl *ProducersCtrl) GetProducers() ([]*models.Producers, error) {
	return ctrl.ProducersService.GetProducers()
}

func (ctrl *ProducersCtrl) GetProducerById(id int) (*models.Producers, error) {
	return ctrl.ProducersService.GetProducerById(id)
}

func (ctrl *ProducersCtrl) CreateProducer(producer *models.Producers) error {
	return ctrl.ProducersService.CreateProducer(producer)
}

func (ctrl *ProducersCtrl) UpdateProducer(producer *models.Producers) error {
	return ctrl.ProducersService.UpdateProducer(producer)
}

func (ctrl *ProducersCtrl) DeleteProducer(id int) error {
	return ctrl.ProducersService.DeleteProducer(id)
}
