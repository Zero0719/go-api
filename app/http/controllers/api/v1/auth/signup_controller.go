package auth

import (

	v1 "github.com/Zero0719/go-api/app/http/controllers/api/v1"
	"github.com/Zero0719/go-api/app/models/user"
	"github.com/Zero0719/go-api/app/requests"
	"github.com/Zero0719/go-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type SignupController struct {
	v1.BaseAPIController
}

func (sc *SignupController) IsEmailExist(c *gin.Context) {
	// 初始化请求对象
    request := requests.SignupEmailExistRequest{}

    if ok := requests.Validate(c, &request, requests.ValidateSignupEmailExist); !ok {
        return
    }
    
    //  检查数据库并返回响应
    response.JSON(c, gin.H{
        "exist": user.IsEmailExist(request.Email),
    })
}