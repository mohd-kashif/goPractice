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

type DetectSquares struct {
	PointsMap map[[2]int]int
}

func Constructor() DetectSquares {
	return DetectSquares{
		PointsMap: make(map[[2]int]int),
	}
}

func (this *DetectSquares) Add(point []int) {
	pt := [2]int{point[0], point[1]}
	val, ok := this.PointsMap[pt]
	if ok {
		this.PointsMap[pt] = val + 1
	} else {
		this.PointsMap[pt] = 1
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (this *DetectSquares) Count(point []int) int {
	px := point[0]
	py := point[1]
	x := px
	y := py
	res := 0
	for pt, freq := range this.PointsMap {
		x = pt[0]
		y = pt[1]
		if (abs(x-px) != abs(y-py)) || (px == x && py == y) {
			continue
		}
		freq1, ok1 := this.PointsMap[[2]int{x, py}]
		freq2, ok2 := this.PointsMap[[2]int{px, y}]
		if ok1 && ok2 {
			res += freq * freq1 * freq2
		}
	}
	return res
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */
