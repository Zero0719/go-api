## 使用步骤

### 1. 克隆模板并初始化项目

```bash
# 克隆模板仓库到你的项目目录
git clone https://github.com/your-username/go-api-template.git my-new-project
cd my-new-project

# 删除原有的 Git 历史记录，准备初始化新的仓库
rm -rf .git
```

### 2. 配置项目信息

#### 2.1 全局替换模块名

在项目根目录执行以下命令，将所有 Go 文件中的模块名替换为你的项目模块名：

```bash
# 查找并替换所有 .go 文件中的模块名
# macOS/Linux 系统使用：
find . -type f -name "*.go" -exec sed -i '' 's|github.com/Zero0719/go-api|github.com/your-username/your-project|g' {} +

# Linux 系统（非 macOS）使用：
# find . -type f -name "*.go" -exec sed -i 's|github.com/Zero0719/go-api|github.com/your-username/your-project|g' {} +
```

#### 2.2 更新 go.mod 文件

```bash
# 修改模块名
go mod edit -module github.com/your-username/your-project

# 整理依赖关系
go mod tidy
```

#### 2.3 修改应用名称

编辑 `main.go` 文件，将应用名称修改为你的项目名称：

```go
// main.go
var rootCmd = &cobra.Command{
    Use: "YourAppName",  // 修改为你的应用名称
    Short: "Your project description",  // 修改为你的项目描述
}
```

### 3. 安装项目依赖

```bash
# 下载所有依赖包
go mod download
```

### 4. 配置环境变量

```bash
# 复制环境变量模板文件
cp .env.example .env

# 编辑环境变量文件，配置数据库、Redis 等连接信息
vim .env
# 或使用你喜欢的编辑器打开 .env 文件进行编辑
```

### 5. 启动开发服务器

```bash
# 方式一：使用 Air 进行热重载开发（推荐）
# Air 会在代码变更时自动重新编译和运行
air

# 方式二：直接运行（需要手动重启）
go run main.go serve
```