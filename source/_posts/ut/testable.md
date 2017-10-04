---
title: testable_code
p: ut/testable
date: 2017-08-05 17:37:45
tags:
- unit test
---
概念
---
> 集成测试对用户的交互行为感兴趣，而单元测试往往仅专注于一小段代码

> 软件IC是人们在讨论可复用性和基于组件的开发时最喜欢使用的比喻。意思是软件组件应该就像集成电路一样进行组合，这只有在你使用的组件已知是可靠地时候才能行之有效。
　　芯片在设计时就考虑了测试——不只是在工厂，在安装时，而且也是在部署现场进行测试。更加复杂的芯片和系统可能还拥有完整的Built-In-SelfTest（BIST）特性，用于在内部运行某种基础级的诊断；或是拥有Test Access Mechanism（TAM），用以提供一种测试装备，允许外部环境提供激励，并收集来自芯片的响应。

untestable的代码
---
- 隐藏的输入
- 依赖过多，构建麻烦
- 耦合度过高
- 非单一职责
- 显式的依赖DB NETWORK 未测试的类等(应该依赖接口或者mock)

验收标准
---
- 覆盖率高(路径被遍历)
- 圈复杂度低
- 每个输入有唯一的输出
- 过去的系统状态和变量可见，或在运行中可查询(例如：事务日志)
- 软件的可测试性特征主要表现是设立观察点、控制点、观察装置、驱动装置、隔离装置

pricinple
---
- 设计之初就考虑了`测试模式`
- 承认单元测试不是万能的，只保证部分正确
- 尽量少的依赖
> First rule of writing testable code to me is having a good dependency management within the code base, or at least having the least possible dependencies in your object/module relations
- 执行路径过多(每个执行路径是不同的代码，都该有个测试与之对应)
> 应该由单个方法来决定if的条件
> 嵌套的panduan 
> == null 
> \> 2

- 可控性指系统的状态可受外部控制改变，而不是由内部模块自发的完成
> 测试环境不可以等待每天晚上1点这个时刻
- 可观测性
> 再次，没有必要的描述如何判断哪些文件/数据被GC掉了。无法观测到执行结果集带来的后果是无法精确的预期测试结果。
- 参数的热设定
- 系统使用信息统计


NOT testable
---
1. 包含不确定的因素
```java
public static string GetTimeOfDay()
{
    DateTime time = DateTime.Now;
    if (time.Hour >= 0 && time.Hour < 6)
    {
        return "Night";
    }
    if (time.Hour >= 6 && time.Hour < 12)
    {
        return "Morning";
    }
    if (time.Hour >= 12 && time.Hour < 18)
    {
        return "Afternoon";
    }
    return "Evening";
}
```
- 依赖了系统时间
- 隐藏了输入（依赖）
- 相同测试不同输出

测什么
---
> CORRECT

实现
---
1. 完全的TDD
2. 实现-> 测试 -> 重构以更testable -> 更多测试 -> 重构测试


todo
---
- 监视圈复杂度