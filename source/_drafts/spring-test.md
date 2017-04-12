```java
@ContextConfiguration //用java/xml配置上下文
@ActiveProfiles
@Test
```

- spring test framework缓存ApplicationContext `siglton`的bean可能会被*污染*(dirtied) 要用`@DirtiesContext`
- `@TestExcutionListeners`添加测试执行监听器s
- 默认对数据库的操作会`@Rollback(true)` 可以用`@Commit / @Rollback(false)`标明写入到数据库

### 数据库相关
```java
@Sql // 执行sql语句在测试之前     
@Transactional //在事务中测试
ResourceDatabasePopulator //用语言执行语句文件
```

### 控制是否执行
```java
@Ignore //不执行
@IfProfileValue //执行如果匹配环境
```

### 控制执行时间
```java
@Timed //必须在指定时间内执行完成
@Repeat //执行几次
```

Integer equals ==
