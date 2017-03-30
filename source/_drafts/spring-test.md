```java
@ContextConfiguration //用java/xml配置上下文
@ActiveProfiles
@Test
```

- spring test framework缓存ApplicationContext `siglton`的bean可能会被*污染*(dirtied) 要用`@DirtiesContext`
- `@TestExcutionListeners`添加测试执行监听器s
- 默认对数据库的操作会`@Rollback(true)` 可以用`@Commit / @Rollback(false)`标明写入到数据库

### @Sql
> 执行sql语句在测试之前

