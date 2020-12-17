package main

import "fmt"

// arr    : 用来存储待排序的元素
// n      : 代表元素数量
// output : 输出过程中的调试信息
// DEBUG = 1 开启调试信息
// DEBUG = 0 关闭调试信息

var (
	arr []int
	n   int
)

func init() {

	arr = []int{
		4, 6, 1, 88, 55, 45, 70, 27, 38,
	}

	n = len(arr)

}

func main() {
	quick_sort(arr, 0, n-1)
	fmt.Print(arr)
}

// 快速排序：对 arr 中 l 到 r 位进行排序
// arr : 待排序数组
// l   : 待排序区间起始坐标
// r   : 待排序区间结束坐标
func quick_sort(arr []int, l, r int) {
	// 递归的边界条件是区间中只有一个元素
	// x  : 记录从前向后扫描的位置
	// y  : 记录从后向前扫描的位置
	// z  : 基准值，选择待排序区间的第一个元素
	// for 循环中是 partition 过程
	// 每一轮，先从后向前扫，再从前向后扫
	if l >= r {
		return
	}
	x := l
	y := r
	z := arr[l]
	for x < y {
		for x < y && arr[y] >= z {
			y--
		}
		if x < y {
			arr[x] = arr[y]
			x++
		}
		for x < y && arr[x] <= z {
			x++
		}
		if x < y {
			arr[y] = arr[x]
			y--
		}
	}
	// 将基准值 z 放入其正确位置数组的 x 位
	// 其实，此时 x==y，所以写成 arr[y] = z 也行
	// 再分别对左右区间，进行快速排序
	arr[x] = z
	output(l, x, r)
	quick_sort(arr, l, x-1)
	quick_sort(arr, x+1, r)
	return
}

func output(l int, x int, r int) {
	fmt.Printf("\n待排序区间范围 【%d,%d】\n", l, r)
	fmt.Printf("基准值：%d\n", arr[x])

	for i := 1; i < l; i++ {
		fmt.Printf("%d ", arr[i])
	}

	fmt.Printf("[")

	for i := 1; i <= r; i++ {
		fmt.Printf("%d ",arr[i])
	}

	fmt.Printf("]")

	for i := r + 1; i <= n-1; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
	fmt.Scanln()
}
