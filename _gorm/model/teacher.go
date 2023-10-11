// Package model @Author:冯铁城 [17615007230@163.com] 2023-10-11 09:00:53
package model

import "time"

// Teacher 教师模型
type Teacher struct {
	ID         uint      `gorm:"type:int(8);primaryKey;autoIncrement;comment:主键ID"`
	Name       string    `gorm:"type:varchar(64);not null;comment:名称"`
	Email      string    `gorm:"type:varchar(64);not null;unique;comment:邮箱"`
	Age        uint      `gorm:"type:int(8);not null;comment:年龄"`
	Subject    string    `gorm:"type:varchar(64);not null;comment:教学科目"`
	CreateTime time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdateTime time.Time `gorm:"type:datetime ON UPDATE CURRENT_TIMESTAMP;default:CURRENT_TIMESTAMP;comment:更新时间"`
}
