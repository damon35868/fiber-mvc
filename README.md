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

## Installation/安装

```bash
$ git clone https://github.com/damon35868/fiber-mvc.git
```

## cmd/命令行

```bash
# 启动本地环境 local mode 获取 .env.local环境变量
$ make dev

# 启动本地环境 development mode 获取 .env.development环境变量
$ make devstart

# 启动本地环境 production mode 获取 .env.production环境变量
$ make devpro

# 执行数据库迁移
$ make migrateup

# 执行数据库回滚
$ make migratedown

# 生成model及增删改查
$ make modelgen

# 生成文档
$ make docgen
```
