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
* 基础用户登陆 
* 平滑重启 ✅

### 基础接口

* GET /me 获取当前登陆用户信息
* POST /login 登陆接口
* POST /refreshToken 刷新token, 该接口需要在 header 中传递 Refresh-Token

### 统一响应

```
utils.ResponseSuccess
utils.ResponseError
```