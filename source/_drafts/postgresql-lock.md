---
title: postgresql_lock
date: 2017-02-07 14:04:07
tags: postgresql;db
---

### 秒杀
1. `update tbl set a = a - 1 where a > 0`
> 有等待 没有获得锁的也要等待

2. `select * from tble where id = :id for share nowait`
> 同一个行做秒杀的话 只有一行 会有些本该秒到的没有秒到
> 可以把每一个商品做成一行

3. `update tbl set a = a - 1 where id=:id and pg_try_advisory_xact_lock(:hash)` 
> hash 可以是uid % number_of_goods


