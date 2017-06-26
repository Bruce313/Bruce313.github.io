1. @ResponseStatus 制定**业务错误**代码

doService 
--
> Servelet处理请求入口

- `FlashMap`用户在不同请求中共享信息, 比如Redirect
- getHandlerChain
- 304
- preHandle
- 

```java
if (mappedHandler == null || mappedHandler.getHandler() == null) {
					noHandlerFound(processedRequest, response);
					return;
				}
//错误处理也可以函数
```

```java
	boolean isGet = "GET".equals(method);
	if (isGet || "HEAD".equals(method)) {
		long lastModified = ha.getLastModified(request, mappedHandler.getHandler());
		if (logger.isDebugEnabled()) {
			logger.debug("Last-Modified value for [" + getRequestUri(request) + "] is: " + lastModified);
		}
		if (new ServletWebRequest(request, response).checkNotModified(lastModified) && isGet) {
			return;
		}
	}
		//处理304
```

> spring mvc 是个框架
框架只写一些依赖的接口和适配器
怎么加载，怎么注入依赖，是使用者去做


> 方法是最简单的接口／一个方法的接口
接口是方法簇