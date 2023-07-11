package common

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"io"
	"mime/multipart"
	"reflect"
)

type BizFunc[Req, Resp any] func(ctx context.Context, t Req) (resp Resp, err error)

func HandlerFuncWrapper[Req, Resp any](bizFunc BizFunc[Req, Resp]) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req Req
		if reflect.TypeOf(req).Kind() == reflect.Ptr {
			v := reflect.ValueOf(&req).Elem()
			v.Set(reflect.New(v.Type().Elem()))
		}
		if err := c.BindAndValidate(&req); err != nil {
			c.JSON(400, map[string]any{
				"code":    400,
				"message": err.Error(),
			})
			return
		}
		// 自定义标签解析
		if err := injectWrapperTagValue(ctx, c, req); err != nil {
			c.JSON(400, map[string]any{
				"code":    400,
				"message": err.Error(),
			})
			return
		}

		resp, err := bizFunc(ctx, req)
		if err != nil {
			c.JSON(500, map[string]any{
				"code":    500,
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, map[string]any{
			"code":    0,
			"message": "ok",
			"data":    resp,
		})
	}
}

type UserRegisterParam struct {
	Username string        `json:"username" vd:"len($)>0"`
	Password string        `json:"password" vd:"len($)>0"`
	Avatar   *FormFileInfo `wrapper:"photo"`
}

func UserRegister(ctx context.Context, params *UserRegisterParam) (map[string]any, error) {
	// 处理业务逻辑

	return map[string]any{
		"user_id": "12345",
	}, nil
}

type FormFileInfo struct {
	Raw  []byte // 文件二进制
	Name string // 原始文件名称
	Size int64  // 文件大小
}

func (f *FormFileInfo) Reader() io.Reader {
	return bytes.NewReader(f.Raw)
}

func injectWrapperTagValue(ctx context.Context, c *app.RequestContext, v any) error {
	kind := reflect.TypeOf(v).Kind()
	if kind != reflect.Ptr {
		return errors.New("inject typeVal must be a ptr kind")
	}
	// 传入的v可能是一个多级指针，需要转换成一级指针
	v = getFirstRankPtr(v)
	if reflect.ValueOf(v).IsNil() {
		// 空指针直接返回
		return nil
	}
	typeVal := reflect.TypeOf(v).Elem() // 获取类型
	vals := reflect.ValueOf(v).Elem()   // 获取值
	for i := 0; i < typeVal.NumField(); i++ {
		tag := typeVal.Field(i).Tag
		if tagVal, ok := tag.Lookup("wrapper"); ok {
			// 存在tag
			formFileInfo, err := getUploadFileInfoWithContext(ctx, c, tagVal)
			if err != nil {
				return err
			}
			if formFileInfo == nil {
				return errors.New("parse form file failed, form file empty")
			}
			// 注入值
			val := vals.Field(i)
			if val.Kind() == reflect.Ptr {
				// 指针类型的属性
				if reflect.TypeOf(val.Interface()).Elem().Name() != "FormFileInfo" {
					return fmt.Errorf("current field [%v] type not support `wrapper` tag", reflect.TypeOf(val.Interface()).Elem().Name())
				}
				val.Set(reflect.ValueOf(formFileInfo))
			} else {
				if val.Type().Name() != "FormFileInfo" {
					return fmt.Errorf("current field [%v] type not support `wrapper` tag", val.Type().Name())
				}
				val.Set(reflect.ValueOf(*formFileInfo))
			}
		}
	}
	return nil
}

// 多级指针转换成一级指针
func getFirstRankPtr(v any) any {
	temp := v
	for reflect.TypeOf(temp).Kind() == reflect.Ptr {
		v = temp
		if reflect.ValueOf(temp).IsNil() {
			// 空指针直接返回，不再取值
			return temp
		}
		temp = reflect.ValueOf(temp).Elem().Interface()
	}
	return v
}

func getUploadFileInfoWithContext(ctx context.Context, c *app.RequestContext, filename string) (*FormFileInfo, error) {
	raw, err := c.FormFile(filename)
	if err != nil {
		return nil, nil
	}
	fileInfo, err := getUploadFileInfo(raw)
	if err != nil {
		return nil, err
	}
	return fileInfo, nil
}

func getUploadFileInfo(fileHeader *multipart.FileHeader) (*FormFileInfo, error) {
	openFile, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer openFile.Close()

	data := make([]byte, fileHeader.Size+10)
	_, err = openFile.Read(data) //读取传入文件的内容
	if err != nil {
		return nil, err
	}
	return &FormFileInfo{
		Raw:  data,
		Name: fileHeader.Filename,
		Size: fileHeader.Size,
	}, nil
}
