package main

import (
	"fmt"
	"os"

)
func main(){
	arr:=[]int{43,67,98,23,56,13,74,68,96,34,22}
	for i:=0;i<len(arr);i++{
		fmt.Printf("排名第 %d 位的元素：%d\n",i,quickSelect(arr,0,len(arr)-1,i))
		os.Exit(0)
	}
}

// arr  : 待查找数组
// l-r : 待查找区间
// k : 待查找元素的排名
// 在 arr 数组的 l 到 r 区间内，查找排名为 k 的元素
func quickSelect(arr []int,l,r,k int)int{
	// 首先选取基准值，完成 partition 操作
	x:=l
	y:=r
	z:=arr[l]
	for x<y {
		fmt.Println(x,y)
		for x<y && arr[y]>=z{
			y--
		}
		if x<y {
			x++
			arr[x] = arr[y]
		}
		for x<y && arr[x]<=z {
			x++
		}
		if x<y {
			y--
			arr[y] = arr[x]
		}
	}
	arr[x]=z
	// ind 为当前基准值的排名
    // 用基准值的排名与 k 做比较
    // 如果相等，则为基准值
    // 如果 ind > k，在前半部分查找排名第 k 位的元素
	// 如果 ind < k, 在后半部分查找排名第 k - ind 位的元素
	ind:=x-l+1
	if ind ==k{
		return arr[x]
	}
	if ind>k {
		return quickSelect(arr,l,x-1,k)
	}
	return quickSelect(arr,x+1,r,k-ind)
}