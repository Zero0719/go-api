package requests

import (

	"github.com/Zero0719/go-api/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	opts := govalidator.Options{
		Data: data,
		Rules: rules,
		Messages: messages,
		TagIdentifier: "valid",
	}

	return govalidator.New(opts).ValidateStruct()
}

type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		response.BadRequest(c, err, "请求解析错误，请确认请求格式是否正确。")
		return false
	}

	errs := handler(obj, c)
	if len(errs) > 0 {
		response.ValidationError(c, errs)
		return false
	}

	return true
}