基于zinx与刘丹冰的视频（BV1wE411d7th），学习zinx的设计从而了解Go在服务器开发的用法。

# 连接

​	为每一个Conn再封装一层，拥有启动连接，停止，获取当前连接Conn、连接Id、客户端ip，以及读取或发送方法，其实有点类似于Netty的SocketChannel

