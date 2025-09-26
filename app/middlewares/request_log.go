package middlewares

import (
	"bytes"
	"go-api/app/utils"
	"io"
	"mime"
	"strings"

	"github.com/gin-gonic/gin"
)

// responseBodyWriter 用于捕获响应内容的写入器
type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write 重写Write方法，同时写入原始响应和缓冲区
func (w responseBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// isFileResponse 判断是否为文件响应
func isFileResponse(contentType string) bool {
	if contentType == "" {
		return false
	}

	// 解析Content-Type
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return false
	}

	// 常见的文件类型
	fileTypes := []string{
		"application/pdf",
		"application/zip",
		"application/octet-stream",
		"image/",
		"video/",
		"audio/",
		"text/csv",
		"application/vnd.ms-excel",
		"application/vnd.openxmlformats-officedocument",
	}

	for _, fileType := range fileTypes {
		if strings.HasPrefix(mediaType, fileType) {
			return true
		}
	}

	return false
}

func RequestLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 记录请求相关信息
		ip := ctx.ClientIP()
		method := ctx.Request.Method
		path := ctx.Request.URL.Path
		query := ctx.Request.URL.RawQuery

		// 获取请求体内容
		var requestBody string
		if ctx.Request.Body != nil {
			bodyBytes, err := io.ReadAll(ctx.Request.Body)
			if err == nil {
				requestBody = string(bodyBytes)
				// 重新设置请求体，因为ReadAll会消耗掉body
				ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}
		}

		// 创建响应写入器来捕获响应内容
		responseWriter := &responseBodyWriter{
			ResponseWriter: ctx.Writer,
			body:           &bytes.Buffer{},
		}
		ctx.Writer = responseWriter

		ctx.Next()

		// 获取响应内容
		responseBody := responseWriter.body.String()
		contentType := ctx.Writer.Header().Get("Content-Type")

		// 判断是否为文件响应
		var responseType string
		if isFileResponse(contentType) {
			responseType = "FILE"
			responseBody = "[文件内容]"
		} else {
			responseType = "JSON/TEXT"
		}

		// 记录日志到专门的请求日志文件
		utils.RequestLogger.Info().
			Str("ip", ip).
			Str("method", method).
			Str("path", path).
			Str("query", query).
			Str("request_body", requestBody).
			Str("response_type", responseType).
			Str("response_body", responseBody).
			Int("status", ctx.Writer.Status()).
			Msg("HTTP Request")
	}
}
