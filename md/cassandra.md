#freatrue
 * 写入可用性
 * 无单点故障
#performace
 * 追加模式，update和insert一样，不用读文件，磁头不需要seek
 * bloom filter + memtable 
 * 操作分成几个阶段，分配给线程，用较少的线程，减少线程调度
