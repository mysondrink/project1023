// 用户数据传输对象定义
package dto

import "assemblyline/project1023/model"

type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

type TrailDto struct {
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
	Create_time     string `json:"create_time"`
	Update_time     string `json:"update_time"`
	Charge_area     uint   `json:"charge_area"`
	Tran_area       uint   `json:"tran_area"`
	Status          string `json:"status"`
}

func ToUserDto(user model.UserTable) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}

func ToTrailDto(trail model.Trail) TrailDto {
	return TrailDto{
		Controller_type: trail.Controller_type,
		Controller_id:   trail.Controller_id,
		Start_time:      trail.Start_time,
		End_time:        trail.End_time,
		Trail_name:      trail.Trail_name,
		Trail_id:        trail.Trail_id,
		Trail_type:      trail.Trail_type,
		Module_name:     trail.Module_name,
		Module_id:       trail.Module_id,
		Module_type:     trail.Module_type,
		Car_id:          trail.Car_id,
		Create_time:     trail.Create_time,
		Update_time:     trail.Update_time,
		Charge_area:     trail.Charge_area,
		Tran_area:       trail.Tran_area,
		Status:          trail.Status,
	}
}
