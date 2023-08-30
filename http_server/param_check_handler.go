// Package http_server @Author:冯铁城 [17615007230@163.com] 2023-06-21 16:52:02
package http_server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	chineseTrans "github.com/go-playground/validator/v10/translations/zh"
	"github.com/gorilla/mux"
)

// 校验路径参数handler
func checkPathParamHandler(w http.ResponseWriter, req *http.Request) {

	//1.获取路径参数
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	log.Printf("request path param is [%v]\n", id)

	//2.校验路径参数
	checkPathParam(w, id, "id", "gt=0,lte=1000")
}

// 校验URL参数handler
func checkUrlParamHandler(w http.ResponseWriter, req *http.Request) {

	//1.获取所有URL参数
	query := req.URL.Query()
	for key := range query {
		log.Printf("request url param %v is %v\n", key, query[key])
	}

	//2.封装为结构体
	id, _ := strconv.Atoi(query.Get("id"))
	age, _ := strconv.Atoi(query.Get("age"))
	student := &Student{Id: id, Name: query.Get("name"), Age: age}

	//3.校验结构体参数
	checkStructParam(w, student)
}

// 校验Body Json参数Handler
func checkBodyJsonParamHandler(w http.ResponseWriter, req *http.Request) {

	//1.读取body参数
	body, err := io.ReadAll(req.Body)
	if nil != err {
		log.Printf("get body error->%v", err.Error())
		return
	}

	//2.解析为JSON
	var student *Student
	if err = json.Unmarshal(body, &student); nil != err {
		log.Printf("parse json error->%v", err.Error())
		return
	}

	//3.打印参数
	log.Printf("request body param is %v\n", student)

	//4.校验结构体参数
	checkStructParam(w, student)
}

// 校验路径参数
func checkPathParam(w http.ResponseWriter, param int, fieldName string, tag string) {

	//1.初始化校验器
	validate, trans, err := initValidate()
	if err != nil {
		log.Printf("register translate error->[%v]", err.Error())
	}

	//2.校验id参数
	if errs := validate.Var(param, tag); errs != nil {

		//3.获取异常
		translateErrs := errs.(validator.ValidationErrors)[:1].Translate(trans)

		//4.格式化异常信息
		var errorMessage string
		for key := range translateErrs {
			errorMessage += fmt.Sprintf("[%v]%v\n", fieldName, translateErrs[key])
		}

		//5.日志打印
		log.Println(errorMessage)

		//6.响应
		http.Error(w, errorMessage, http.StatusBadRequest)
	}
}

// 校验结构体参数
func checkStructParam(w http.ResponseWriter, s interface{}) {

	//1.初始化校验器
	validate, trans, err := initValidate()
	if err != nil {
		log.Printf("register translate error->[%v]", err.Error())
	}

	//2.验证结构体
	if errs := validate.Struct(s); errs != nil {

		//3.获取异常
		validationErrors := errs.(validator.ValidationErrors)

		//4.定义异常描述切片
		errorMessages := make([]string, 0)

		//5.循环异常
		for _, e := range validationErrors {

			//6.翻译异常
			translateErrorMessage := e.Translate(trans)

			//7.存入集合
			errorMessages = append(errorMessages, translateErrorMessage)
		}

		//8.日志打印
		log.Println(errorMessages)

		//9.响应
		http.Error(w, strings.Join(errorMessages, "\n"), http.StatusBadRequest)
	}
}

// 初始化校验器
func initValidate() (*validator.Validate, ut.Translator, error) {

	//1.初始化校验器
	validate := validator.New()

	//2.初始化中文翻译器
	trans, _ := ut.New(zh.New()).GetTranslator("zh")

	//3.注册中文翻译器到校验器
	err := chineseTrans.RegisterDefaultTranslations(validate, trans)

	//4.返回
	return validate, trans, err
}
