package controllers

import (
	interfaces2 "github.com/Jarozin/interfaces2"
	models "github.com/Jarozin/models"
)

type SeasonsCtrl struct {
	SeasonsService interfaces2.IRepoSeasons
}

func NewSeasonsCtrl(service interfaces2.IRepoSeasons) *SeasonsCtrl {
	return &SeasonsCtrl{SeasonsService: service}
}

func (ctrl *SeasonsCtrl) GetSeasons() ([]*models.Seasons, error) {
	return ctrl.SeasonsService.GetSeasons()
}

func (ctrl *SeasonsCtrl) GetSeasonById(id int) (*models.Seasons, error) {
	return ctrl.SeasonsService.GetSeasonById(id)
}

func (ctrl *SeasonsCtrl) GetSeasonsBySerialId(id int) ([]*models.Seasons, error) {
	return ctrl.SeasonsService.GetSeasonsBySerialId(id)
}

func (ctrl *SeasonsCtrl) CreateSeason(season *models.Seasons) error {
	return ctrl.SeasonsService.CreateSeason(season)
}

func (ctrl *SeasonsCtrl) UpdateSeason(season *models.Seasons) error {
	return ctrl.SeasonsService.UpdateSeason(season)
}

func (ctrl *SeasonsCtrl) DeleteSeason(id int) error {
	return ctrl.SeasonsService.DeleteSeason(id)
}
