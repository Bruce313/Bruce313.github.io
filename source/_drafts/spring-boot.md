---
title: spring-boot-code
tags: spring-boot;source
---
> The second class-level annotation is @EnableAutoConfiguration. This annotation tells Spring Boot to **“guess”** how you will want to configure Spring, based on the jar dependencies that you have added. Since spring-boot-starter-web added Tomcat and Spring MVC, the auto-configuration will assume that you are developing a `web` application and setup Spring accordingly.

> 超类中定义空函数，子类修改这个值
这样超类定义行为模式，子类定义具体实现
```java
 public void refresh() throws BeansException, IllegalStateException {
        Object var1 = this.startupShutdownMonitor;
        synchronized(this.startupShutdownMonitor) {
            this.prepareRefresh();
            ConfigurableListableBeanFactory beanFactory = this.obtainFreshBeanFactory();
            this.prepareBeanFactory(beanFactory);

            try {
                this.postProcessBeanFactory(beanFactory);
                this.invokeBeanFactoryPostProcessors(beanFactory);
                this.registerBeanPostProcessors(beanFactory);
                this.initMessageSource();
                this.initApplicationEventMulticaster();
                this.onRefresh();
                this.registerListeners();
                this.finishBeanFactoryInitialization(beanFactory);
                this.finishRefresh();
            } catch (BeansException var9) {
                if(this.logger.isWarnEnabled()) {
                    this.logger.warn("Exception encountered during context initialization - cancelling refresh attempt: " + var9);
                }

                this.destroyBeans();
                this.cancelRefresh(var9);
                throw var9;
            } finally {
                this.resetCommonCaches();
            }

        }
    }
	protected void onRefresh() throws BeansException {
    }
```


**不传递参数的，就是事件 否则就是接口调用**



ApplicationContext Listeners
----
```java
0 = "org.springframework.boot.ClearCachesApplicationListener"
1 = "org.springframework.boot.builder.ParentContextCloserApplicationListener"
2 = "org.springframework.boot.context.FileEncodingApplicationListener"
3 = "org.springframework.boot.context.config.AnsiOutputApplicationListener"
4 = "org.springframework.boot.context.config.ConfigFileApplicationListener"
5 = "org.springframework.boot.context.config.DelegatingApplicationListener"
6 = "org.springframework.boot.liquibase.LiquibaseServiceLocatorApplicationListener"
7 = "org.springframework.boot.logging.ClasspathLoggingApplicationListener"
8 = "org.springframework.boot.logging.LoggingApplicationListener"
9 = "org.springframework.boot.autoconfigure.BackgroundPreinitializer"
10 = "org.springframework.cloud.bootstrap.BootstrapApplicationListener"
11 = "org.springframework.cloud.bootstrap.LoggingSystemShutdownListener"
12 = "org.springframework.cloud.context.restart.RestartListener"
```

从`META-INF/spring.factories`加载factory是
`SpringFactoriesLoader`的默认路径
所以factory是程序起点？