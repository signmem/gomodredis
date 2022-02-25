# gomodredis  

# comment 

> 用于对 redis 进行压测  
> 根据配置文件生成 "conncurrency" 并发写入 redis
> 配置文件 maxactive 控制了最大并发
> 每个主机名 value 为当前写入的 timestamp     



# http heal check

> 客户端可以通过下面 api 获取一些监控数据
> curl http://localhost:8081/health (返回 ok 即程序正常)  
