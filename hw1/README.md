# HW1

## 题目

```
HW1
编写一个 HTTP服务器。
1.接收客户端request，并将request中带的 header写入response header
2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
3.Server端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
4.当访问localhost/healthz 时，应返回 200
```

## 执行
```shell
make run
```
运行端口为：8080
