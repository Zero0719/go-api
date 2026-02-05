// Package response 响应处理工具
package response

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// JSON 响应 200 和 JSON 数据
func JSON(c *gin.Context, data interface{}) {
    c.JSON(http.StatusOK, data)
}

// Success 响应 200 和预设『操作成功！』的 JSON 数据
// 执行某个『没有具体返回数据』的『变更』操作成功后调用，例如删除、修改密码、修改手机号
func Success(c *gin.Context, data interface{}, message ...string) {
    JSON(c, gin.H{
        "code": 0,
        "message": defaultMessage("操作成功！", message...),
        "data": data,
    })
}

func Error(c *gin.Context, err error, code ...int) {
    errCode := 1
    if len(code) > 0 {
        errCode = code[0]
    } else {
        errCode = 1
    }
    JSON(c, gin.H{
        "code": errCode,
        "message": defaultMessage(err.Error()),
        "data": []interface{}{},
    })
}


// defaultMessage 内用的辅助函数，用以支持默认参数默认值
// Go 不支持参数默认值，只能使用多变参数来实现类似效果
func defaultMessage(defaultMsg string, msg ...string) (message string) {
    if len(msg) > 0 {
        message = msg[0]
    } else {
        message = defaultMsg
    }
    return
}