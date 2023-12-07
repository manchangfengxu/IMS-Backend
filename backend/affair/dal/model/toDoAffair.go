package model

import (
	"gorm.io/gorm"
)
//TODO:time,taq
type ToDoAffair struct {
	gorm.Model
	Title   string `gorm:"type:char;size:20;comment:确定工作性质;not null"`
	State   bool   `gorm:"type:bool;default:false;commernt:事务是否完成"`
	Content string `gorm:"type:string;comment:事务详细内容;not null"`
}

func (ToDoAffair) TableName() string {
	return "ims_TODO_toDoAffair"
}
