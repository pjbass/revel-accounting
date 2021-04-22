package controllers

import (
  "github.com/revel/revel"
	gorm "github.com/revel/modules/orm/gorm/app"
  "accounting/app/models"
)

func initializeDB() {
  gorm.DB.AutoMigrate(&models.Asset{})
}

func init() {
  revel.OnAppStart(initializeDB)
}
