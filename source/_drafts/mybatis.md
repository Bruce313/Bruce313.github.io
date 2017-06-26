## 设置
- cacheEnabled: default true
> 表的列类型变化了会产生错误

## typeHandlers
> 所有的数据库字段都通过`typeHandler`转换成了Java字段

## 参数
- timeout
- useCache 有二级缓存
```xml
<select timeout="true" flushCache = "false" useCache = "true">
    select * from user
</select>
```
- `useGeneratedKeys` keyProperty 和`useGeneratadKeys`一起使用 
> 拿到自动生成的值, 多个用逗号分隔, 多行插入也能拿到

## 属性映射
```xml
<resultMap id="userResultMap" type="User">
  <id property="id" column="user_id" />
  <result property="username" column="user_name"/>
  <result property="password" column="hashed_password"/>
  <assocation property="propA" resultMap="另一个ResultMap">
  </assocation>
</resultMap>
```

## 缓存
```xml
<cache flushInteval = "1000" size="1024" readOnly = "true"/>
```
- <cache type="实现了Cache接口的类"/> 自定义缓存


