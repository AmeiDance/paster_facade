# paster_facade

Paster 服务端门面模块，以 [Gin](https://github.com/gin-gonic/gin) 作为 Web 服务端框架，使用字节跳动开源的微服务 RPC 框架 [KiteX](https://github.com/cloudwego/kitex)
，通过 [Apache Thrift](https://github.com/apache/thrift) 协议与下游核心模块 [paster_core](https://github.com/AmeiDance/paster_core) 通信。

[Sequence Diagram](https://mermaid-js.github.io/mermaid-live-editor/edit/#eyJjb2RlIjoic2VxdWVuY2VEaWFncmFtXG4gICAgcGFydGljaXBhbnQgZmVcbiAgICBwYXJ0aWNpcGFudCBmYWNhZGVcbiAgICBwYXJ0aWNpcGFudCBjb3JlXG4gICAgcGFydGljaXBhbnQgcmVkaXNcbiAgICBwYXJ0aWNpcGFudCBteXNxbFxuXG4gICAgZmUtPj5mZTogL1xuICAgIGFjdGl2YXRlIGZlXG4gICAgZGVhY3RpdmF0ZSBmZVxuXG4gICAgZmUtPj5mYWNhZGU6IFBPU1QgL3Bvc3Qvc2F2ZS9cbiAgICBhY3RpdmF0ZSBmZVxuICAgIGFjdGl2YXRlIGZhY2FkZVxuICAgIGZhY2FkZS0-PnJlZGlzOiBnZXRcbiAgICBhY3RpdmF0ZSByZWRpc1xuICAgIHJlZGlzLS0-PmZhY2FkZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSByZWRpc1xuICAgIGZhY2FkZS0-PmNvcmU6IFNhdmVQb3N0XG4gICAgYWN0aXZhdGUgY29yZVxuICAgIGNvcmUtPj5teXNxbDogd3JpdGVcbiAgICBhY3RpdmF0ZSBteXNxbFxuICAgIG15c3FsLS0-PmNvcmU6IHJldHVyblxuICAgIGRlYWN0aXZhdGUgbXlzcWxcbiAgICBjb3JlLS0-PmZhY2FkZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBjb3JlXG4gICAgZmFjYWRlLT4-cmVkaXM6IHNldFxuICAgIGFjdGl2YXRlIHJlZGlzXG4gICAgcmVkaXMtLT4-ZmFjYWRlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIHJlZGlzXG4gICAgZmFjYWRlLS0-PmZlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIGZhY2FkZVxuICAgIGRlYWN0aXZhdGUgZmVcblxuICAgIGZlLT4-ZmU6IC88aWQ-XG4gICAgYWN0aXZhdGUgZmVcbiAgICBkZWFjdGl2YXRlIGZlXG5cbiAgICBmZS0-PmZhY2FkZTogR0VUIC9jb21tZW50L2dldC9cbiAgICBhY3RpdmF0ZSBmZVxuICAgIGFjdGl2YXRlIGZhY2FkZVxuICAgIGZhY2FkZS0-PmNvcmU6IEdldENvbW1lbnRzXG4gICAgYWN0aXZhdGUgY29yZVxuICAgIGNvcmUtPj5teXNxbDogcmVhZFxuICAgIGFjdGl2YXRlIG15c3FsXG4gICAgbXlzcWwtLT4-Y29yZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBteXNxbFxuICAgIGNvcmUtLT4-ZmFjYWRlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIGNvcmVcbiAgICBhY3RpdmF0ZSByZWRpc1xuICAgIGZhY2FkZS0tPj5mZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBmYWNhZGVcbiAgICBkZWFjdGl2YXRlIGZlXG5cbiAgICBmZS0-PmZhY2FkZTogUE9TVCAvcG9zdC9nZXQvXG4gICAgYWN0aXZhdGUgZmVcbiAgICBhY3RpdmF0ZSBmYWNhZGVcbiAgICBmYWNhZGUtPj5jb3JlOiBHZXRQb3N0XG4gICAgYWN0aXZhdGUgY29yZVxuICAgIGNvcmUtPj5teXNxbDogcmVhZFxuICAgIGFjdGl2YXRlIG15c3FsXG4gICAgbXlzcWwtLT4-Y29yZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBteXNxbFxuICAgIGNvcmUtLT4-ZmFjYWRlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIGNvcmVcbiAgICBmYWNhZGUtPj5jb3JlOiAoRGVsZXRlUG9zdClcbiAgICBhY3RpdmF0ZSBjb3JlXG4gICAgY29yZS0tPj5mYWNhZGU6IHJldHVyblxuICAgIGRlYWN0aXZhdGUgY29yZVxuICAgIGZhY2FkZS0tPj5mZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBmYWNhZGVcbiAgICBkZWFjdGl2YXRlIGZlXG5cbiAgICBmZS0-PmZhY2FkZTogUE9TVCAvY29tbWVudC9zYXZlL1xuICAgIGFjdGl2YXRlIGZlXG4gICAgYWN0aXZhdGUgZmFjYWRlXG4gICAgZmFjYWRlLT4-cmVkaXM6IGdldFxuICAgIGFjdGl2YXRlIHJlZGlzXG4gICAgcmVkaXMtLT4-ZmFjYWRlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIHJlZGlzXG4gICAgZmFjYWRlLT4-Y29yZTogU2F2ZUNvbW1lbnRcbiAgICBhY3RpdmF0ZSBjb3JlXG4gICAgY29yZS0-Pm15c3FsOiB3cml0ZVxuICAgIGFjdGl2YXRlIG15c3FsXG4gICAgbXlzcWwtLT4-Y29yZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBteXNxbFxuICAgIGNvcmUtLT4-ZmFjYWRlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIGNvcmVcbiAgICBmYWNhZGUtPj5yZWRpczogc2V0XG4gICAgYWN0aXZhdGUgcmVkaXNcbiAgICByZWRpcy0tPj5mYWNhZGU6IHJldHVyblxuICAgIGRlYWN0aXZhdGUgcmVkaXNcbiAgICBmYWNhZGUtLT4-ZmU6IHJldHVyblxuICAgIGRlYWN0aXZhdGUgZmFjYWRlXG4gICAgZGVhY3RpdmF0ZSBmZVxuIiwibWVybWFpZCI6IntcbiAgXCJ0aGVtZVwiOiBcImRlZmF1bHRcIlxufSIsInVwZGF0ZUVkaXRvciI6ZmFsc2UsImF1dG9TeW5jIjp0cnVlLCJ1cGRhdGVEaWFncmFtIjpmYWxzZX0)

![](https://mermaid.ink/svg/eyJjb2RlIjoic2VxdWVuY2VEaWFncmFtXG4gICAgcGFydGljaXBhbnQgZmVcbiAgICBwYXJ0aWNpcGFudCBmYWNhZGVcbiAgICBwYXJ0aWNpcGFudCBjb3JlXG4gICAgcGFydGljaXBhbnQgcmVkaXNcbiAgICBwYXJ0aWNpcGFudCBteXNxbFxuXG4gICAgZmUtPj5mZTogL1xuICAgIGFjdGl2YXRlIGZlXG4gICAgZGVhY3RpdmF0ZSBmZVxuXG4gICAgZmUtPj5mYWNhZGU6IFBPU1QgL3Bvc3Qvc2F2ZS9cbiAgICBhY3RpdmF0ZSBmZVxuICAgIGFjdGl2YXRlIGZhY2FkZVxuICAgIGZhY2FkZS0-PnJlZGlzOiBnZXRcbiAgICBhY3RpdmF0ZSByZWRpc1xuICAgIHJlZGlzLS0-PmZhY2FkZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSByZWRpc1xuICAgIGZhY2FkZS0-PmNvcmU6IFNhdmVQb3N0XG4gICAgYWN0aXZhdGUgY29yZVxuICAgIGNvcmUtPj5teXNxbDogd3JpdGVcbiAgICBhY3RpdmF0ZSBteXNxbFxuICAgIG15c3FsLS0-PmNvcmU6IHJldHVyblxuICAgIGRlYWN0aXZhdGUgbXlzcWxcbiAgICBjb3JlLS0-PmZhY2FkZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBjb3JlXG4gICAgZmFjYWRlLT4-cmVkaXM6IHNldFxuICAgIGFjdGl2YXRlIHJlZGlzXG4gICAgcmVkaXMtLT4-ZmFjYWRlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIHJlZGlzXG4gICAgZmFjYWRlLS0-PmZlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIGZhY2FkZVxuICAgIGRlYWN0aXZhdGUgZmVcblxuICAgIGZlLT4-ZmU6IC88aWQ-XG4gICAgYWN0aXZhdGUgZmVcbiAgICBkZWFjdGl2YXRlIGZlXG5cbiAgICBmZS0-PmZhY2FkZTogR0VUIC9jb21tZW50L2dldC9cbiAgICBhY3RpdmF0ZSBmZVxuICAgIGFjdGl2YXRlIGZhY2FkZVxuICAgIGZhY2FkZS0-PmNvcmU6IEdldENvbW1lbnRzXG4gICAgYWN0aXZhdGUgY29yZVxuICAgIGNvcmUtPj5teXNxbDogcmVhZFxuICAgIGFjdGl2YXRlIG15c3FsXG4gICAgbXlzcWwtLT4-Y29yZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBteXNxbFxuICAgIGNvcmUtLT4-ZmFjYWRlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIGNvcmVcbiAgICBhY3RpdmF0ZSByZWRpc1xuICAgIGZhY2FkZS0tPj5mZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBmYWNhZGVcbiAgICBkZWFjdGl2YXRlIGZlXG5cbiAgICBmZS0-PmZhY2FkZTogUE9TVCAvcG9zdC9nZXQvXG4gICAgYWN0aXZhdGUgZmVcbiAgICBhY3RpdmF0ZSBmYWNhZGVcbiAgICBmYWNhZGUtPj5jb3JlOiBHZXRQb3N0XG4gICAgYWN0aXZhdGUgY29yZVxuICAgIGNvcmUtPj5teXNxbDogcmVhZFxuICAgIGFjdGl2YXRlIG15c3FsXG4gICAgbXlzcWwtLT4-Y29yZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBteXNxbFxuICAgIGNvcmUtLT4-ZmFjYWRlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIGNvcmVcbiAgICBmYWNhZGUtPj5jb3JlOiAoRGVsZXRlUG9zdClcbiAgICBhY3RpdmF0ZSBjb3JlXG4gICAgY29yZS0tPj5mYWNhZGU6IHJldHVyblxuICAgIGRlYWN0aXZhdGUgY29yZVxuICAgIGZhY2FkZS0tPj5mZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBmYWNhZGVcbiAgICBkZWFjdGl2YXRlIGZlXG5cbiAgICBmZS0-PmZhY2FkZTogUE9TVCAvY29tbWVudC9zYXZlL1xuICAgIGFjdGl2YXRlIGZlXG4gICAgYWN0aXZhdGUgZmFjYWRlXG4gICAgZmFjYWRlLT4-cmVkaXM6IGdldFxuICAgIGFjdGl2YXRlIHJlZGlzXG4gICAgcmVkaXMtLT4-ZmFjYWRlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIHJlZGlzXG4gICAgZmFjYWRlLT4-Y29yZTogU2F2ZUNvbW1lbnRcbiAgICBhY3RpdmF0ZSBjb3JlXG4gICAgY29yZS0-Pm15c3FsOiB3cml0ZVxuICAgIGFjdGl2YXRlIG15c3FsXG4gICAgbXlzcWwtLT4-Y29yZTogcmV0dXJuXG4gICAgZGVhY3RpdmF0ZSBteXNxbFxuICAgIGNvcmUtLT4-ZmFjYWRlOiByZXR1cm5cbiAgICBkZWFjdGl2YXRlIGNvcmVcbiAgICBmYWNhZGUtPj5yZWRpczogc2V0XG4gICAgYWN0aXZhdGUgcmVkaXNcbiAgICByZWRpcy0tPj5mYWNhZGU6IHJldHVyblxuICAgIGRlYWN0aXZhdGUgcmVkaXNcbiAgICBmYWNhZGUtLT4-ZmU6IHJldHVyblxuICAgIGRlYWN0aXZhdGUgZmFjYWRlXG4gICAgZGVhY3RpdmF0ZSBmZVxuIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjpmYWxzZSwiYXV0b1N5bmMiOnRydWUsInVwZGF0ZURpYWdyYW0iOmZhbHNlfQ)