type MyLinkedList struct {
	Head *ListNode
	Tail *ListNode
}

type ListNode struct {
	Prev *ListNode
	Val  int
	Next *ListNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	// CHANGED: negative indexes are invalid
	if index < 0 {
		return -1
	}

	curr := this.Head

	for i := 0; i < index && curr != nil; i++ {
		curr = curr.Next
	}

	if curr == nil {
		return -1
	}

	return curr.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &ListNode{
		Prev: nil,
		Val:  val,
		Next: this.Head,
	}

	if this.Head != nil {
		this.Head.Prev = newNode
	} else {
		this.Tail = newNode
	}

	this.Head = newNode
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &ListNode{
		Prev: this.Tail,
		Val:  val,
		Next: nil,
	}

	if this.Tail != nil {
		this.Tail.Next = newNode
	} else {
		this.Head = newNode
	}

	this.Tail = newNode
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= 0 {
		this.AddAtHead(val)
		return
	}

	curr := this.Head
	i := 0

	for i < index && curr != nil {
		curr = curr.Next
		i++
	}

	if i < index {
		return
	}

	if curr == nil {
		this.AddAtTail(val)
		return
	}

	newNode := &ListNode{
		Val:  val,
		Prev: curr.Prev,
		Next: curr,
	}

	curr.Prev.Next = newNode
	curr.Prev = newNode
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 {
		return
	}

	curr := this.Head

	for i := 0; i < index && curr != nil; i++ {
		curr = curr.Next
	}

	if curr == nil {
		return
	}

	if curr.Prev == nil {
		this.Head = curr.Next
	} else {
		curr.Prev.Next = curr.Next
	}

	if curr.Next == nil {
		this.Tail = curr.Prev
	} else {
		curr.Next.Prev = curr.Prev
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */