// Package _gorm @Author:冯铁城 [17615007230@163.com] 2023-10-11 09:46:46
package _gorm

import (
	"fmt"
	"log"

	"go-study-net/_gorm/model"
	"gorm.io/gorm"
)

// SimpleSelect 初级查询
func SimpleSelect(db *gorm.DB) {

	//1.查询全部学生信息,不携带任何条件
	var students []*model.Student
	if result := db.Find(&students); result.Error != nil {
		panic(fmt.Sprintf("select data error:[%v]", result.Error))
	}
	fmt.Println("\nget all data:")
	for _, student := range students {
		log.Print(fmt.Sprintf("student=[%v]", student))
	}

	//2.获取ID正序排列的第一个学生
	var firstStudent *model.Student
	if result := db.First(&firstStudent); result.Error != nil {
		panic(fmt.Sprintf("select data error:[%v]", result.Error))
	}
	fmt.Println("\nget order by id asc first data:")
	log.Print(fmt.Sprintf("firstStudent=[%v]", firstStudent))

	//3.获取ID倒序排列的第一个学生
	var lastStudent *model.Student
	if result := db.Last(&lastStudent); result.Error != nil {
		panic(fmt.Sprintf("select data error:[%v]", result.Error))
	}
	fmt.Println("\nget order by id desc first data:")
	log.Print(fmt.Sprintf("lastStudent=[%v]", lastStudent))

	//4.根据主键ID查询学生
	var studentById *model.Student
	if result := db.First(&studentById, lastStudent.ID); result.Error != nil {
		panic(fmt.Sprintf("select data error:[%v]", result.Error))
	}
	fmt.Println("\nget data by id:")
	log.Print(fmt.Sprintf("studentById=[%v]", studentById))

	//5.根据主键ID查询学生
	studentByIdAndHasId := &model.Student{ID: lastStudent.ID}
	if result := db.First(&studentByIdAndHasId); result.Error != nil {
		panic(fmt.Sprintf("select data error:[%v]", result.Error))
	}
	fmt.Println("\nget data by id:")
	log.Print(fmt.Sprintf("studentByIdAndHasId=[%v]", studentByIdAndHasId))

	//6.根据主键ID批量查询学生
	var studentsByIds []*model.Student
	if result := db.Find(&studentsByIds, []uint{firstStudent.ID, lastStudent.ID}); result.Error != nil {
		panic(fmt.Sprintf("select data error:[%v]", result.Error))
	}
	fmt.Println("\nget batch data by id:")
	for _, studentsById := range studentsByIds {
		log.Print(fmt.Sprintf("studentsById=[%v]", studentsById))
	}
}

