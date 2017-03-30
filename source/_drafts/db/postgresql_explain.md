postgresql explain 用法
===
执行计划中代表的意思
---
> `Scan method` on `table`(cost=`启动耗时`..`如果没有limit等 所有的行被读取消耗的时间` rows=`会影响的行` width=`所有行的平均宽度(bytes)`)
>    (actual time=`准备时间`...`实际时间` rows=`实际查询时间` loops=`这个子查询被执行的次数`)

> 输出是个树行结构 执行也是按照树的依赖来决定顺序和串并行

影响耗时的因素
---
1. 执行时间
2. 读取时间
3. 网络消耗
4. 分析语句时间（比如其中涉及的系统调用）
5. 大表和小表得到的执行计划会不同
> 影响比较大的是scan method & join method

Scan method
---
1. seq scan 
> 没有办法或者认为这是最快的方式 比如行比较少
2. index scan
> 爬树(or hash) 找到位置 `立即`取数据(random access)
3. index only scan
> 同index scan但是不用**random access**
4. bitmap [heap / index ] scan
> 同index scan 不过要中间有个bitmap的结构存储index scan的结果 然后读出符合条件的page 再`cond recheck`

join method
----
1. nested loop
> 有条件的一般作为外层表 有索引用index scan
2. hash join
> 把小表hash. scan大表再去小表中查找 hash一般都耗内存...
3. merge join
> 把两个表并行排序 然后遍历大表 查找对应的小表

影响planner的因素
---
- pg_class的一些汇总信息 

具体的方法
----
- 小表join小表再join大表 而不是让planner自己决定顺序
- 为了不让analyse不执行 可以包在begin commit中（改重要数据也是一样) 
- 尽量使用index only scan
```sql 
create index on id;
create index on (a, b)
select id from tab where a = 1
-- 并不是index only scan
```
其他
---
1. 优化器选择执行计划是要靠统计信息的 所以在**production**环境分析执行才有意义
