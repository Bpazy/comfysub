# comfysub 
畅想订阅（Comfortable subscription），适用于 Shadowsocks, V2ray 等订阅间互相转换。

## 支持状况
1. [x] ShadowsocksD -> Shadowsocks(SIP002)

## 使用
1. 在 releases 中下载；
2. 运行：
```
$ chmod +x ...
$ ./gfw-converter -port :8080 
$ curl http://127.0.0.1:8080/ssd2ss?url=http://airport.org/ssd-subscription
```
## 附录
1. [SIP002 规范](https://shadowsocks.org/en/spec/SIP002-URI-Scheme.html)
2. [v2rayN 订阅功能说明](https://github.com/2dust/v2rayN/wiki/%E8%AE%A2%E9%98%85%E5%8A%9F%E8%83%BD%E8%AF%B4%E6%98%8E)
3. [ShadowsocksD HTTP订阅协定](https://github.com/TheCGDF/SSD-Windows/wiki/HTTP%E8%AE%A2%E9%98%85%E5%8D%8F%E5%AE%9A)
