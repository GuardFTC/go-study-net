// Package _gorm @Author:冯铁城 [17615007230@163.com] 2023-10-11 14:32:40
package _gorm

import (
	"time"

	"go-study-net/_gorm/model"
	"gorm.io/gorm"
)

// Update 更新数据
func Update(db *gorm.DB) {

	//1.获取夏洛
	var student *model.Student
	db.Where("name = ?", "夏洛").First(&student)

	//2.单个全量更新
	student.Age = 66
	student.UpdateTime = time.Now().Local()
	db.Save(&student)

	//3.单个更新单列
	db.Model(&student).Update("email", "hhhhhh@xltfn.cn")

	//4.批量更新单列
	db.Model(&model.Student{}).Where("age > ?", 18).Update("create_time", time.Now().Local())

	//5.批量更新多列
	db.Model(&model.Student{}).Where("age <= ?", 18).Updates(&model.Student{UpdateTime: time.Now().Local(), Class: 11})

	//6.不带条件的全部更新
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&model.Teacher{}).Update("create_time", time.Now().Local())
}
