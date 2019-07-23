# Carey's rpc

期望完成的内容：

- [x] 服务注册与发现 (consul)
- [x] 健康检查
- [x] 调用链追踪 (opentracing + zipkin)
- [x] 链路日志 (logrus)
- [ ] 各项指标监控（prometheus）

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

4. 访问http://127.0.0.1:9411/zipkin/可以查看调用链