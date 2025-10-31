package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}

	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}

	if l.Head == nil {
		l.Tail = nil
		return
	}

	prev := l.Head
	current := l.Head.Next

	for current != nil {
		if current.Data == data_ref {
			prev.Next = current.Next
			if current.Next == nil {
				l.Tail = prev
			}
		} else {
			prev = current
		}
		current = current.Next
	}
}
