type MinStack struct {
	stack []int
	mStack []int
}

func Constructor() MinStack {
	return MinStack{stack: []int{}, mStack: []int{}}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.mStack) == 0 || val <= this.mStack[len(this.mStack)-1]{
		this.mStack = append(this.mStack, val)
	}
}

func (this *MinStack) Pop() {
	top := this.stack[len(this.stack)-1]
	if top == this.mStack[len(this.mStack)-1]{
		this.mStack = this.mStack[:len(this.mStack)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.mStack[len(this.mStack)-1]
}
