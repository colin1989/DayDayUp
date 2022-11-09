package Fibonacci

// 一只青蛙一次可以跳上一级台阶，也可以跳上二级台阶……，也可以跳n级，求该青蛙跳上一个n级的台阶总共需要多少种跳法。
// 解法 f(n) = f(n-1) + f(n-2) + ... + f(1) + f(0)
//	f(n-1) = f(n-2) + ... + f(1) + f(0)
// f(n) = f(n-1) + f(n-1)
// f(n) = f(n-1)*2
// 2^n-1
// 跟斐波那契数列差不多
func JumpFloor(n int) int {
	if n < 1 {
		return 0
	}
	return 1 << (n - 1)
}
