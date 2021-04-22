package models

import (
  "github.com/jinzhu/gorm"
)

type Asset struct {
  gorm.Model 
  Name string `gorm:"size:64"`
  Type string `gorm:"size:16"`
  Balance float64
}
