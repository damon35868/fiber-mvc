## Description/描述

基于 go fiber 搭建的 MVC 结构, 开箱即用！

## Features/主要功能

- sqlc 集成，并提供自动构建命令，快速生成 model 及增删改查代码
- 抽离配置文件
- 抽离控制器、服务、中间件等结构
- 集成 swagger
- 集成 mysql 操作
- 集成 redis 操作
- 分布式锁
- 结构化环境
- 命令行工具
- 抽离公共异常处理及返回体
- 集成定时任务

## Installation/安装

```bash
$ git clone https://github.com/damon35868/fiber-mvc.git
```

## initialization/初始化

```bash
# install golang-migrate 非Mac请下载二进制 https://github.com/golang-migrate/migrate/releases
$ brew install golang-migrate

# install sqlc
$ go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# install air
$ go install github.com/air-verse/air@latest

# 添加环境变量并设置你的参数
$ cp .env.example ./.env.local && cp .env.example ./.env.development && cp .env.example ./.env.production
```

## cmd/命令行

```bash
# 启动本地环境 local mode 获取 .env.local环境变量
$ make dev

# 启动本地环境 development mode 获取 .env.development环境变量
$ make dev.start

# 启动本地环境 production mode 获取 .env.production环境变量
$ make dev.pro

# 执行数据库迁移
$ make migrate

# 执行数据库回滚
$ make migrate.rollback

# 生成model及增删改查
$ make gen.model

# 生成文档
$ make gen.doc
```
