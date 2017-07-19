---
title: spring mvc处理http请求的过程
date: 2017-07-18 17:49:29
tags:
- spring
- source
category:
- java
- spring
- webmvc
---
> 介绍servlet
> 实际上是遵循了[servlet](/java/servlet/)规范

1. doPost / doGet 调用`DispatcherServlet.doService`
2. request可以存取attribute, 存储原有attrs
3. 设置常用attrs
	- context
	- localeResolver **修改点**
	- themeResolver **修改点**
	- 处理并设置`FlashMap`([FlashMap](/java/spring/webmvc/spring-flashmap/)是不同请求间共享数据的方式)
4. 处理multipart 
	- 使用**multipartResovler**处理
	- 如果错误写入错误到attrs(attrs是**MultiValueMap**)

5. 在`List<HandlerMapping>`中找第一个`HandlerExecutionChain` **修改点**
6. 如果是Get请求, 处理302
	- ETag
	- lastModified
	- 设置302
	- 写回`Last-Modified`
7. 请求`SimpleUrlHandlerMapping`(**求改点**)得到`HandlerExcutionChain`
	- 如果是handler是String，那么去context里拿到Bean作为handler
	- 在Mapping中获得匹配的(`Interceptor.matches`)来加入Chain
8. 分配请求到`HandlerExecutionChain`
	- 调用`HandlerInterceptor.preHandle`, 一次调用，有false结束
	- 调用`HandlerInterceptor.handle`, 获得MoSimpleControllerHandlerAdapterdelAndView; [handle内部](/java/spring/webmvc/handlerSimpleControllerHandlerAdapter)
	- 调用`HandlerInterceptor.postHandle`
> 常用实现 `SimpleControllerHandlerAdapter`
> 每个Handler的实现自己提供**Adapter**, 因为handler是个object



9. 如果发生`Exception | Throwable`，调用`preHandle`已经返回true的`Interceptor.triggerAfterCompletion` **修改点**
10. 清理multipart **修改点**

其他
---
- [异步处理](/java/spring/webmvc/async)