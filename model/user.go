// 存储数据模型
package model

import (
	"time"

	"gorm.io/gorm"
)

type UserTable struct {
	gorm.Model
	Name      string    `gorm:"type:varchar(20);not null"`
	Telephone string    `gorm:"type:varchar(11);not null;unique"`
	Password  string    `gorm:"size:255;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Controller struct {
	gorm.Model
	Controller_id   string `gorm:"type:varchar(255);not null"`
	Controller_type string `gorm:"type:varchar(255);not null"`
	Start_time      string `gorm:"type:varchar(255);"`
	End_time        string `gorm:"type:varchar(255);"`
	Trail_name      string `gorm:"type:varchar(255);"`
	Trail_id        uint
	Trail_type      string `gorm:"type:varchar(255);"`
	Module_name     string `gorm:"type:varchar(255);"`
	Module_id       uint
	Module_type     string `gorm:"type:varchar(255);"`
	Car_id          uint
	// Create_time     string `gorm:"type:varchar(255);"`
	// Update_time     string `gorm:"type:varchar(255);"`
	Charge_area uint
	Trans_area  uint
	Status      string    `gorm:"type:varchar(255);"`
	Tx_status   string    `gorm:"type:varchar(255);"`
	Rx_status   string    `gorm:"type:varchar(255);"`
	Rates       string    `gorm:"type:varchar(255);"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
}

type Car struct {
	gorm.Model
	Start_time    string `gorm:"type:varchar(255);"`
	End_time      string `gorm:"type:varchar(255);"`
	Car_name      string `gorm:"type:varchar(255);"`
	Car_id        uint
	Car_type      string `gorm:"type:varchar(255);"`
	Sample_status string `gorm:"type:varchar(255);"`
	// Create_time   string `gorm:"type:varchar(255);"`
	// Update_time   string `gorm:"type:varchar(255);"`
	Charge_area   uint
	Trans_area    uint
	Status        string `gorm:"type:varchar(255);"`
	Sample_name   string `gorm:"type:varchar(255);"`
	Sample_type   string `gorm:"type:varchar(255);"`
	Sample_id     uint
	Sample_code   uint
	Power         uint
	Target        string `gorm:"type:varchar(255);"`
	Position      string `gorm:"type:varchar(255);"`
	Valid         uint
	Username      string    `gorm:"type:varchar(255);"`
	Controller_id string    `gorm:"type:varchar(255);"`
	CreatedAt     time.Time `gorm:"not null"`
	UpdatedAt     time.Time `gorm:"not null"`
}

type Toy struct {
	gorm.Model
	Start_time string `gorm:"type:varchar(255);"`
	End_time   string `gorm:"type:varchar(255);"`
	// Car_name      string `gorm:"type:varchar(255);"`
	Toy_id     uint
	Color      uint
	Default    uint
	Position_x uint
	Position_y uint
	// Car_type      string `gorm:"type:varchar(255);"`
	// Sample_status string `gorm:"type:varchar(255);"`
	// Create_time   string `gorm:"type:varchar(255);"`
	// Update_time   string `gorm:"type:varchar(255);"`
	// Charge_area   uint
	// Trans_area    uint
	// Status        string `gorm:"type:varchar(255);"`
	// Sample_name   string `gorm:"type:varchar(255);"`
	// Sample_type   string `gorm:"type:varchar(255);"`
	// Sample_id     uint
	// Sample_code   uint
	// Power         uint
	// Target        string `gorm:"type:varchar(255);"`
	// Position      string `gorm:"type:varchar(255);"`
	// Valid         uint
	// Username      string    `gorm:"type:varchar(255);"`
	// Controller_id string    `gorm:"type:varchar(255);"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}
