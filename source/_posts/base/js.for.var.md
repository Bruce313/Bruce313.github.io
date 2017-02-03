##作用域，for不提供,var不起作用
```js
for () {
    var con = createConnecntion () ;
    doSth(function () {
        con.close();
    });
}
```
> for 没有新建作用域，var 不起作用，con的地址一样，闭包传递了地址
