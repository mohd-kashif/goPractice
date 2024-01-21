package stack

import "strconv"

const LIMIT = 4 // DONOT CHANGE IT!

type Stack struct {
	ix   int // first free position, so data[ix] == 0
	data [LIMIT]int
}

func (st *Stack) Push(n int) {
	if st.ix == LIMIT {
		return
	}
	st.data[st.ix] = n
	st.ix++
}

func (st *Stack) Pop() int {
	if st.ix == 0 {
		return 0
	}
	st.ix--
	val := st.data[st.ix]
	st.data[st.ix] = 0
	return val
}

func (st Stack) String() string {
	result := ""
	for k, v := range st.data {
		result = result + "[" + strconv.Itoa(k) + ":" + strconv.Itoa(v) + "] "
	}
	return result
}
