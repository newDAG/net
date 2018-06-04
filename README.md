# net
newDAG 轻量级网络库

在不依赖于其他复杂核心库的情况下，可进行单独测试

使用内存Peer来模拟:

1.节点网络连接及收发信息情况

2.事件及区块的同步情况，

3.RPC的处理情况



使用方法:

go test -v -run TestNetworkTransport_Sync 

你也可换成其他测试文件中的测试函数