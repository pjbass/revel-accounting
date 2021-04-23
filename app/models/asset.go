package models

import (
  "github.com/jinzhu/gorm"
)

/**
 * Structure modeling an asset or a liability. Contains the name of the asset
 * or liability, the type (either 'Asset' or 'Liability') and the balance on
 * the asset or liability.
 */
type Asset struct {
  gorm.Model 
  Name string `gorm:"size:64"`
  Type string `gorm:"size:16"`
  Balance float64
}
