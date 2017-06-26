## 遇到问题
1. 升级gradle
2. 关闭daemon `gradle --stop`

## 概念
1. project
2. task
3. DAG `gradle.taskGraph.hasTask()`访问DAG 
4. java plugin: source set可定义要操作的源文件地址

## 参数
- -q    *quite*
- --profile 保存构建日志 自动构建有用
- -i -d INFO/DEBUG级别的日志

## task

```groovy
task hello() {
    println 'hello'
}

task taskDep(dependsOn: hello) {
    println 'hello after'
}
```

## 常用任务
1. clean (rebuild?)
2. assemble : build - ut
3. check

