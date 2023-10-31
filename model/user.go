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

type Trail struct {
	gorm.Model
	Start_time  string `gorm:"type:varchar(255);not null"`
	End_time    string `gorm:"type:varchar(255);not null"`
	Trail_name  string `gorm:"type:varchar(255);not null"`
	Module_type string `gorm:"type:varchar(255);not null"`
	Trail_type  string `gorm:"type:varchar(255);not null"`
	Trail_id    uint
	Car_id      uint
	Car_type    string `gorm:"type:varchar(255);not null"`
	Status      string `gorm:"type:varchar(255);not null"`
	Create_time string `gorm:"type:varchar(255);not null"`
	Update_time string `gorm:"type:varchar(255);not null"`
	Charge_area uint
	Tran_area   uint
}
