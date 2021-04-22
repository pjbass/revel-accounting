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

func (c App) Index() revel.Result {
  
  var assets = []models.Asset{}
  c.Txn.Find(&assets)
  
	return c.Render(assets)
}

func (c App) Add(name string, aorl string, balance float32) revel.Result {
  
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
