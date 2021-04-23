package controllers

import (
  "github.com/revel/revel"
	gorm "github.com/revel/modules/orm/gorm/app"
  "accounting/app/models"
)

/**
 * Initializes the database based on the models present.
 */
func initializeDB() {
  gorm.DB.AutoMigrate(&models.Asset{})
}

/**
 * Sets initializeDB as a function to run on application start up.
 */
func init() {
  revel.OnAppStart(initializeDB)
}
