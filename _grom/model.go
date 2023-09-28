// Package _grom  @Author:冯铁城 [17615007230@163.com] 2023-09-12 09:20:03
package _grom

import "time"

// Student 学生模型
type Student struct {
	ID         uint      `gorm:"type:int(8);primaryKey;autoIncrement;comment:主键ID"`
	Name       string    `gorm:"type:varchar(64);not null;comment:名称"`
	Email      string    `gorm:"type:varchar(64);not null;unique;comment:邮箱"`
	Age        uint      `gorm:"type:int(8);not null;comment:年龄"`
	Grade      uint      `gorm:"type:int(8);not null;comment:年级编号"`
	Class      uint      `gorm:"type:int(8);not null;comment:班级编号"`
	CreateTime time.Time `gorm:"type:datetime;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdateTime time.Time `gorm:"type:datetime ON UPDATE CURRENT_TIMESTAMP;default:CURRENT_TIMESTAMP;comment:更新时间"`
}

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

// TeacherAndStudent 学生老师关联表模型
type TeacherAndStudent struct {
	ID        uint `gorm:"type:int(8);primaryKey;autoIncrement;comment:主键ID"`
	TeacherId uint `gorm:"type:int(8);comment:教师ID"`
	StudentId uint `gorm:"type:int(8);comment:学生ID"`
}
