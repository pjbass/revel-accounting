package controllers

import (
	"github.com/revel/revel"
  gormc "github.com/revel/modules/orm/gorm/app/controllers"
  gorm "github.com/revel/modules/orm/gorm/app"
  "accounting/app/models"
)

/**
 * App struct that acts as the main controller for the application.
 * Handles adding, deleting, and displaying assets and liabilities.
 */
type App struct {
  gormc.TxnController
}

/**
 * Function used to calculate total assets, total liabilities, and net worth.
 * Takes an array of Asset models and returns the total assets, total 
 * liabilities, and net worth, in that order.
 */
func calculateTotals(assets []models.Asset) (float64, float64, float64) {
  
  var (
    assetTotal float64 = 0.0
    liabilityTotal float64 = 0.0
    netWorth float64 = 0.0
  )
  
  for _, asset := range assets {
    if asset.Type == "Asset" {
      assetTotal += asset.Balance
    } else {
      liabilityTotal += asset.Balance
    }
  }
  
  netWorth = assetTotal - liabilityTotal
  return assetTotal, liabilityTotal, netWorth
}

/**
 * The main page of the application. Displays assets and liabilities, as well
 * as total assets, total liabilities, and net worth.
 */
func (c App) Index() revel.Result {
  
  var assets = []models.Asset{}
  c.Txn.Find(&assets)
  
  assetTotal, liabilityTotal, netWorth := calculateTotals(assets)
  
	return c.Render(assets, assetTotal, liabilityTotal, netWorth)
}

/**
 * Endpoint that allows deleting an asset by id. Redirects to the index
 * upon completion.
 */
func (c App) Delete(id uint) revel.Result {
  
  gorm.DB.Delete(&models.Asset{}, id)
  
  return c.Redirect(App.Index)
}

/**
 * Endpoint that allows adding an asset or liability. Validates input before
 * adding, and will flash a warning when validation fails. Redirects to the
 * index page.
 */
func (c App) Add(name string, aorl string, balance float64) revel.Result {
  
  c.Validation.Required(name).Message("A name for the asset/liability is required!")
  c.Validation.Required(aorl).Message("The type ('Asset' or 'Liability') is required!")
  c.Validation.Required(aorl == "Asset" || aorl == "Liability").Message("The type must be either an 'Asset' or a 'Liability'")
  c.Validation.Required(balance).Message("The balance for the asset/liability is required!")
  
  if c.Validation.HasErrors() {
    c.Validation.Keep()
    c.FlashParams()
  } else {
    var asset = models.Asset{Name: name, Type: aorl, Balance: balance}
    gorm.DB.Create(&asset)
  }
  
  return c.Redirect(App.Index)
}
