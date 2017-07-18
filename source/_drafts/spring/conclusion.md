---
title: spring框架概括
tags: spring-framework
---
Bean
---
1. 方便的操作对象的属性
2. 方便的调用方法
3. 方便的获得对象的原信息
4. 递归的执行以上操作到*级联*的对象


BeanFactory
---
1. 生成Bean
2. Bean有预加载的**单例**模式和*原型**模式等
3. 通过配置来生成`BeanFactory`然后生成**Bean**

####概念
1. Singleton
2. Early Early Instance

Spring Boot
---
1. 用spring.factories来配置启动过程
中的各个节点的执行类(Delegate)
2. Initializer Listener PostProcessor


all
==
1. 抽象和模型不是被设计出来的, 是在模拟
2. 实现尽可能多的声明自己的功能(实现多接口)
3. 接口定义功能 虚拟类定义模型属性
4. 抽象类有**组合接口**和**提供默认实现**的功能