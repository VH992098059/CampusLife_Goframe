# 基于GoFrame的活动管理系统

## 项目简介

这是一个采用前后端分离架构的综合性活动管理平台，为用户提供活动发布、报名、管理等全流程服务。项目采用现代化的技术栈，具备完整的用户认证、权限管理和数据持久化功能。

## 技术架构

### 后端技术栈
- **框架**: GoFrame v2.9.0（Go语言企业级开发框架）
- **数据库**: MySQL 8.0 + Redis（数据存储与缓存）
- **认证方式**: GToken（基于Redis的分布式Token认证）
- **第三方服务**: 阿里云短信服务（验证码发送）
- **配置管理**: Nacos（分布式配置中心）
- **API设计**: RESTful风格，支持跨域访问

### 前端技术栈
- **框架**: React 18 + TypeScript
- **UI组件库**: Ant Design 5.24.5
- **构建工具**: Vite 6.0
- **路由管理**: React Router DOM 7.1
- **HTTP客户端**: Axios
- **样式处理**: Sass + CSS Modules

## 核心功能

### 1. 用户管理系统
- 用户注册/登录（支持手机号验证码登录）
- 多平台登录支持
- 用户信息管理（头像、昵称、联系方式等）
- 基于Redis的分布式会话管理

### 2. 活动管理系统
- 活动发布与编辑（支持富文本编辑）
- 活动状态管理（草稿、待审核、报名中、进行中、已结束）
- 活动搜索与分类筛选
- 活动热度统计
- 报名人数限制与统计

### 3. 报名管理系统
- 在线报名功能
- 报名审核流程
- 报名状态跟踪
- 活动参与者管理

### 4. 权限与安全
- JWT Token认证机制
- 密码加盐加密存储
- 跨域请求处理
- 中间件权限验证

## 项目结构

```
demo3/
├── api/                    # API接口定义
│   └── backapi/           # 后端API结构体
├── front_web/             # 前端React项目
│   ├── src/
│   │   ├── api/          # API调用
│   │   ├── page/         # 页面组件
│   │   ├── layout/       # 布局组件
│   │   └── router/       # 路由配置
│   └── package.json
├── internal/              # 后端核心代码
│   ├── cmd/              # 启动命令
│   ├── controller/       # 控制器层
│   ├── service/          # 业务逻辑层
│   ├── dao/              # 数据访问层
│   └── model/            # 数据模型
├── utility/               # 工具包
│   ├── gtoken/           # Token认证
│   ├── aliyunCode/       # 阿里云短信
│   └── middleware/       # 中间件
├── manifest/              # 部署配置
│   ├── config/           # 配置文件
│   └── docker/           # Docker配置
└── linuxdemo.sql         # 数据库脚本
```

## 快速开始

### 环境要求
- Go 1.23+
- Node.js 18+
- MySQL 8.0+
- Redis 6.0+

### 后端启动

1. 克隆项目
```bash
git clone <repository-url>
cd demo3
```

2. 安装依赖
```bash
go mod tidy
```

3. 配置数据库
```bash
# 导入数据库脚本
mysql -u root -p < linuxdemo.sql
```

4. 配置文件
```bash
# 修改 hack/config.yaml 中的数据库和Redis配置
```

5. 启动服务
```bash
go run main.go
```

### 前端启动

1. 进入前端目录
```bash
cd front_web
```

2. 安装依赖
```bash
pnpm install
```

3. 启动开发服务器
```bash
pnpm dev
```

## API文档

### 用户相关
- `POST /backend/register` - 用户注册
- `POST /backend/login` - 用户登录
- `GET /backend/info` - 获取用户信息
- `GET /backend/sendCode` - 发送验证码

### 活动相关
- `GET /backend/activityList` - 获取活动列表
- `GET /backend/GetActivity` - 获取活动详情
- `GET /backend/activitySearch` - 搜索活动
- `POST /backend/join` - 参加活动

## 数据库设计

### 核心数据表
- `user_info` - 用户基础信息表
- `activity_info` - 活动详情表
- `activity_order` - 活动报名订单表

## 部署说明

### Docker部署
```bash
# 构建镜像
docker build -t demo3 .

# 运行容器
docker run -p 8000:8000 demo3
```

### 生产环境配置
- 配置Redis集群
- 配置MySQL主从
- 配置Nginx反向代理
- 配置SSL证书

## 开发规范

- 遵循GoFrame开发规范
- 使用Git Flow工作流
- 代码提交前进行单元测试
- API接口需要编写文档

## 许可证

MIT License

## 联系方式

如有问题，请提交Issue或联系开发团队。