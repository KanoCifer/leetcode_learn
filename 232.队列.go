package leetcodelearn

type MyQueue struct {
	stackIn  []int
	stackOut []int
}

func Constructor() MyQueue {
	return MyQueue{
		stackIn:  make([]int, 0),
		stackOut: make([]int, 0),
	}
}

func (q *MyQueue) Push(x int) {
	q.stackIn = append(q.stackIn, x)
}

func (q *MyQueue) transfer() {
	if len(q.stackOut) == 0 {
		for len(q.stackIn) > 0 {
			n := len(q.stackIn)
			top := q.stackIn[n-1]
			q.stackIn = q.stackIn[:n-1]
			q.stackOut = append(q.stackOut, top)
		}
	}
}

func (q *MyQueue) Pop() int {
	q.transfer()
	n := len(q.stackOut)
	if n == 0 {
		return 0 // 实际工程里应返回 error
	}
	v := q.stackOut[n-1]
	q.stackOut = q.stackOut[:n-1]
	return v
}

func (q *MyQueue) Peek() int {
	q.transfer()
	n := len(q.stackOut)
	if n == 0 {
		return 0
	}
	return q.stackOut[n-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.stackIn) == 0 && len(q.stackOut) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
