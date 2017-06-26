概念
---
1. Logger
- 感知事件events
- 通知**合适**的`Appender`

2. Appender(Handler)
- 用`Layout`把事件产生的信息传递到**目的地(destination)**

3. Layout(Formatter)
- 格式化事件的信息(view, template)

4. Filter 
- 决定`Logger`是否广播事件
- 决定`Appender`是否响应事件

5. Marker
- 事件再分类

框架
---
1. java.util.logging
2. Log4j Logback tinylog 实现
3. Apache Commons Logging & SLF4J 提供**前端**抽象，解耦底层实现


配置
---
1. log4j2
- 在`classpath`下搜寻**log4j.yaml|xml|properties**作为配置


使用
----
1. 创建Logger
- 通过Logger创建Logger, 级联的创建下去
- 每个名字的Logger都是*单例的*
- 使用**ClassName**作为Logger的名字
- 默认的Filter是通过`Level`来过滤日志的
- `logp`带上打印日志的类和方法
- `logrb` 
- `entering`
- `exiting`

2. Appenders
- 通过代码或配置添加Appender
- ConsoleAppender
- FileAppender || RollingFileAppender
- Syslogappender
- SMTPAppender
- FailoverAppender 失败时转给其他Appender
- UDP -> ELK

3. Layouts
- `%xEx`占位符来写入**stack trace**