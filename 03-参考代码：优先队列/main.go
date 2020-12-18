package main

import (
	"errors"
	"fmt"

)


var (
	n    int
	op   int
	val  int
	loop = true
	q    *Heap
)

/*
* 下面我会用小顶堆，演示一个优先队列
* 你可以自行改成大顶堆
* */
type Heap struct {
	// 存储数据的数组, 默认每个元素的类型都是整型
	data []int
	// 现有元素数量
	n int
	// 最大容量
	size int
}

func main() {
	fmt.Println("请输入size:")
	fmt.Scanln(&n)
	q = newHeap(n)

	fmt.Printf("优先队列结构学习，请输入操作命令(0--6)：\n")
	fmt.Printf("\t0.查看堆中元素情况\n")
	fmt.Printf("\t1.插入元素\n")
	fmt.Printf("\t2.移除队首元素\n")
	fmt.Printf("\t3.查看队首元素\n")
	fmt.Printf("\t4.退出\n")
	for loop {
		fmt.Println("请输入操作命令所对应的数字：")
		fmt.Scanf("%d\r\n", &op)
		switch op {
		case 0:
			output(q)
		case 1:
			fmt.Println("请输入要插入的数字：")
			fmt.Scanf("%d\r\n", &val)
			push(q, val)
		case 2:
			pop(q)
		case 3:
			fmt.Println(top(q))
		case 4:
			loop = false

		}
	}
}

func top(q *Heap) (int,error) {
	if q.n==0 {
		return 0,errors.New("data is empty")
	}
	return q.data[0],nil
}

// 出队操作
// 先覆盖元素，再向下调整
func pop(q *Heap) {
	if q.n == 0 {
		return
	}
	q.data[0] = q.data[q.n-1]
	q.n--
	q.data = q.data[:q.n]
	fmt.Println("data is:", q.data)
	downUpdate(q, 0)
}

// 堆的向下调整操作
// ind 指向当前需要向下调整的节点
// temp 指向两个子节点中值较小的节点
// ind * 2 <= q->n 说明当前节点还有子节点
func downUpdate(q *Heap, ind int) {
	for (ind+1)*2 <= q.n {
		temp := ind*2 + 1
		if temp+1 < q.n && q.data[temp+1] < q.data[temp] {
			temp = temp + 1
		}
		if q.data[ind] <= q.data[temp] {
			break
		}
		q.data[ind], q.data[temp] = q.data[temp], q.data[ind]
	}
}

// 入队操作
// 先放置元素，再向上调整
func push(q *Heap, element int) bool {
	if q.n == q.size {
		return false
	}

	q.n++
	q.data[q.n-1] = element
	upUpdate(q, q.n-1)
	return true
}

// 堆的向上调整操作
// ind 指向当前需要向上调整的节点
// father 指向 ind 的父节点
func upUpdate(q *Heap, ind int) {
	var father int
	for ind != 0 {
		father = (ind - 1) / 2
		if q.data[father] <= q.data[ind] {
			break
		}
		q.data[ind], q.data[father] = q.data[father], q.data[ind]
	}
}

func output(q *Heap) {
	fmt.Println("data is", q.data, "n is:", q.n)
}

// 初始化一个最大容量为 size 的优先队列
func newHeap(size int) *Heap {
	q := new(Heap)
	q.size = size
	q.data = make([]int, size)
	q.n = 0
	return q
}
