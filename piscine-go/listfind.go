package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	for current := l.Head; current != nil; current = current.Next {
		if comp(current.Data, ref) {
			return &current.Data
		}
	}
	return nil
}
