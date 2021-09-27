# Go Garden 
[![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/panco95/go-garden) [![Go Report Card](https://goreportcard.com/badge/github.com/panco95/go-garden)](https://goreportcard.com/report/github.com/panco95/go-garden) 


Go Garden是一款面向分布式系统架构的微服务框架

## 概念

Go Garden为分布式系统架构的开发提供了核心需求，包括微服务的基础架构支持，例如gateway网关路由分发、负载均衡、服务调用链路追踪、接口级粒度可配服务限流、熔断策略等。

## 结构图

![struct](docs/struct.png "struct")


## 特性

- **注册发现**

- **网关路由**

- **负载均衡**

- **服务限流**

- **服务熔断**

- **服务重试**

- **超时控制**

- **动态路由**

- **自动同步**

- **安全认证**

- **链路追踪**

## 开发者必读

Go Garden考虑到开发者的使用门槛，并没有自行造轮子实现Http框架，而是选择使用较为成熟的Gin作为底层Http服务，这样可以大大减少开发者的学习路径，如果你是Go初学者，请提前阅读Gin框架文档。


## 快速开始

```golang
import "github.com/panco95/go-garden/core"

var service *core.Garden

func main() {
    service = core.New()
    service.Run(nil, nil)
}
```

## 教程：基于Go Garden快速构建微服务
访问 [基于Go Garden快速构建微服务](docs/tutorial.md) 学习如何使用Go
Garden

## 示例代码
访问 [examples](examples) 查看教程

## 许可证

Go Garden 包含 Apache 2.0 许可证