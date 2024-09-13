package controllers

import (
	interfaces2 "github.com/Jarozin/interfaces2"
	models "github.com/Jarozin/models"
)

type FavouritesCtrl struct {
	FavouritesService interfaces2.IRepoFavourites
}

func NewFavouritesCtrl(service interfaces2.IRepoFavourites) *FavouritesCtrl {
	return &FavouritesCtrl{FavouritesService: service}
}

func (ctrl *FavouritesCtrl) GetFavourites() ([]*models.Favourites, error) {
	return ctrl.FavouritesService.GetFavourites()
}

func (ctrl *FavouritesCtrl) GetFavouriteById(id int) (*models.Favourites, error) {
	return ctrl.FavouritesService.GetFavouriteById(id)
}

func (ctrl *FavouritesCtrl) UpdateFavourite(favourite *models.Favourites) error {
	return ctrl.FavouritesService.UpdateFavourite(favourite)
}

func (ctrl *FavouritesCtrl) DeleteFavourite(id int) error {
	return ctrl.FavouritesService.DeleteFavourite(id)
}
