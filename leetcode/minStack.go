package leetcode

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		minStack: make([]int, 0),
		stack:    make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	n := len(this.minStack)
	if n == 0 {
		this.minStack = append(this.minStack, val)
		return
	}
	if val < this.minStack[n-1] {
		this.minStack = append(this.minStack, val)
	} else {
		this.minStack = append(this.minStack, this.minStack[n-1])
	}

}

func (this *MinStack) Pop() {
	n := len(this.minStack)
	this.stack = this.stack[:n-1]
	this.minStack = this.minStack[:n-1]

}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