// WhereSelect 条件查询
func WhereSelect(db *gorm.DB) {

	//1.等值查询，获取符合条件的第一个对象
	var studentByNameFirst *model.Student
	if result := db.Where("name = ?", "马冬梅").First(&studentByNameFirst); result.Error != nil {
		panic(fmt.Sprintf("select data error:[%v]", result.Error))
	}
	fmt.Println("\nget where equal first data:")
	log.Print(fmt.Sprintf("studentByNameFirst=[%v]", studentByNameFirst))

	//2.等值查询，获取符合条件的全部对象
	var studentsByGrade []*model.Student
	if result := db.Where("grade = ?", 3).Find(&studentsByGrade); result.Error != nil {
		panic(fmt.Sprintf("select data error:[%v]", result.Error))
	}
	fmt.Println("\nget where equal all data:")
	for _, studentByGrade := range studentsByGrade {
		log.Print(fmt.Sprintf("studentByGrade=[%v]", studentByGrade))
	}

	//3.不等值查询，获取符合条件的全部对象
	var studentsByGradeNotEqual []*model.Student
	if result := db.Where("grade <> ?", 3).Find(&studentsByGradeNotEqual); result.Error != nil {
		panic(fmt.Sprintf("select data error:[%v]", result.Error))
	}
	fmt.Println("\nget where not equal all data:")
	for _, studentByGradeNotEqual := range studentsByGradeNotEqual {
		log.Print(fmt.Sprintf("studentByGradeNotEqual=[%v]", studentByGradeNotEqual))
	}

	//4.in查询，获取符合条件的全部对象
	var studentsIn []*model.Student
	if result := db.Where("age in ?", []uint{17, 19}).Find(&studentsIn); result.Error != nil {
		panic(fmt.Sprintf("select data error:[%v]", result.Error))
	}
	fmt.Println("\nget where in all data:")
	for _, studentIn := range studentsIn {
		log.Print(fmt.Sprintf("studentIn=[%v]", studentIn))
	}

	//5.like查询，获取符合条件的全部对象
	var studentsLike []*model.Student
	if result := db.Where("email like ?", "%zhang%").Find(&studentsLike); result.Error != nil {
		panic(fmt.Sprintf("select data error:[%v]", result.Error))
	}
	fmt.Println("\nget where like all data:")
	for _, studentLike := range studentsLike {
		log.Print(fmt.Sprintf("studentLike=[%v]", studentLike))
	}

	//6.大于/大于等于查询
	var studentsGqOrGe []*model.Student
	if result := db.Where("age > ?", 18).Find(&studentsGqOrGe); result.Error != nil {
		panic(fmt.Sprintf("select data error:[%v]", result.Error))
	}
	fmt.Println("\nget where gq/ge all data:")
	for _, studentGqOrGe := range studentsGqOrGe {
		log.Print(fmt.Sprintf("studentGqOrGe=[%v]", studentGqOrGe))
	}

	//7.小于/小于等于查询
	var studentsLqOrLe []*model.Student
	if result := db.Where("age <= ?", 18).Find(&studentsLqOrLe); result.Error != nil {
		panic(fmt.Sprintf("select data error:[%v]", result.Error))
	}
	fmt.Println("\nget where lq/le all data:")
	for _, studentLqOrLe := range studentsLqOrLe {
		log.Print(fmt.Sprintf("studentLqOrLe=[%v]", studentLqOrLe))
	}

	//8.AND查询
	var studentsAnd []*model.Student
	if result := db.Where("age >= ? and name = ?", 18, "陈凯").Find(&studentsAnd); result.Error != nil {
		panic(fmt.Sprintf("select data error:[%v]", result.Error))
	}
	fmt.Println("\nget where and all data:")
	for _, studentAnd := range studentsAnd {
		log.Print(fmt.Sprintf("studentAnd=[%v]", studentAnd))
	}

	//9.OR查询
	var studentsOr []*model.Student
	if result := db.Where("age > ?", 20).Or("name = ?", "马冬梅").Find(&studentsOr); result.Error != nil {
		panic(fmt.Sprintf("select data error:[%v]", result.Error))
	}
	fmt.Println("\nget where or all data:")
	for _, studentOr := range studentsOr {
		log.Print(fmt.Sprintf("studentOr=[%v]", studentOr))
	}

	//10.结构映射查询
	param := &model.Student{Age: 18}
	var students []*model.Student
	if result := db.Where(param).Find(&students); result.Error != nil {
		panic(fmt.Sprintf("select data error:[%v]", result.Error))
	}
	fmt.Println("\nget where all data:")
	for _, student := range students {
		log.Print(fmt.Sprintf("student=[%v]", student))
	}

	//11.内联查询
	var inlineStudents []*model.Student
	if result := db.Find(&inlineStudents, "age = ?", 18); result.Error != nil {
		panic(fmt.Sprintf("select data error:[%v]", result.Error))
	}
	fmt.Println("\nget where inline all data:")
	for _, inlineStudent := range inlineStudents {
		log.Print(fmt.Sprintf("inlineStudent=[%v]", inlineStudent))
	}

	//12.NOT查询
	var notStudents []*model.Student
	if result := db.Not("age in ?", []uint{17, 19}).Find(&notStudents); result.Error != nil {
		panic(fmt.Sprintf("select data error:[%v]", result.Error))
	}
	fmt.Println("\nget where not all data:")
	for _, notStudent := range notStudents {
		log.Print(fmt.Sprintf("notStudent=[%v]", notStudent))
	}

	//13.查询指定字段
	var selectStudents []*model.Student
	if result := db.Select("id", "name").Where("age = ?", 17).Find(&selectStudents); result.Error != nil {
		panic(fmt.Sprintf("select data error:[%v]", result.Error))
	}
	fmt.Println("\nget where select all data:")
	for _, selectStudent := range selectStudents {
		log.Print(fmt.Sprintf("selectStudent=[%v]", selectStudent))
	}

	//14.查询数量
	var count int64
	if result := db.Model(&model.Student{}).Count(&count); result.Error != nil {
		panic(fmt.Sprintf("select data error:[%v]", result.Error))
	}
	fmt.Println("\nget count:")
	log.Print(fmt.Sprintf("count=[%v]", count))
}

// group查询结果
type groupResult struct {
	Age   int
	Total int
}

// SpecialSelect 特殊查询
func SpecialSelect(db *gorm.DB) {

	//1.排序，获取全部数据，根据年龄倒序，id正序
	var orderStudents []*model.Student
	if result := db.Order("age desc,id asc").Find(&orderStudents); result.Error != nil {
		panic(fmt.Sprintf("select data error:[%v]", result.Error))
	}
	fmt.Println("\nget order data:")
	for _, orderStudent := range orderStudents {
		log.Print(fmt.Sprintf("orderStudent=[%v]", orderStudent))
	}

	//2.偏移量查询
	var limitStudents []*model.Student
	if result := db.Order("age desc,id asc").Offset(2).Limit(2).Find(&limitStudents); result.Error != nil {
		panic(fmt.Sprintf("select data error:[%v]", result.Error))
	}
	fmt.Println("\nget offset and limit data:")
	for _, limitStudent := range limitStudents {
		log.Print(fmt.Sprintf("limitStudent=[%v]", limitStudent))
	}

	//3.group by and having查询
	var groupResult *groupResult
	if result := db.Model(&model.Student{}).Select("age,count(1) as total").Group("age").Having("total > 1").Find(&groupResult); result.Error != nil {
		panic(fmt.Sprintf("select data error:[%v]", result.Error))
	}
	fmt.Println("\nget group by data:")
	log.Print(fmt.Sprintf("result=[%v]", groupResult))

	//4.join查询
	var teacher *model.Teacher
	db.Where("name = ?", "王老师").Find(&teacher)

	var joinStudents []*model.Student
	db.Model(&model.Student{}).Select("students.id,students.name,students.age,students.grade,students.class").
		Joins("inner join study.teacher_and_students tas on students.id = tas.student_id").
		Where("tas.teacher_id = ?", teacher.ID).
		Order("students.id desc").
		Scan(&joinStudents)
	fmt.Println("\nget join data:")
	for _, joinStudent := range joinStudents {
		log.Print(fmt.Sprintf("joinStudent=[%v]", joinStudent))
	}
}
