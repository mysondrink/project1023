// 用户数据传输对象定义
package dto

import "assemblyline/project1023/model"

type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

type TrailDto struct {
	Start_time  string `json:"start_time"`
	End_time    string `json:"end_time"`
	Trail_name  string `json:"trail_name"`
	Module_type string `json:"module_type"`
	Trail_type  string `json:"trail_type"`
	Trail_id    uint   `json:"trail_id"`
	Car_id      uint   `json:"car_id"`
	Car_type    string `json:"car_type"`
	Status      string `json:"status"`
	Create_time string `json:"create_time"`
	Update_time string `json:"update_time"`
	Charge_area uint   `json:"charge_area"`
	Tran_area   uint   `json:"tran_area"`
}

func ToUserDto(user model.UserTable) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}

func ToTrailDto(trail model.Trail) TrailDto {
	return TrailDto{
		Start_time:  trail.Start_time,
		End_time:    trail.End_time,
		Trail_name:  trail.Trail_name,
		Module_type: trail.Module_type,
		Trail_type:  trail.Trail_type,
		Trail_id:    trail.Trail_id,
		Car_id:      trail.Car_id,
		Car_type:    trail.Car_type,
		Status:      trail.Status,
		Create_time: trail.Create_time,
		Update_time: trail.Update_time,
		Charge_area: trail.Charge_area,
		Tran_area:   trail.Tran_area,
	}
}
