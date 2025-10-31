package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil || pos < 0 {
		return nil
	}

	currentNode := l
	index := 0

	for currentNode != nil {
		if index == pos {
			return currentNode
		}
		currentNode = currentNode.Next
		index++
	}
	return nil
}
