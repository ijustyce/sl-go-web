# sl-go-web
go 后端开发的框架太多了，不像 java，基本就是 spring 全家桶了。  
go 后端开发框架可以看：[go-web-framework-stars](https://github.com/mingrammer/go-web-framework-stars)  
性能评测可以看：[go-web-framework-benchmark](https://github.com/smallnest/go-web-framework-benchmark)  

目前可选的框架有：gin、beego、echo、fiber、go-zero等  
其中，echo 和 fiber 基于 fasthttp，fasthttp 仅支持 http 1.1 不支持 http 1.0 和 http 2.0，这是它的局限，
但性能确实像它的名字一样，快，在我的测试中，gin 可以跑到 6.81万 qps，而 fiber 可以达到 8.19w qps，
spring boot 只有 2.42w qps，单从效率来看，spring boot 确实弱爆了。  
go-zero 官方的文档看，性能和 beego 以及 gin 差不多，beego 是国人开发，作者已经创业，维护不如 gin，维护上，
个人感觉 fiber 应该是第一，beego 官网已经打不开，文档也打不开，基本半残废，echo 也是基于 fasthttp，
不过后面的规划是移除它且不再支持 fasthttp，个人觉得，值的关注的，依次为：gin、go-zero、fiber、echo  

至于 json 处理，github.com/goccy/go-json 就好，它兼容标准的 encoding/json，且性能更好。  

正如前面所说，可以使用的框架有：gin、go-zero、fiber、echo 等，在尚不清楚未来走向的情况下，使用 go 开发后端，更好的
办法是自己封装一层，使得应用程序感知不到自己使用的是哪个框架，这便是该项目的用途。

## todo
- [x] bind header
- [x] set header
- [x] get cookie
- [x] set cookie