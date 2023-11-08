这里记录了我学习 《气七天用Go从零实现系列之web框架》

设计Context
- `http.ResponseWriter`和 `http.Request` 对象的接口颗粒度太细
- 每次都需要设置 `Content-Type`

- 对 `http.ResponseWriter` `http.Request` 做封装 
- Context随着每次请求的出现而产生，请求的结束而销毁，和当前请求相关的信息都应由Context承载
- 设计 Context 结构，扩展性和复杂性留在了**内部**，而对外简化了接口。路由的处理函数，以及将要实现的中间件，参数都统一使用 Context 实例， Context 就像一次会话的百宝箱，可以找到任何东西。

![](D:\Go_WorkSpace\Gee_webFrame\前缀树.png)