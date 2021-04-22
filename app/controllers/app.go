package controllers

import (
	"github.com/revel/revel"
  gormc "github.com/revel/modules/orm/gorm/app/controllers"
  gorm "github.com/revel/modules/orm/gorm/app"
  "accounting/app/models"
)

type App struct {
  gormc.TxnController
}

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

func (c App) Index() revel.Result {
  
  var assets = []models.Asset{}
  c.Txn.Find(&assets)
  
  assetTotal, liabilityTotal, netWorth := calculateTotals(assets)
  
	return c.Render(assets, assetTotal, liabilityTotal, netWorth)
}

func (c App) Delete(id uint) revel.Result {
  gorm.DB.Delete(&models.Asset{}, id)
  
  return c.Redirect(App.Index)
}

func (c App) Add(name string, aorl string, balance float64) revel.Result {
  
  c.Validation.Required(name).Message("A name for the asset/liability is required!")
  c.Validation.Required(aorl).Message("The type ('Asset' or 'Liability') is required!")
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
