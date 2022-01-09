## 0. mimo结构图
![](./static/mimo-prod2.png)

## 1. 用途
用于测试websocket连接发消息

## 2. 使用
构建
> go build

帮助
```shell script
./ws-test -h
Usage:
  -addr string
        websocket连接地址 (default "ws://localhost:10000/ws/connect")
  -conns int
        连接用户数 (default 2)
  -uid string
        用户名前缀 (default "UserA-")
  -room string
        房间ID (default "room001")

  -forever
        是否循环发送房间消息. (default true)
  -interval int
        每个步骤的间隔时间,需要等待连接、加房间等.单位:秒 (default 5)
  -time int
        持续发送消息N秒后,停止程序
  -wait int
        发送消息的间隔时间.单位:毫秒 (default 5000)
  -waitn int
        添加时间误差范围,减少并发的压力.单位:毫秒 (default 2000)

  -content string
        消息内容
  -keeplive
        是否发送心跳包 (default true)
  -onlyWsconn
        只保持websocket连接 (default false)
  -level string
        输出日志级别 (default "info")
```  

运行
```shell script
# 已连接的用户只发送心跳包(keeplive=true,content="")
./ws-test -conns 100

# 已连接的用户不断发送自定义的消息体
./ws-test -conns 2 -time 10 -wait 500 -content "这是消息体" -keeplive=false

# 只发送心跳包，保持websocket连接，不加房间不发送消息
./ws-test -keeplive=true -onlyWsconn=true
```

输出
```shell script
1. 已创建2个用户连接
2. 正在启动读协程,等待5秒
3. 正在加入房间,等待5秒
4. 正在发送房间消息: 2020-05-27 18:41:15.7809834 +0800 CST m=+10.064575701
5. 注意: 程序在 10s 后停止运行
{"applicationId":"mimo-test","level":"error","msg":"server is closed.","time":"2020-05-27T18:41:25+08:00"}
{"applicationId":"mimo-test","level":"info","msg":"[UserA-0] 共收到 20 条消息, TPS:2.00, Time:10","time":"2020-05-27T18:41:25+08:00"}
{"applicationId":"mimo-test","level":"info","msg":"[UserA-1] 共收到 20 条消息, TPS:2.00, Time:10","time":"2020-05-27T18:41:25+08:00"}
{"applicationId":"mimo-test","level":"info","msg":"所有消息统计: 数量:40, TPS:4, 时间:10","time":"2020-05-27T18:41:25+08:00"}
0. 程序正在停止.请等待5秒
```

## 3. 压测
