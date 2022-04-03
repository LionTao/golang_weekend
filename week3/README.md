# 第三周作业

## 问题描述

基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

## 实现方案

根据描述信息，可以简单汇总成3块内容：

- 实现 HTTP server 的启动和关闭
- 监听 linux signal信号，支持 kill -9 或 Ctrl+C 的中断操作操作
- errgroup 实现多个 goroutine 的级联退出

按照实现方案，我们将任务拆分成多个小的功能。

1. 实现 http server 的启动和关闭功能
2. 使用 chan 实现对 linux signal 中断的注册和处理
3. 通过 errgroup + context 的形式，对 1、2中的 goroutine 进行级联注销

## 代码

[main.go](./main.go)