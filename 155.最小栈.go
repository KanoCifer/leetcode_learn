package leetcodelearn

type MinStack struct {
	stack []int
	min   []int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		min:   []int{},
	}
}

func (this *MinStack) Push(value int) {
	this.stack = append(this.stack, value)
	if len(this.min) == 0 || value <= this.min[len(this.min)-1] {
		this.min = append(this.min, value)
	}

}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	if len(this.min) == 0 {
		return
	}
	pop := this.stack[len(this.stack)-1]
	if pop == this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.min) == 0 {
		return 0
	}
	return this.min[len(this.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(value);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
