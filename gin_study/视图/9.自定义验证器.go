package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
)

// 返回结构体的msg参数
func _GetValidMsg(err error, obj any) string {
	getObj := reflect.TypeOf(obj)
	//将err接口断言为具体类型
	if errs, ok := err.(validator.ValidationErrors); ok {
		//断言成功
		for _, e := range errs {
			//循环每一个错误信息
			//根据报错字段名，获取结构体的具体字段
			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
				msg := f.Tag.Get("msg")
				return msg
			}
		}
	}
	return ""
}

type User struct {
	Name string `json:"name" binding:"required,sign" msg:"用户名校验失败"`
	Age  int    `json:"age" binding:"required" msg:"请输入年龄"`
}

func signValid(fl validator.FieldLevel) bool {
	var nameList []string = []string{"杰", "123"}
	name := fl.Field().Interface().(string)
	for _, str := range nameList {
		if name == str {
			return false
		}
	}
	return true
}

func main() {
	router := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("sign", signValid)
	}

	router.POST("/", func(c *gin.Context) {

		var user User
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(200, gin.H{"msg": _GetValidMsg(err, &user)})
			return
		}
		c.JSON(200, gin.H{"data": user})
		return
	})
	router.Run(":80")
}
