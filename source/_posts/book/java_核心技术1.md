----
title: java核心技术1卷
tags: 
- java
- 读书笔记
----

## 类
### getter setter好处 
1. 只要一个位置能改数据
2. 改之前可以检查数据正确性
3. 改了可以触发事件
4. 可以更改表现（view）给getter

### 静态方法只能访问静态字段
1. 静态字段用`static final`
2. 如果字段用`final` 那么初始化后*必须被设置*了 且不能更改

### 所有方法调用都是`值调用`
1. 如果改变基本类型(`int boolean`) 参数值的改变无效

```java
int p = 100;
public void changeParam(int param) {
    param *= 3;
}
changeParam(p);
println(p);//100
```
2. 如果是对象引用 那么指针被复制 可以通过指针指向的对象的方法改变对象 **然而**并不改变指针

```java
public void changeObj(Sample s) {
    s.eat();
}
```

### 初始化
1. int初始化0 boolean=false 
2. 对象分两个阶段 `声明&初始化`

```java
Object o;
o.id();//编译错误 未初始化
```
3. 对象被初始化那么属性跟着初始化 所以里面的对象就变成了`null`
> 然而使用**默认/隐藏 动作**不是一个专业程序员所应该做的: )

**坑** //也许该给坑一个图标 
```java
class Employee {
    private String name;
    public Employee(name) {
        this.name = name ? name : "";
    }
}
```
> 不该用同名 会有**访问哪一个**的问题

4. 可以有两个代码块会在构造函数之前执行

```java
class Obj {
    {
        //每次构造对象都执行
    }
    static {
        //初始化类时就会执行
        //所以会在上面之前
    }
}
```

### 类的设计
1. private所有字段
> 需求经常变但是类的使用方使用的方法变化小
> 易于调试bug

2. 显式的初始化各个字段
>  隐藏的代码/错误/默认动作 越少 就越清晰越难出现难以调试的bug

3. 归并多个字段到一个类 
> 减少类中字段过多 抽象子类 更容易修改

4. 考虑使用包的私有类 类的私有方法
5. 类的单一职责


### 继承 `is-a`关系
1. 使用超类方法
2. 覆盖超类方法
3. 如果构造过程复杂且构造参数多 可以考虑使用builder模式
4. 如果不想让子类
5. 超类对象转成子类对象要**强制类型转换**

```java
class Manager extends Employee {
}
Employee e = new Manager();
if(e instanceof Manager) {
    Manager m = (Manager)e;
}
try {
    Manager m = (Manager) e;
} catch(ClassCastException e) {
    logger.info("not a manager");
}
```

6. 抽象方法是占位用

> 和接口有啥区别?
> 

7. 