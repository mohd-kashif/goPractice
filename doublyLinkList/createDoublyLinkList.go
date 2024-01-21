package doublylinklist

import "container/list"

func InsertListElements(n int) *list.List { // add elements in list from 1 to n
	lst := list.New()
	for i := 1; i <= n; i++ {
		lst.PushBack(i) // insertion here
	}
	return lst
}
