---
title: regexp
date: 2016-04-19 16:27:38
tags:
---
1. sds.Trim
> `/abc$/`不能匹配"abcdfeabc"因为`$`代表换行
> `/^abc/`能, 因为开始总有*new line*
