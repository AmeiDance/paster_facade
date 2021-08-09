# paster_facade

Paster 服务端门面模块，使用 Web 服务端框架 [Gin](https://github.com/gin-gonic/gin) 提供 HTTP 接口，使用字节跳动开源的微服务 RPC
框架 [KiteX](https://github.com/cloudwego/kitex) 通过 [Thrift](https://github.com/apache/thrift)
协议与下游核心模块 [paster_core](https://github.com/ameidance/paster_core) 通信。

<details>
<summary><b>时序图</b></summary>
<a href="https://mermaid-js.github.io/mermaid-live-editor/edit/##eyJjb2RlIjoic2VxdWVuY2VEaWFncmFtXG4gICAgcGFydGljaXBhbnQgZmVcbiAgICBwYXJ0aWNpcGFudCBmYWNhZGVcbiAgICBwYXJ0aWNpcGFudCBjb3JlXG4gICAgcGFydGljaXBhbnQgcmVkaXNcbiAgICBwYXJ0aWNpcGFudCBteXNxbFxuXG4gICAgZmUtPj5mZTogL1xuICAgIGFjdGl2YXRlIGZlXG4gICAgZGVhY3RpdmF0ZSBmZVxuXG4gICAgZmUtPj5mYWNhZGU6IFBPU1QgL3Bvc3Qvc2F2ZS9cbiAgICBhY3RpdmF0ZSBmZVxuICAgIGFjdGl2YXRlIGZhY2FkZVxuICAgIGZhY2FkZS0-PnJlZGlzOiBnZXRcbiAgICBhY3RpdmF0ZSByZWRpc1xuICAgIHJlZGlzLS0-PmZhY2FkZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSByZWRpc1xuICAgIGZhY2FkZS0-PnJlZGlzOiBzZXRcbiAgICBhY3RpdmF0ZSByZWRpc1xuICAgIHJlZGlzLS0-PmZhY2FkZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSByZWRpc1xuICAgIGZhY2FkZS0-PmNvcmU6IFNhdmVQb3N0XG4gICAgYWN0aXZhdGUgY29yZVxuICAgIGNvcmUtPj5teXNxbDogd3JpdGVcbiAgICBhY3RpdmF0ZSBteXNxbFxuICAgIG15c3FsLS0-PmNvcmU6IHJldHVyblxuICAgIGRlYWN0aXZhdGUgbXlzcWxcbiAgICBjb3JlLS0-PmZhY2FkZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBjb3JlXG4gICAgZmFjYWRlLS0-PmZlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIGZhY2FkZVxuICAgIGRlYWN0aXZhdGUgZmVcblxuICAgIGZlLT4-ZmU6IC88aWQ-XG4gICAgYWN0aXZhdGUgZmVcbiAgICBkZWFjdGl2YXRlIGZlXG5cbiAgICBmZS0-PmZhY2FkZTogR0VUIC9jb21tZW50L2dldC9cbiAgICBhY3RpdmF0ZSBmZVxuICAgIGFjdGl2YXRlIGZhY2FkZVxuICAgIGZhY2FkZS0-PmNvcmU6IEdldENvbW1lbnRzXG4gICAgYWN0aXZhdGUgY29yZVxuICAgIGNvcmUtPj5teXNxbDogcmVhZFxuICAgIGFjdGl2YXRlIG15c3FsXG4gICAgbXlzcWwtLT4-Y29yZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBteXNxbFxuICAgIGNvcmUtLT4-ZmFjYWRlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIGNvcmVcbiAgICBhY3RpdmF0ZSByZWRpc1xuICAgIGZhY2FkZS0tPj5mZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBmYWNhZGVcbiAgICBkZWFjdGl2YXRlIGZlXG5cbiAgICBmZS0-PmZhY2FkZTogUE9TVCAvcG9zdC9nZXQvXG4gICAgYWN0aXZhdGUgZmVcbiAgICBhY3RpdmF0ZSBmYWNhZGVcbiAgICBmYWNhZGUtPj5jb3JlOiBHZXRQb3N0XG4gICAgYWN0aXZhdGUgY29yZVxuICAgIGNvcmUtPj5teXNxbDogcmVhZFxuICAgIGFjdGl2YXRlIG15c3FsXG4gICAgbXlzcWwtLT4-Y29yZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBteXNxbFxuICAgIGNvcmUtLT4-ZmFjYWRlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIGNvcmVcbiAgICBmYWNhZGUtPj5jb3JlOiAoRGVsZXRlUG9zdClcbiAgICBhY3RpdmF0ZSBjb3JlXG4gICAgY29yZS0tPj5mYWNhZGU6IHJldHVyblxuICAgIGRlYWN0aXZhdGUgY29yZVxuICAgIGZhY2FkZS0tPj5mZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBmYWNhZGVcbiAgICBkZWFjdGl2YXRlIGZlXG5cbiAgICBmZS0-PmZhY2FkZTogUE9TVCAvY29tbWVudC9zYXZlL1xuICAgIGFjdGl2YXRlIGZlXG4gICAgYWN0aXZhdGUgZmFjYWRlXG4gICAgZmFjYWRlLT4-cmVkaXM6IGdldFxuICAgIGFjdGl2YXRlIHJlZGlzXG4gICAgcmVkaXMtLT4-ZmFjYWRlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIHJlZGlzXG5cbiAgICBmYWNhZGUtPj5jb3JlOiBTYXZlQ29tbWVudFxuICAgIGFjdGl2YXRlIGNvcmVcbiAgICBjb3JlLT4-bXlzcWw6IHdyaXRlXG4gICAgYWN0aXZhdGUgbXlzcWxcbiAgICBteXNxbC0tPj5jb3JlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIG15c3FsXG4gICAgY29yZS0tPj5mYWNhZGU6IHJldHVyblxuICAgIGRlYWN0aXZhdGUgY29yZVxuICAgIGZhY2FkZS0tPj5mZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBmYWNhZGVcbiAgICBkZWFjdGl2YXRlIGZlXG4iLCJtZXJtYWlkIjoie1xuICBcInRoZW1lXCI6IFwiZGVmYXVsdFwiXG59IiwidXBkYXRlRWRpdG9yIjpmYWxzZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOmZhbHNlfQ"><img src="https://mermaid.ink/svg/eyJjb2RlIjoic2VxdWVuY2VEaWFncmFtXG4gICAgcGFydGljaXBhbnQgZmVcbiAgICBwYXJ0aWNpcGFudCBmYWNhZGVcbiAgICBwYXJ0aWNpcGFudCBjb3JlXG4gICAgcGFydGljaXBhbnQgcmVkaXNcbiAgICBwYXJ0aWNpcGFudCBteXNxbFxuXG4gICAgZmUtPj5mZTogL1xuICAgIGFjdGl2YXRlIGZlXG4gICAgZGVhY3RpdmF0ZSBmZVxuXG4gICAgZmUtPj5mYWNhZGU6IFBPU1QgL3Bvc3Qvc2F2ZS9cbiAgICBhY3RpdmF0ZSBmZVxuICAgIGFjdGl2YXRlIGZhY2FkZVxuICAgIGZhY2FkZS0-PnJlZGlzOiBnZXRcbiAgICBhY3RpdmF0ZSByZWRpc1xuICAgIHJlZGlzLS0-PmZhY2FkZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSByZWRpc1xuICAgIGZhY2FkZS0-PnJlZGlzOiBzZXRcbiAgICBhY3RpdmF0ZSByZWRpc1xuICAgIHJlZGlzLS0-PmZhY2FkZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSByZWRpc1xuICAgIGZhY2FkZS0-PmNvcmU6IFNhdmVQb3N0XG4gICAgYWN0aXZhdGUgY29yZVxuICAgIGNvcmUtPj5teXNxbDogd3JpdGVcbiAgICBhY3RpdmF0ZSBteXNxbFxuICAgIG15c3FsLS0-PmNvcmU6IHJldHVyblxuICAgIGRlYWN0aXZhdGUgbXlzcWxcbiAgICBjb3JlLS0-PmZhY2FkZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBjb3JlXG4gICAgZmFjYWRlLS0-PmZlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIGZhY2FkZVxuICAgIGRlYWN0aXZhdGUgZmVcblxuICAgIGZlLT4-ZmU6IC88aWQ-XG4gICAgYWN0aXZhdGUgZmVcbiAgICBkZWFjdGl2YXRlIGZlXG5cbiAgICBmZS0-PmZhY2FkZTogR0VUIC9jb21tZW50L2dldC9cbiAgICBhY3RpdmF0ZSBmZVxuICAgIGFjdGl2YXRlIGZhY2FkZVxuICAgIGZhY2FkZS0-PmNvcmU6IEdldENvbW1lbnRzXG4gICAgYWN0aXZhdGUgY29yZVxuICAgIGNvcmUtPj5teXNxbDogcmVhZFxuICAgIGFjdGl2YXRlIG15c3FsXG4gICAgbXlzcWwtLT4-Y29yZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBteXNxbFxuICAgIGNvcmUtLT4-ZmFjYWRlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIGNvcmVcbiAgICBhY3RpdmF0ZSByZWRpc1xuICAgIGZhY2FkZS0tPj5mZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBmYWNhZGVcbiAgICBkZWFjdGl2YXRlIGZlXG5cbiAgICBmZS0-PmZhY2FkZTogUE9TVCAvcG9zdC9nZXQvXG4gICAgYWN0aXZhdGUgZmVcbiAgICBhY3RpdmF0ZSBmYWNhZGVcbiAgICBmYWNhZGUtPj5jb3JlOiBHZXRQb3N0XG4gICAgYWN0aXZhdGUgY29yZVxuICAgIGNvcmUtPj5teXNxbDogcmVhZFxuICAgIGFjdGl2YXRlIG15c3FsXG4gICAgbXlzcWwtLT4-Y29yZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBteXNxbFxuICAgIGNvcmUtLT4-ZmFjYWRlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIGNvcmVcbiAgICBmYWNhZGUtPj5jb3JlOiAoRGVsZXRlUG9zdClcbiAgICBhY3RpdmF0ZSBjb3JlXG4gICAgY29yZS0tPj5mYWNhZGU6IHJldHVyblxuICAgIGRlYWN0aXZhdGUgY29yZVxuICAgIGZhY2FkZS0tPj5mZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBmYWNhZGVcbiAgICBkZWFjdGl2YXRlIGZlXG5cbiAgICBmZS0-PmZhY2FkZTogUE9TVCAvY29tbWVudC9zYXZlL1xuICAgIGFjdGl2YXRlIGZlXG4gICAgYWN0aXZhdGUgZmFjYWRlXG4gICAgZmFjYWRlLT4-cmVkaXM6IGdldFxuICAgIGFjdGl2YXRlIHJlZGlzXG4gICAgcmVkaXMtLT4-ZmFjYWRlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIHJlZGlzXG4gICAgZmFjYWRlLT4-cmVkaXM6IHNldFxuICAgIGFjdGl2YXRlIHJlZGlzXG4gICAgcmVkaXMtLT4-ZmFjYWRlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIHJlZGlzXG4gICAgZmFjYWRlLT4-Y29yZTogU2F2ZUNvbW1lbnRcbiAgICBhY3RpdmF0ZSBjb3JlXG4gICAgY29yZS0-Pm15c3FsOiB3cml0ZVxuICAgIGFjdGl2YXRlIG15c3FsXG4gICAgbXlzcWwtLT4-Y29yZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBteXNxbFxuICAgIGNvcmUtLT4-ZmFjYWRlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIGNvcmVcbiAgICBmYWNhZGUtLT4-ZmU6IHJldHVyblxuICAgIGRlYWN0aXZhdGUgZmFjYWRlXG4gICAgZGVhY3RpdmF0ZSBmZVxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjpmYWxzZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOmZhbHNlfQ"></img></a>
</details>

### 环境 & 运行

> ⚠️ 由于 [KiteX](https://github.com/cloudwego/kitex) 框架依赖网络库 [Netpoll](https://github.com/cloudwego/netpoll) **不支持 Windows**，目前项目只能在 **Linux**, **macOS** 环境下构建。

本服务模块依赖以下组件：

- [Redis](https://github.com/redis/redis)
- [Consul](https://github.com/hashicorp/consul)

首先复制 `conf` 目录下所有 `*.example.yml` 配置文件并重命名为 `*.yml`，之后逐一配置。

- `gin.yml` 为 Gin 框架服务启动配置，默认端口号为 5000
- `redis.yml` 为 Redis 连接配置
- `consul.yml` 为 Consul 连接配置

配置完成后在项目根目录下执行以下命令来构建、运行项目。

```bash
sh build.sh
sh output/bootstrap.sh
```

> ⚠️ 请确保在运行项目前已经按照配置信息启动 Redis 和 Consul 服务端。

### Todo

- [x] 实现 KiteX 服务发现扩展接口，使用 [Consul](https://github.com/hashicorp/consul) 服务注册与发现
- [x] 新增 frame 层中间件优化 KiteX 框架请求响应日志
- [ ] 将项目打包成 Dokcer 镜像
- [ ] ...
