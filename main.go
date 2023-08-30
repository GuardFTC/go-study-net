// @Author:冯铁城 [17615007230@163.com] 2023-06-21 15:03:33
package main

import "go-study-net/http_request"

func main() {

	//1.测试HTTP服务器
	testHttpServer()

	//2.测试gin
	testGin()

	//3.测试HTTP请求
	testHttpRequest()
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
	http_request.GetRequest()

	//2.post请求
	http_request.PostRequest()

	//3.put请求
	http_request.PutRequest()

	//4.put请求
	http_request.PatchRequest()

	//5.delete请求
	http_request.DeleteRequest()

	//6.上传本地文件
	http_request.UploadFileFromLocal()
}
