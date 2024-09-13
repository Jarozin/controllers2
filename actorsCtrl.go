package controllers

import (
	interfaces2 "github.com/Jarozin/interfaces2"
	models "github.com/Jarozin/models"
)

type ActorsCtrl struct {
	ActorsService interfaces2.IRepoActors
}

func NewActorsCtrl(service interfaces2.IRepoActors) *ActorsCtrl {
	return &ActorsCtrl{ActorsService: service}
}

func (ctrl *ActorsCtrl) GetActors() ([]*models.Actors, error) {
	return ctrl.ActorsService.GetActors()
}

func (ctrl *ActorsCtrl) GetActorById(id int) (*models.Actors, error) {
	return ctrl.ActorsService.GetActorById(id)
}

func (ctrl *ActorsCtrl) CreateActor(actor *models.Actors) error {
	return ctrl.ActorsService.CreateActor(actor)
}

func (ctrl *ActorsCtrl) UpdateActor(actor *models.Actors) error {
	return ctrl.ActorsService.UpdateActor(actor)
}

func (ctrl *ActorsCtrl) DeleteActor(id int) error {
	return ctrl.ActorsService.DeleteActor(id)
}

func (ctrl *ActorsCtrl) CheckActor(actor *models.Actors) bool {
	return ctrl.ActorsService.CheckActor(actor)
}
