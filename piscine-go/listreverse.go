package piscine

func ListReverse(l *List) {
	var prev *NodeL
	currentNode := l.Head

	l.Tail = l.Head

	for currentNode != nil {
		next := currentNode.Next
		currentNode.Next = prev
		prev = currentNode
		currentNode = next
	}

	l.Head = prev
}
