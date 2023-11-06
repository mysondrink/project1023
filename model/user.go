// 存储数据模型
package model

import (
	"gorm.io/gorm"
)

type UserTable struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"type:varchar(11);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}

type Controller struct {
	gorm.Model
	Controller_id   string `gorm:"type:varchar(255);not null"`
	Controller_type string `gorm:"type:varchar(255);not null"`
	Start_time      string `gorm:"type:varchar(255);not null"`
	End_time        string `gorm:"type:varchar(255);not null"`
	Trail_name      string `gorm:"type:varchar(255);not null"`
	Trail_id        uint
	Trail_type      string `gorm:"type:varchar(255);not null"`
	Module_name     string `gorm:"type:varchar(255);not null"`
	Module_id       uint
	Module_type     string `gorm:"type:varchar(255);not null"`
	Car_id          uint
	Create_time     string `gorm:"type:varchar(255);not null"`
	Update_time     string `gorm:"type:varchar(255);not null"`
	Charge_area     uint
	Trans_area      uint
	Status          string `gorm:"type:varchar(255);not null"`
	Tx_status       string `gorm:"type:varchar(255);not null"`
	Rx_status       string `gorm:"type:varchar(255);not null"`
	Rates           string `gorm:"type:varchar(255);not null"`
}
