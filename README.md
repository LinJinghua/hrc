# 消息驱动编程

> HTTP Reactive Client 是一个典型的消息（事件）驱动的案例。[^pmlpml]


----------


## 动机
从[Motivation for Reactive Client Extension](https://jersey.github.io/documentation/latest/rx-client.html#d0e5556)这一节中总结Reactive动机：
 - 同步方法性能差。使用同步方法处理每个请求的缺点是速度慢。因为需要顺序处理所有的独立请求，这意味着浪费资源----有时不必要阻塞线程，否则可用于一些其他工作。
 - 异步回调使代码变得复杂。(如Callback Hell)
 - Reactive Approach一定程度上克服了这两个缺点。

## Naive Approach
流程如下图[^synchronousway]。
![synchronous way](https://jersey.github.io/documentation/latest/images/rx-client-sync-approach.png)
结果：
```
> cli -m sync
[Sync Service]  |   Start  | 2017-12-29 20:40:35.3830525 +0800 CST m=+0.003985000
{
  "customers": "8675309",
  "duration": "150.1473ms"
}

{
  "destinations": "Destinations",
  "duration": "250.967ms"
}

{
  "weather": 1,
  "duration": "170.7828ms"
}

... ...

{
  "quoting": 10,
  "duration": "330.6603ms"
}

[Sync Service]  |  Finish  | 2017-12-29 20:40:40.8890264 +0800 CST m=+5.509958900
[Sync Service]  | Duration | 5.5099855s
```

## Channel 搭建基于消息的异步机制
流程如下图[^asynchronousway]。
![asynchronous way](https://jersey.github.io/documentation/latest/images/rx-client-async-approach.png)
结果：
```
> cli -m async
[Async Service] |   Start  | 2017-12-29 21:09:18.103084132 +0800 +08 m=+0.001677987
{
  "customers": "8675309",
  "duration": "150.8119ms"
}

{
  "destinations": "Destinations",
  "duration": "250.5256ms"
}

{
  "weather": 1,
  "duration": "170.4075ms"
}

... ... 

{
  "quoting": 8,
  "duration": "332.0998ms"
}

[Async Service] |  Finish  | 2017-12-29 21:09:18.843827823 +0800 +08 m=+0.742421694
[Async Service] | Duration | 740.758048ms
```

## go 异步 REST 服务协作的优势
从上两节结果可见：完成同样的任务，同步方法使用了约5.4秒，而异步REST服务协作只用了约0.74秒。异步REST服务协作的优势不言而喻：高效利用CPU资源，更快完成任务。

## 一般性的解决方案
利用 Channel 搭建基于消息的异步机制。


----------
### 附录
测试说明：
```
> go get github.com/LinJinghua/hrc
> go install github.com/LinJinghua/hrc && hrc&
> go install github.com/LinJinghua/hrc/cli
> cli -m sync
> cli -m async
```


[^pmlpml]:[理解 goroutine 的并发](http://blog.csdn.net/pmlpml/article/details/78850661#t6)

[^synchronousway]:[Time consumed to create a response for the client – synchronous way](https://jersey.github.io/documentation/latest/rx-client.html#d0e5556)

[^asynchronousway]:[Time consumed to create a response for the client – asynchronous way](https://jersey.github.io/documentation/latest/rx-client.html#d0e5556)
