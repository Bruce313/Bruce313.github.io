1. 通过where子句缩小范围，遇到大量计算可以先用简单计算比如> <去判断，然后计算
  > 计算离你最近的10个酒店
  > order by TJJKDJFKDJF(lat, long) limit 10
  __可以__ **where lat > a and long > b order by**
