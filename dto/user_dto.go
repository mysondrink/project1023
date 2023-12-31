// 用户数据传输对象定义
package dto

import (
	"assemblyline/project1023/model"
	"time"
)

type UserDto struct {
	Name      string    `json:"name"`
	Telephone string    `json:"telephone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ControllerDto struct {
	Controller_id   string `json:"controller_id"`
	Controller_type string `json:"controller_type"`
	Start_time      string `json:"start_time"`
	End_time        string `json:"end_time"`
	Trail_name      string `json:"trail_name"`
	Trail_id        uint   `json:"trail_id"`
	Trail_type      string `json:"trail_type"`
	Module_name     string `json:"module_name"`
	Module_id       uint   `json:"module_id"`
	Module_type     string `json:"module_type"`
	Car_id          uint   `json:"car_id"`
	// Create_time     string `json:"create_time"`
	// Update_time     string `json:"update_time"`
	Charge_area uint      `json:"charge_area"`
	Trans_area  uint      `json:"trans_area"`
	Status      string    `json:"status"`
	Tx_status   string    `json:"tx_status"`
	Rx_status   string    `json:"rx_status"`
	Rates       string    `json:"rates"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CarDto struct {
	Start_time string `json:"start_time"`
	End_time   string `json:"end_time"`
	Car_name   string `json:"car_name"`
	Car_id     uint   `json:"car_id"`
	Car_type   string `json:"car_type"`
	// Create_time   string `json:"create_time"`
	// Update_time   string `json:"update_time"`
	Charge_area   uint      `json:"charge_area"`
	Trans_area    uint      `json:"trans_area"`
	Status        string    `json:"status"`
	Sample_name   string    `json:"sample_name"`
	Sample_status string    `json:"sample_status"`
	Sample_type   string    `json:"sample_type"`
	Sample_id     uint      `json:"sample_id"`
	Sample_code   uint      `json:"sample_code"`
	Power         uint      `json:"power"`
	Target        string    `json:"target"`
	Position      string    `json:"position"`
	Valid         uint      `json:"valid"`
	Username      string    `json:"username"`
	Controller_id string    `json:"controller_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type ToyDto struct {
	Start_time string `json:"start_time"`
	End_time   string `json:"end_time"`
	// Car_name   string `json:"car_name"`
	Toy_id     uint `json:"toy_id"`
	Color      uint `json:"color"`
	Default    uint `json:"default"`
	Position_x uint `json:"position_x"`
	Position_y uint `json:"position_y"`
	// Car_type   string `json:"car_type"`
	// Create_time   string `json:"create_time"`
	// Update_time   string `json:"update_time"`
	// Charge_area   uint      `json:"charge_area"`
	// Trans_area    uint      `json:"trans_area"`
	// Status        string    `json:"status"`
	// Sample_name   string    `json:"sample_name"`
	// Sample_status string    `json:"sample_status"`
	// Sample_type   string    `json:"sample_type"`
	// Sample_id     uint      `json:"sample_id"`
	// Sample_code   uint      `json:"sample_code"`
	// Power         uint      `json:"power"`
	// Target        string    `json:"target"`
	// Position      string    `json:"position"`
	// Valid         uint      `json:"valid"`
	// Username      string    `json:"username"`
	// Controller_id string    `json:"controller_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToUserDto(user model.UserTable) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}

func ToControllerDto(controller []model.Controller) []ControllerDto {
	// for item := range controller {
	// }
	var result []ControllerDto
	for _, item := range controller {
		template := ControllerDto{
			Controller_type: item.Controller_type,
			Controller_id:   item.Controller_id,
			Start_time:      item.Start_time,
			End_time:        item.End_time,
			Trail_name:      item.Trail_name,
			Trail_id:        item.Trail_id,
			Trail_type:      item.Trail_type,
			Module_name:     item.Module_name,
			Module_id:       item.Module_id,
			Module_type:     item.Module_type,
			Car_id:          item.Car_id,
			CreatedAt:       item.CreatedAt,
			UpdatedAt:       item.UpdatedAt,
			Charge_area:     item.Charge_area,
			Trans_area:      item.Trans_area,
			Status:          item.Status,
			Tx_status:       item.Tx_status,
			Rx_status:       item.Rx_status,
			Rates:           item.Rates,
		}
		result = append(result, template)
	}
	return result
}

func ToCarDto(car []model.Car) []CarDto {
	// for item := range car {
	// }
	var result []CarDto
	for _, item := range car {
		template := CarDto{
			Start_time:    item.Start_time,
			End_time:      item.End_time,
			Car_name:      item.Car_name,
			Car_id:        item.Car_id,
			Car_type:      item.Car_type,
			Sample_status: item.Sample_status,
			CreatedAt:     item.CreatedAt,
			UpdatedAt:     item.UpdatedAt,
			Charge_area:   item.Charge_area,
			Trans_area:    item.Trans_area,
			Status:        item.Status,
			Sample_name:   item.Sample_name,
			Sample_type:   item.Sample_type,
			Sample_id:     item.Sample_id,
			Sample_code:   item.Sample_code,
			Power:         item.Power,
			Target:        item.Target,
			Position:      item.Position,
			Valid:         item.Valid,
			Username:      item.Username,
			Controller_id: item.Controller_id,
		}
		result = append(result, template)
	}
	return result
}

func ToToyDto(toy []model.Toy) []ToyDto {
	// for item := range car {
	// }
	var result []ToyDto
	for _, item := range toy {
		template := ToyDto{
			Start_time: item.Start_time,
			End_time:   item.End_time,
			Toy_id:     item.Toy_id,
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
			Position_x: item.Position_x,
			Position_y: item.Position_y,
			Default:    item.Default,
			Color:      item.Color,
		}
		result = append(result, template)
	}
	return result
}

/*
func ToColumnDto(car []model.Car) []CarDto {
	// for item := range car {
	// }
	var result []CarDto
	for _, item := range car {
		template := CarDto{
			Start_time:    item.Start_time,
			End_time:      item.End_time,
			Car_name:      item.Car_name,
			Car_id:        item.Car_id,
			Car_type:      item.Car_type,
			Sample_status: item.Sample_status,
			CreatedAt:     item.CreatedAt,
			UpdatedAt:     item.UpdatedAt,
			Charge_area:   item.Charge_area,
			Trans_area:    item.Trans_area,
			Status:        item.Status,
			Sample_name:   item.Sample_name,
			Sample_type:   item.Sample_type,
			Sample_id:     item.Sample_id,
			Sample_code:   item.Sample_code,
			Power:         item.Power,
			Target:        item.Target,
			Position:      item.Position,
			Valid:         item.Valid,
			Username:      item.Username,
			Controller_id: item.Controller_id,
		}
		result = append(result, template)
	}
	return result
}
*/
