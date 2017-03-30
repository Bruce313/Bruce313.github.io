---
title: redis源码阅读
date: 2016-04-18 12:08:35
tags: redis; source
---
##arch
- 对读优化
1. 各种count/free/length冗余
- 不适合计算任务，比如key \*, 作为缓存，适合快速读取
1. 单进程，非抢占式
2. 定时任务可能会有延迟
