# 中文验证器使用指南

## 概述

本项目集成了基于 `github.com/go-playground/validator/v10` 的中文验证器，支持中文错误提示和自定义验证规则。

## 功能特性

- ✅ 中文错误提示
- ✅ 自定义字段名翻译
- ✅ 支持所有标准验证规则
- ✅ 自定义验证规则翻译
- ✅ 与 Gin 框架完美集成

## 验证规则

### 常用验证规则

| 规则 | 说明 | 示例 |
|------|------|------|
| `required` | 必填字段 | `validate:"required"` |
| `min` | 最小长度/值 | `validate:"min=3"` |
| `max` | 最大长度/值 | `validate:"max=20"` |
| `len` | 固定长度 | `validate:"len=11"` |
| `email` | 邮箱格式 | `validate:"email"` |
| `gte` | 大于等于 | `validate:"gte=1"` |
| `lte` | 小于等于 | `validate:"lte=120"` |
| `gt` | 大于 | `validate:"gt=0"` |
| `lt` | 小于 | `validate:"lt=100"` |
| `oneof` | 枚举值 | `validate:"oneof=red green blue"` |
| `numeric` | 数字 | `validate:"numeric"` |
| `alpha` | 字母 | `validate:"alpha"` |
| `alphanum` | 字母数字 | `validate:"alphanum"` |
| `url` | URL格式 | `validate:"url"` |
| `uuid` | UUID格式 | `validate:"uuid"` |

## 使用方法

### 1. 定义结构体

```go
type UserRequest struct {
    Username string `json:"username" label:"用户名" validate:"required,min=3,max=20"`
    Email    string `json:"email" label:"邮箱" validate:"required,email"`
    Password string `json:"password" label:"密码" validate:"required,min=6,max=20"`
    Age      int    `json:"age" label:"年龄" validate:"required,gte=1,lte=120"`
    Phone    string `json:"phone" label:"手机号" validate:"required,len=11"`
}
```

**重要说明：**
- 使用 `validate` 标签而不是 `binding` 标签
- `label` 标签用于自定义字段名显示
- 如果没有 `label` 标签，会使用 `json` 标签
- 如果都没有，会使用字段名

### 2. 在 Gin 处理器中使用

```go
func registerHandler(c *gin.Context) {
    var req UserRequest
    
    // 绑定 JSON 数据
    if err := c.ShouldBindJSON(&req); err != nil {
        response.Error(c, gin.H{}, err.Error(), http.StatusBadRequest)
        return
    }

    // 使用中文验证器验证
    if err := validator.ValidateStruct(req); err != nil {
        response.Error(c, gin.H{}, err.Error(), http.StatusBadRequest)
        return
    }

    // 验证通过，处理业务逻辑
    response.Success(c, gin.H{}, "注册成功")
}
```

### 3. 直接使用验证器

```go
// 验证结构体并返回所有错误
err := validator.ValidateStruct(user)
if err != nil {
    fmt.Println(err.Error()) // 输出中文错误信息
}

// 验证结构体并返回第一个错误
err := validator.ValidateStructFirstError(user)
if err != nil {
    fmt.Println(err.Error()) // 输出第一个中文错误信息
}

// 直接使用验证器实例
err := validator.Validate.Struct(user)
if err != nil {
    // 翻译错误信息
    translatedErr := validator.TranslateError(err)
    fmt.Println(translatedErr)
}
```

## 错误信息示例

### 输入数据
```json
{
    "username": "ab",
    "email": "invalid-email",
    "password": "123",
    "age": 150,
    "phone": "138001380"
}
```

### 输出错误信息
```
用户名长度不能少于3个字符; 邮箱必须是有效的邮箱地址; 密码长度不能少于6个字符; 年龄必须小于或等于120; 手机号长度必须是11个字符
```

## 自定义验证规则

如果需要添加自定义验证规则，可以在 `pkg/validator/validator.go` 中修改：

```go
// 在 registerCustomTranslations 函数中添加
customTranslations := map[string]string{
    "custom_rule": "{0}不符合自定义规则",
    // ... 其他规则
}
```

## 注意事项

1. **标签使用**：必须使用 `validate` 标签，不能使用 `binding` 标签
2. **字段名**：优先使用 `label` 标签，其次 `json` 标签，最后字段名
3. **初始化**：验证器在包导入时自动初始化，无需手动初始化
4. **错误处理**：建议使用 `validator.ValidateStruct()` 函数，它会自动翻译错误信息

## 测试

运行测试文件验证功能：

```bash
go run cmd/scripts/test/main.go
```

运行 Gin 示例：

```bash
go run cmd/scripts/test/gin_example.go
```

然后使用 curl 测试：

```bash
# 测试注册接口
curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "ab",
    "email": "invalid-email",
    "password": "123",
    "age": 150,
    "phone": "138001380"
  }'
```
