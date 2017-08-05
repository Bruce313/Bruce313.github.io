---
title: controller-handle
p: spring/mvc/controller-handle
date: 2017-07-20 11:11:14
tags:
- spring
- mvc
categroy:
- java
- spring
---

![总图]()
![mvc图]()
![处理流程图]

流程
---
- 处理`OPTIONS`请求 allowHeader**可配置**
- 检查Method是否被支持
- 检查[session](/java/spring/session)
- 处理[Cache-Controller](/common/http/cache) **可配置**
- 处理[Vary](/common/http/cache) **可配置**
- handleRequestInternal
	+ 

TODO:
- Session 何时多个线程同步一个session
