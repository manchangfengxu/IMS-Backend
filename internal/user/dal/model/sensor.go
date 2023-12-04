package model

import (
	"gorm.io/gorm"
	MqCo "ims-server/pkg/Mqtt"
)

type Sensor struct {
	gorm.Model
	Type       MqCo.SensorType `gorm:"type:int;size:15;comment:传感器类型;not null"`
	SensorDate interface{}     `gorm:"type:int;size:10;comment:传感器数据;unique"`
	TerminalID string          `gorm:"type:string;comment:传感器ID;not null"`
}

func (Sensor) TableName() string {
	return "ims_aaa_sensor"
}
