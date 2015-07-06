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
