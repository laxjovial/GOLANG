package piscine

func ListSize(l *List) int {
	count := 0
	node := l.Head
	for node != nil {
		count++
		node = node.Next
	}
	return count
}
