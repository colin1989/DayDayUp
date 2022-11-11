package CuttingRope

import "math"

// 面试题14：剪绳子
// 题目：给你一根长度为n绳子，请把绳子剪成m段（m、n都是整数，n>1并且m≥1）。
// 每段的绳子的长度记为k[0]、k[1]、……、k[m]。k[0]*k[1]*…*k[m]可能的最大乘
// 积是多少？例如当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此
// 时得到最大的乘积18。

// 动态规划(Dynamic Programming)
func cuttingRopeDP(n int) int {
	if n < 2 {
		return 0
	} else if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	}

	// 使用空间换时间,从低到高记录每段绳子切割后的最大乘积
	products := make([]int, n+1)
	// 默认0-3
	products[0] = 0
	products[1] = 1
	products[2] = 2
	products[3] = 3
	for i := 4; i <= n; i++ {
		// 切割一次,在从对应的记录里获取乘积
		for j := 1; j <= i/2; j++ {
			product := products[j] * products[i-j]
			if products[i] < product {
				products[i] = product
			}
		}
	}
	return products[n]
}

// 贪婪算法(Greedy Algorithm)
func cuttingRopeGA(n int) int {
	if n < 2 {
		return 0
	} else if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	}

	// 尽可能剪掉长度为3的绳子
	timesLen3 := n / 3
	// 当剪完所有长度3的绳子后,剩余1时,应该留一段4的绳子
	// 因为 2*2>3*1
	if n-timesLen3*3 == 1 {
		timesLen3 -= 1
	}
	timesLen2 := (n - timesLen3*3) / 2

	return int(math.Pow(3, float64(timesLen3)) * math.Pow(2, float64(timesLen2)))
}
