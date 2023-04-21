> 参考极客时间-消息队列高手课 [13讲](https://time.geekbang.org/column/article/119988) [加餐](https://time.geekbang.org/column/article/126504)

## 问题

张大爷在胡同口遇见了李大爷  

张大爷：“吃了没，您呐？”  
李大爷：“刚吃。”  
李大爷：“您这，嘛去？”  
张大爷：“嗨，没事儿溜溜弯儿。”  
李大爷：“有空家里坐坐啊。”  
张大爷：“回头去给老太太请安。”

假设张大爷为服务端，李大爷为客户端，俩儿大爷每次遇见都会有这三组对
话，请你以最快的方式实现俩儿大爷遇见十万次，也就是三十万次对话

## 实现方式

基本实现 -> communication.go  
耗时：8.0s

## 优化

1. 写  
   + 耗时：4.0s
   + communication_reduce_write.go  
   + 把三个写操作变为一个写操作，并且所有的写操作通过一个线程来执行
2. <font color=red>批量读</font>
   + 耗时：1.8s
   + communication_reduce_read.go
   + 把多次读一个变为一次多多个
3. <font color=red>批量写</font>
   + 370ms
   + communication_batch_write.go
   + 把一次写一个变为一次写多个