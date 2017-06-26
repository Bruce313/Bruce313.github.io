---
title: job\_log
date: 2016-04-18 15:15:44
tags:
---
2016/04/18
---
- linux
*ENOSPEC* no space空间不足(/tmp比如)


4-24
---
1. node发送请求
```
ECONNRESET
```
怀疑是网络不对
通过`curl`可以
使用`tcpdump`监控那里不一样
发现服务没有接受到请求
> url里的中文必须`encodeURIComponent`不然客户端连发都发不出去

## 关于写日志
> One reason is that it is easy to get lost in details of complicated data structures and control flow; we find stepping through a program less productive than thinking harder and adding output statements and self-checking code at critical places


## Log4j
概念
---
- LoggerConfig
- LoggerContext
- LoggerContext.Configuration (only on active)
- LoggerContext.Appenders
- LoggerContext.Context-WideFilters
- LoggerManager
```java
Logger log = LoggerManager.getLogger("name");
Logger logO = LoggerManager.getLogger("name");
assert log == logO; //same obj
```


过程
---
1. `Logger`根据名字创建于`LoggerManager`
2. `Logger`关联一个`LoggerConfig`
3. 首先找到name相同的配置给`LoggerConfig`
4. 否则使用**parent package**的配置
5. 最终使用**root LoggerConfig**

特性
---
1. 级联(Hierarchy) Logger
> Logger不再有级联关系，关系转到`LoggerConfig`

2. 

实践
---
1. **ClassName**作为Logger的名字