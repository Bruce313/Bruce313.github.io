---
title: junit
date: 2017-02-28 18:35:07
tags: 
- 单元测试
- 学习
---
1. assert : 第一个参数是消息
2. RunWith 自定义TestRunner
3. Ignore
4. Suite 组套测试 分组
5. Timeout 要求性能
6. expected 测试异常
7. assertThat(value, Matcher)

## 概念
1. Test
2. Runner
3. Suite分类，Category贴标签
4. Assert 
5. Rule 多个测试可重用的资源
    - ErrorCollector
    - Verifier
    - TestWacher
    - TimeoutRule
    - ExpectedException
6. Matcher
7. Assumption & Ignore & Runner filter来过滤测试

## Runner
1. Parameterized