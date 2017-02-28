---
title: api_design
date: 2017-02-20 16:53:25
tags: pattern arch
---

1. /:username/abc 而不是 /:uid/abc
2. 列表中有url指示实体资源所在位置
3. 列表events用s 组成树形结构 不是一种模型就在一层上
4. 全局统一的type名称
5. 过滤: 字段，条件
6. 排序
7. 创建和更新返回最新
8. 统一的type给列表
9. 统一的错误码
10. 返回next_page 使用游标(last_tick(id/time))分页
