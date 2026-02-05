package requests

import (
	"errors"
	"fmt"
	"net/http"

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
		response.Error(c, errors.New("请求解析错误，请确认请求格式是否正确。"), http.StatusBadRequest)
		fmt.Println("test")
		return false
	}

	errs := handler(obj, c)
	if len(errs) > 0 {
		for _, err := range errs {
			response.Error(c, errors.New(err[0]), http.StatusUnprocessableEntity)
			return false
		}
	}

	return true
}