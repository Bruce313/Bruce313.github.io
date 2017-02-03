pg 查询优化
===
Index Only Scan
---
- pg的索引和数据分离
- 根据索引查询数据后 未索引的列会`random access`的读数据区
- 如果只读带有索引的列 就会快些
- `select x from tab where y = 1` create index on (y, x)会触发index only scan

确保索引使用
---
- 使用真实数据测试
- analyze always
- 强制使用索引 分析消耗
- 汇总数据可能误导planner pg_class

方式 
---
- 用表达式创建索引 `create index on tab(func(col))`
- 创建view 对读做优化(额 其实建索引也是牺牲写性能)
- with大法有时会阻止优化器 所以试试子查询
- 大多数时候`select a, b, c`要比`select *`要快
- prepared statement
- 执行`cacuum analyze`更新统计信息 然后`create index`会更智能
- join要比subquery(select a, (select b from c) as c) 要对优化器友好
- `inner join`要比`left join` 要对优化器友好
- 避免pattern match把%放在前面`like '%abc'`, 不会使用索引
- 使用`部分索引(partical index)`排除null等
- 用`order by limit`代替`max min` 因为pg对aggr(尤其count)不友好
- 使用汇总字段 用冗余对读优化
- 使用冗余 把外表直接写入表 减少join
- 使用limit影响join
