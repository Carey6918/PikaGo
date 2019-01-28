# Carey's rpc

期望完成的内容：

- [x] 服务注册与发现 (consul)
- [x] 健康检查
- [ ] 调用链追踪 (opentracing + zipkin)
- [ ] 调用链路日志 (zap or logrus)
- [ ] 各项指标监控

使用指南：

1. 样例通过docker启动consul，因此需要下载[docker](https://www.docker.com/get-started)

2. 运行启动脚本

```bash
$ cd example/server
$ ./build.sh
$ ./run.sh
```

3. 运行客户端

```bash
$ cd example/client
$ ./run.sh
```

注：目前Logger尚未实现，因此暂时无法查看日志。