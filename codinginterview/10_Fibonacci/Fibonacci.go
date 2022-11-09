package Fibonacci

// 面试题10：斐波那契数列
// 题目：写一个函数，输入n，求斐波那契（Fibonacci）数列的第n项。
func fibonacci1(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci1(n-1) + fibonacci1(n-2)
}

func fibonacci2(n int) int {
	if n < 2 {
		return n
	}
	f0, f1, fn := 0, 1, 0
	for i := 2; i <= n; i++ {
		fn = f0 + f1
		f0 = f1
		f1 = fn
	}
	return fn
}

//写一个函数，输入 n ，求斐波那契（Fibonacci）数列的第 n 项（即 F(N)）。斐波那契数列的定义如下：
//
//F(0) = 0, 	F(1) = 1
//F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
//斐波那契数列由 0 和 1 开始，之后的斐波那契数就是由之前的两数相加而得出。
//
//答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
func fibonacci3(n int) int {
	if n < 2 {
		return n
	}
	f0, f1, fn, mod := 0, 1, 0, int(1e9+7)
	for i := 2; i <= n; i++ {
		fn = (f0 + f1) % mod
		f0, f1 = f1, fn
	}
	return fn
}
