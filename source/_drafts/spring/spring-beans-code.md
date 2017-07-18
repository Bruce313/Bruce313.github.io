---
title: spring bean 代码
tags: spring;源码
---

```java
PropertyDescriptor(PropertyDescriptor x, PropertyDescriptor y) {
	super(x,y);

	if (y.baseName != null) {
		baseName = y.baseName;
	} else {
		baseName = x.baseName;
	}
}
```
1. 以**同类对象**作为参数的构造函数
2. Constructor不一定public

```java
public class PropertyDescriptor extends FeatureDescriptor {
    private Reference<? extends Class<?>> propertyTypeRef;
}
```
1. Reference是对原对象的引用 
2. 虚拟类可以集成具体类

```java
Assert.isTrue(!(result instanceof Optional), "Multi-level Optional usage not supported");
```
1. 底层需要**断言编程**

```java
public static <T> T instantiateClass(Class<T> clazz) throws BeanInstantiationException {
		Assert.notNull(clazz, "Class must not be null");
		if (clazz.isInterface()) {
			throw new BeanInstantiationException(clazz, "Specified class is an interface");
		}
		try {
			return instantiateClass(clazz.getDeclaredConstructor());
		}
		catch (NoSuchMethodException ex) {
			throw new BeanInstantiationException(clazz, "No default constructor found", ex);
		}
	}
```
1. 将类似于断言和抛出错误的代码放入utils

```java
		public PropertyEditor getDefaultEditor(Class<?> requiredType) {
		if (!this.defaultEditorsActive) {
			return null;
		}
		if (this.overriddenDefaultEditors != null) {
			PropertyEditor editor = this.overriddenDefaultEditors.get(requiredType);
			if (editor != null) {
				return editor;
			}
		}
		if (this.defaultEditors == null) {
			createDefaultEditors();
		}
		return this.defaultEditors.get(requiredType);
	}

	/**
	 * Actually register the default editors for this registry instance.
	 */
	private void createDefaultEditors() {
	}
```
1. 工作方法private, 修饰方法public

模式
----
1. 大量的Factory
2. Delegate
	- TypeConverter: `TypeConverterDelegater`
3. 接口定义**功能**, 虚拟类定义**模型**, 然后提供内置实现*support*


```java
public interface HierarchicalBeanFactory extends BeanFactory {

	/**
	 * Return the parent bean factory, or {@code null} if there is none.
	 */
	BeanFactory getParentBeanFactory();

	/**
	 * Return whether the local bean factory contains a bean of the given name,
	 * ignoring beans defined in ancestor contexts.
	 * <p>This is an alternative to {@code containsBean}, ignoring a bean
	 * of the given name from an ancestor bean factory.
	 * @param name the name of the bean to query
	 * @return whether a bean with the given name is defined in the local factory
	 * @see BeanFactory#containsBean
	 */
	boolean containsLocalBean(String name);

}
```
1. 拆分接口不一定是平级的
2. 对这个功能一定还有别方法

```java
@Override
public final T getObject() throws Exception {
	if (isSingleton()) {
		return (this.initialized ? this.singletonInstance : getEarlySingletonInstance());
	}
	else {
		return createInstance();
	}
}
protected abstract T createInstance() throws Exception;
```
1. 虚拟类定义行为模式 具体类定义实现



all
===
1. 使用boolean