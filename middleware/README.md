# 中间件

将客户端的请求URL重写后转发给上游

中间件的出参：

- `bool`，值为true表示中止后续逻辑
- `error`，中间件执行时的错误信息
