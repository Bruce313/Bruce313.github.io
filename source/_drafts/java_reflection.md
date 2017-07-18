---
title: java反射
tags: java;advanced
---
- 性能下降：优化失效了
- 每个Object都会有.Class对象来标识其元信息&创建同类对象

得到`Class`对象
----
1. `instance.getClass()`
2. `Class.forName('')`
3. `Class c = Double.TYPE, Class v = Void.TYPE; `
4. `Class c = Class.getSuperClass();`

`Class`常用方法
---
1. `getModifiers()` 获得set
2. `getAnnotations()`
3. `getGenericInterfaces()`
> [Class的文档](https://docs.oracle.com/javase/8/docs/api/java/lang/Class.html)

