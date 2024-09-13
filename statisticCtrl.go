package controllers

import (
	interfaces2 "github.com/Jarozin/interfaces2"

	models "github.com/Jarozin/models"
)

type StatisticCtrl struct {
	StatisticService interfaces2.IRepoStatistic
}

func NewStatisticCtrl(Stservice interfaces2.IRepoStatistic) *StatisticCtrl {
	return &StatisticCtrl{StatisticService: Stservice}
}

func (ctrl *StatisticCtrl) GetStatistic() (*models.Statistic, error) {
	return ctrl.StatisticService.GetStatistic()
}

func (ctrl *StatisticCtrl) UpdateStatistic(stat *models.Statistic) error {
	return ctrl.StatisticService.UpdateStatistic(stat)
}
