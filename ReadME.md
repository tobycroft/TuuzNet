# TuuzNet-NetworkPackage

TuuzNet is a full function simple but high performant http client

### What is TuuzNet?

TuuzNet is a Network package for http request curl and websocket, focus in high performance but simple use.

### Update Notice

There will be a huge difference between v1.0 and 1.1 DO NOT UPDATE, the opreration way is different!

# Advantage

- simple use 简单易用
- High performance 高性能
- Full function 全功能

# How to use(v1.1.0+)

范例1，使用Websocket作为客户端链接到远程，这里以Shamrock（机器人）作为简单的Demo

```go
var ws Net.WsClient
ws.NewConnect("ws://10.0.1.102:5801")

send := Send{Action: "send_private_msg", Params: struct {
UserId  int    `json:"user_id"`
Message string `json:"message"`
}(struct {
UserId  int
Message string
}{UserId: 710209520, Message: "test"}), Echo: "test",
}
bt, _ := sonic.Marshal(send)
go func () {
time.Sleep(3 * time.Second)
ws.WriteChannel <- bt
}()
for {
fmt.Println(string(<-ws.ReadChannel))
}
```

范例2，使用NetPost