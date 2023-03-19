# gchat
基于golang实现的轻量IM

## 介绍
- 百万TPS
- 可以快速集成到当前应用也可以单服务部署grpc-server
- 用户对用户使用写扩散，群消息使用读扩散减小开销
- 高扩展hook
- 支持redis和文件两种持久化方式
- 历史消息查询
- SessionTicket会话恢复
- 以用户为单位，单个用户可以有多个客户端

## Install
```shell script
go mod tidy
```