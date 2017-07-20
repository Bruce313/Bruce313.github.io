---
title: 设计模式：适配器模式
p: dp/adapter
date: 2017-07-20 10:28:07
tags:
- design pattern
- 工程
category:
- dp 
---
定义
---
> 适配器模式(Adapter Pattern)：将一个接口转换成客户希望的另一个接口，使接口不兼容的那些类可以一起工作，其别名为包装器(Wrapper)

形象
---
- 手机充电口

参与角色
---
- **Adaptee**
- **AbstractAdapter** 固定的一端确定
- **ConcreteAdapter** 
- **Target** (可以使调用方，也可以是被调用方)

案例
---
- springmvc中对handler的处理
```java
//AbstractAdapter
public interface HandlerAdapter {
	boolean supports(Object handler);
	ModelAndView handle(HttpServletRequest request, HttpServletResponse response, Object handler) throws Exception;
}
//Adaptee A
public interface Controller {
	ModelAndView handleRequest(HttpServletRequest, HttpServletResponse);
}
//ConcreteAdapter A
public class HttpRequestHandlerAdapter implements HandlerAdapter {

	@Override
	public boolean supports(Object handler) {
		return (handler instanceof HttpRequestHandler);
	}

	@Override
	public ModelAndView handle(HttpServletRequest request, HttpServletResponse response, Object handler)
			throws Exception {

		((HttpRequestHandler) handler).handleRequest(request, response);
		return null;
	}
}
//Adaptee B
interface Servlet {
    public void service(ServletRequest req, ServletResponse res)
}
//ConcreteAdapter B
public class SimpleServletHandlerAdapter implements HandlerAdapter {

	@Override
	public boolean supports(Object handler) {
		return (handler instanceof Servlet);
	}

	@Override
	public ModelAndView handle(HttpServletRequest request, HttpServletResponse response, Object handler)
			throws Exception {

		((Servlet) handler).service(request, response);
		return null;
	}
}
```
