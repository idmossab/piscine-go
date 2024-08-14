package piscine

func ListReverse(l *List) {
	prev := new(NodeL)
	prev = nil
	current := l.Head
	next := new(NodeL)
	next = nil
	l.Head = l.Tail

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Tail = prev
}
