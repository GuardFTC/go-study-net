// @Author:冯铁城 [17615007230@163.com] 2023-06-21 15:03:33
package main

import (
	"fmt"

	"go-study-net/_gorm"
	"go-study-net/_gorm/model"
)

func main() {

	//1.测试HTTP服务器
	testHttpServer()

	//2.测试gin
	testGin()

	//3.测试HTTP请求
	testHttpRequest()

	//4.测试Gorm框架
	testGormDB()
}

// 测试HTTP服务器
func testHttpServer() {

	//1.hello world server
	//http_server.StartHelloWorldServer()

	//2.http server
	//http_server.StartHttpServer()
}

// 测试Gin框架
func testGin() {

	//1.hello world server
	//gin.StartHelloWorldServer()

	//2.gin server
	//gin.StartGinServer()
}

// 测试HTTP请求
func testHttpRequest() {

	//1.get请求
	//http_request.GetRequest()

	//2.post请求
	//http_request.PostRequest()

	//3.put请求
	//http_request.PutRequest()

	//4.put请求
	//http_request.PatchRequest()

	//5.delete请求
	//http_request.DeleteRequest()

	//6.上传本地文件
	//http_request.UploadFileFromLocal()
}

// 测试Gorm框架
func testGormDB() {

	//1.初始化数据库连接
	db := _gorm.InitDB()

	//2.创建学生数据库
	err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='学生表'").AutoMigrate(&model.Student{})
	if err != nil {
		panic(fmt.Sprintf("create table error:[%v]", err))
	}

	//3.创建老师数据库
	err = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='老师表'").AutoMigrate(&model.Teacher{})
	if err != nil {
		panic(fmt.Sprintf("create table error:[%v]", err))
	}

	//4.创建老师学生中间表数据库
	err = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='老师学生表'").AutoMigrate(&model.TeacherAndStudent{})
	if err != nil {
		panic(fmt.Sprintf("create table error:[%v]", err))
	}

	//5.创建数据
	_gorm.Create(db)

	//6.初级查询
	_gorm.SimpleSelect(db)

	//7.where查询
	_gorm.WhereSelect(db)

	//8.特殊查询
	_gorm.SpecialSelect(db)

	//9.更新数据
	_gorm.Update(db)

	//10.删除数据
	_gorm.Delete(db)
}
