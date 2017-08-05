---
title: spring-config-classes
p: source/_posts/spring/spring-config-classes
date: 2017-07-21 14:39:37
tags:
- spring

---
> 通过`getBean(Class<?> clazz)`来加载配置
配置通过定义一些类，来识别固定位置的配置
`@Bean`或继承具体类来定义bean -> 拿到BeanDefinition
这样就通过类型达到插入固定配置位置的目的

这样的类有:
- `CorsConfiguration`
- `ContentNegotiationManager`
- UserRoleAuthorizationInterceptor 鉴权