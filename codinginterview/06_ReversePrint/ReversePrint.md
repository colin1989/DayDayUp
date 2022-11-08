# 反转输出

## 实现了4种方案
1. 递归+append
2. 递归计算出深度值,在最深处创建一个大小刚好合适的 `slice`,通过下标赋值
3. 反转链表
4. 遍历种`append`,后续通过 `golang` 语言特性的值交换

## 分析
使用命令
``go
go test -bench . -run none -benchmem -cpuprofile cpuprofile.out -memprofile memprofile.out
go tool pprof -http=":8081" .\cpuprofile.out
``
减少 `growslice` 的触发可以提升性能


## 复杂度
* 时间复杂度 O(N)
* 空间复杂度 O(N)