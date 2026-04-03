# TuuzNet-NetworkPackage

TuuzNet is a full function simple but high performant http client

For more please visit project wiki

### What is TuuzNet?

TuuzNet is a Network package for http request curl and websocket, focus in high performance but simple use.

### Update Notice

There will be a huge difference between v1.0 and 1.1 DO NOT UPDATE, the opreration way is different!

After v1.1 the version upgrades will be under compatible considered

# Advantage

- simple use 简单易用
- High performance 高性能
- Full function 全功能

# How to use(v1.1.4+)

Initial with

go get github.com/tobycroft/TuuzNet

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

The old version:

```go
data, err: = Net.Post(botinfo["url"].(string) + "/get_group_info", nil, post, nil, nil)
if err != nil {
return GroupInfo {}, err
}
var ret1 GroupInfoRet
jsr: = jsoniter.ConfigCompatibleWithStandardLibrary
err = jsr.UnmarshalFromString(data, & ret1)
if err != nil {
return GroupInfo {}, err
}
if ret1.Retcode == 0 {
return ret1.Data, nil
} else {
return GroupInfo {}, errors.New(ret1.Status)
}
```

Ver+ v1.1.4

Ver 1.11.0+ Support Builder Mode

```go
get := new(Net.PostBuilder).New()
get.SetUrl(url_stats)
get.SetHeader(map[string]string{
    "user-agent": "your user agent",
    "ver":        "1.1.0",
    "platform":   "1",
})
//if you need proxy then you can use this
if proxy_on {
    auth := &proxy.Auth{
        User:     socks5_user_if_need,
        Password: socks5_password_if_have,
    }
if !need_auth {
        auth = nil
    }
    get.ProxySocks5("tcp", "127.0.0.1:10808", auth)
	//if there is no auth needed 
	//and you dont need to care about whether it will have in the future 
	//you can directly use nil
    //get.ProxySocks5("tcp", "127.0.0.1:10808", nil)
}
get.SetPostData(map[string]any{
    "tags":   "",
    "system": 28,
    "lang":   lang,
})
retbyte, err := get.PostUrlXEncode().RetBytes()

```
