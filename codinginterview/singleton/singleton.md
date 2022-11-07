# 单例模式

单例模式实现比较简单,主要注意并发时的竞争

`singleton1`在初始化的时候直接生成单例

`singleton2`触发时通过`sync.Once`生成单例