// Package _gorm @Author:冯铁城 [17615007230@163.com] 2023-10-11 09:05:35
package _gorm

import (
	"fmt"

	"go-study-net/_gorm/model"
	"gorm.io/gorm"
)

// Create 创建数据
func Create(db *gorm.DB) {

	//1.单个添加学生对象
	student1 := &model.Student{Name: "夏洛", Email: "xialuo@xltfn.cn", Age: 18, Grade: 3, Class: 2}
	if result := db.Create(student1); result.Error != nil {
		panic(fmt.Sprintf("create data error:[%v]", result.Error))
	}

	//2.批量添加多个学生
	student2 := &model.Student{Name: "马冬梅", Email: "madongmei@xltfn.cn", Age: 18, Grade: 3, Class: 2}
	student3 := &model.Student{Name: "大傻春", Email: "dashachun@xltfn.cn", Age: 21, Grade: 3, Class: 2}
	student4 := &model.Student{Name: "袁华", Email: "yuanhua@xltfn.cn", Age: 17, Grade: 4, Class: 1}
	students := []*model.Student{student2, student3, student4}
	if result := db.Create(students); result.Error != nil {
		panic(fmt.Sprintf("create data error:[%v]", result.Error))
	}

	//3.批量添加学生，指定每次添加数量
	student5 := &model.Student{Name: "张扬", Email: "zhangyang@xltfn.cn", Age: 19, Grade: 3, Class: 2}
	student6 := &model.Student{Name: "秋雅", Email: "qiuya@xltfn.cn", Age: 18, Grade: 3, Class: 2}
	student7 := &model.Student{Name: "陈凯", Email: "chenkai@xltfn.cn", Age: 27, Grade: 8, Class: 4}
	students = []*model.Student{student5, student6, student7}
	if result := db.CreateInBatches(students, 2); result.Error != nil {
		panic(fmt.Sprintf("create data error:[%v]", result.Error))
	}

	//4.添加老师数据
	teacher1 := &model.Teacher{Name: "王老师", Email: "tearcher_wang@xltfn.cn", Age: 45, Subject: "语文"}
	teacher2 := &model.Teacher{Name: "校长", Email: "xiaozhang@xltfn.cn", Age: 56, Subject: "所有科目"}
	teachers := []*model.Teacher{teacher1, teacher2}
	if result := db.Create(teachers); result.Error != nil {
		panic(fmt.Sprintf("create data error:[%v]", result.Error))
	}

	//5.添加关联数据
	teacherAndStudents := []*model.TeacherAndStudent{
		{TeacherId: teacher1.ID, StudentId: student1.ID},
		{TeacherId: teacher1.ID, StudentId: student2.ID},
		{TeacherId: teacher1.ID, StudentId: student3.ID},
		{TeacherId: teacher1.ID, StudentId: student4.ID},
		{TeacherId: teacher1.ID, StudentId: student5.ID},
		{TeacherId: teacher1.ID, StudentId: student6.ID},
		{TeacherId: teacher1.ID, StudentId: student7.ID},
	}
	if result := db.Create(teacherAndStudents); result.Error != nil {
		panic(fmt.Sprintf("create data error:[%v]", result.Error))
	}
}
