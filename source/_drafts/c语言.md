---
title: c语言
date: 2016-04-18 12:13:34
tags: lang
---

- flag, 二进制, 用来标识
1.用作多个is_\* 字段 
```c
CLIENT_MULTI = (1 << 3);
void multiCommand(client *c) {
    if (c->flags & CLIENT_MULTI) {
        addReplyError(c,"MULTI calls can not be nested");
        return;
    }
    c->flags |= CLIENT_MULTI;
    addReply(c,shared.ok);
}
```
