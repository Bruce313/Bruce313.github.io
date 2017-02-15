1. 增加汇总字段 各种count
2. 需要外联找到的字段 比如uid或者uname 写入常用这些字段的表 再限制uname的修改
3. unique来避免 exists(select from tbl)
