package controllers

import (
	interfaces2 "github.com/Jarozin/interfaces2"
	models "github.com/Jarozin/models"
)

type EpisodesCtrl struct {
	EpisodesService interfaces2.IRepoEpisodes
}

func NewEpisodesCtrl(service interfaces2.IRepoEpisodes) *EpisodesCtrl {
	return &EpisodesCtrl{EpisodesService: service}
}

func (ctrl *EpisodesCtrl) GetEpisodes() ([]*models.Episodes, error) {
	return ctrl.EpisodesService.GetEpisodes()
}

func (ctrl *EpisodesCtrl) GetEpisodeById(id int) (*models.Episodes, error) {
	return ctrl.EpisodesService.GetEpisodeById(id)
}

func (ctrl *EpisodesCtrl) GetEpisodesBySeasonId(id int) ([]*models.Episodes, error) {
	return ctrl.EpisodesService.GetEpisodesBySeasonId(id)
}

func (ctrl *EpisodesCtrl) CreateEpisode(episode *models.Episodes) error {
	return ctrl.EpisodesService.CreateEpisode(episode)
}

func (ctrl *EpisodesCtrl) UpdateEpisode(episode *models.Episodes) error {
	return ctrl.EpisodesService.UpdateEpisode(episode)
}

func (ctrl *EpisodesCtrl) DeleteEpisode(id int) error {
	return ctrl.EpisodesService.DeleteEpisode(id)
}
