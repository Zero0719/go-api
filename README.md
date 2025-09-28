## GO API

基于 `gin` 实现的简单 `api` 项目

### Feature

* CORS 跨域处理 ✅
* 统一响应处理 ✅
* 异常处理 ✅
* 日志记录 ✅
* 请求/响应信息记录 ✅
* JWT 验证 ✅
* 配置读取 ✅
* 基础用户登陆 ✅
* 平滑重启 ✅
* 引入mysql 和 redis ✅

### 使用

项目拉下来以后，修改 `config.yaml` 中相关配置

### 基础接口

* GET / 测试接口
* GET /me 获取当前登陆用户信息
* POST /login 登陆接口
* POST /refreshToken 刷新token, 该接口需要在 header 中传递 Refresh-Token
* POST /user 创建用户，可根据自身需要移除该接口

### 统一响应

```
utils.ResponseSuccess
utils.ResponseError
```

### 使用日志

```
utils.Logger.Info().Msg("this is log msg")
```

### 使用数据库

详细用法请查看 `gorm`

```
var user User
if err := DB.Where("username = ?", username).Limit(1).Find(&user).Error; err != nil {
    return User{}, err
}
return user, nil
```

### 使用 reids

详细请查看相关的包文档

```
utils.Redis.Set(ctx.Background(), "test-key", "test-value", 0)
```

