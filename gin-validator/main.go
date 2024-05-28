package main

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type DeserialBinder interface {
	DeserializeBind(ctx *gin.Context) EmsError
}

type LoginParam struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password"`
}

func (l *LoginParam) DeserializeBind(ctx *gin.Context) EmsError {
	err := ctx.ShouldBindWith(l, binding.JSON)
	if err != nil {
		if validateErrors, ok := err.(validator.ValidationErrors); ok {
			switch validateErrors[0].Namespace() {
			case "LoginParam.Username":
				return New(INVALID_PARAM, "The login name is missing.")
			case "LoginParam.Password":
				return New(INVALID_PARAM, "The login password is missing.")
			}
		}
		return New(INVALID_PARAM, err.Error())
	}
	return nil
}

type RequestParams interface {
	LoginParam | LogoutParam
}

type LogoutParam struct {
	Token string `json:"token" binding:"required"`
}

func main() {
	g := gin.New()
	g.POST("/login", GetHandlerFunc(LoginHandlerFunc))
	_ = g.Run(":8080")
}

func GetHandlerFunc(handlerFunc interface{}) func(*gin.Context) {
	return func(ctx *gin.Context) {
		err := proxyHandlerFunc(ctx, handlerFunc)
		if err != nil {
			statusCode := http.StatusOK
			switch err.Code() {
			case INVALID_PARAM:
				statusCode = http.StatusBadRequest
			case NO_RECORD:
				statusCode = http.StatusNotFound
			case INTERNAL_ERROR:
				statusCode = http.StatusInternalServerError
			}
			ctx.JSON(statusCode, &gin.H{"code": err.Code(), "msg": err.Error()})
			ctx.Abort()
		}
		ctx.Next()
	}
}

func LoginHandlerFunc(ctx *gin.Context, param *LoginParam) EmsError {
	ctx.JSON(http.StatusOK, &gin.H{"code": 0, "msg": "success"})
	return nil
}

func proxyHandlerFunc(ctx *gin.Context, handlerFunc interface{}) EmsError {
	funcType := reflect.TypeOf(handlerFunc)
	if funcType.Kind() != reflect.Func {
		panic("the route handlerFunc must be a function")
	}
	funcValue := reflect.ValueOf(handlerFunc)
	typeParam := funcType.In(1).Elem()
	val := reflect.New(typeParam).Interface()
	deser, ok := val.(DeserialBinder)
	if !ok {
		var err error
		if ctx.Request.Method == http.MethodGet {
			err = ctx.ShouldBindQuery(val)
		} else {
			err = ctx.ShouldBind(val)
		}
		if err != nil {
			return New(INVALID_PARAM, err.Error())
		}
	} else {
		err := deser.DeserializeBind(ctx)
		if err != nil {
			return New(INVALID_PARAM, err.Error())
		}
		val = reflect.ValueOf(deser).Interface()
	}
	ret := funcValue.Call(valOf(ctx, val))[0].Interface()
	if ret != nil {
		return ret.(EmsError)
	}
	return nil
}

func valOf(i ...interface{}) []reflect.Value {
	var rt []reflect.Value
	for _, i2 := range i {
		rt = append(rt, reflect.ValueOf(i2))
	}
	return rt
}
