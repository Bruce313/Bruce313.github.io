#start-dfs.sh
    **启动不起来**
    **slave也启动datanode**
    > 需要namenode format
    > 需要core-site.xml
        ```xml
        <property>
            <name>fs.defaultFS</name>
            <value>hdfs://namenode:8020</value>
        </property>
        ```
#同一时间只有一个*writer*对一个文件
#chrome crash while download
 > rm extensions
