package miscellaneous

import "fmt"

/**
 * 设计一个栈，除了提供push（压栈），pop（出栈），peak（取栈顶元*素）操作以外，
 * 还能提供max（取栈中最大值）的功能，并使得时间复杂度最小
 */

// Stack stack struct
type Stack struct {
	Slice       []int
	SortedSlice []int
}

// Top return the top element of the stack
func (s *Stack) Top() int {
	if len(s.Slice) == 0 {
		return 0
	}
	return s.Slice[len(s.Slice)-1]
}

// Pop pop the top element of the stack
func (s *Stack) Pop() int {
	if len(s.Slice) == 0 {
		return 0
	}
	stackLen := len(s.Slice)
	top := s.Slice[stackLen-1]
	s.Slice = s.Slice[:stackLen-1]
	s.SortedSlice = s.SortedSlice[:stackLen-1]
	return top
}

// Push push a element to the top
// element of the stack
func (s *Stack) Push(val int) {
	s.Slice = append(s.Slice, val)
	if len(s.SortedSlice) == 0 {
		s.SortedSlice = append(s.SortedSlice, val)
	} else {
		// 关键：只与最上元素进行对比
		// 若小于，则push最上元素
		// Min函数与此类似
		top := s.SortedSlice[len(s.SortedSlice)-1]
		if val > top {
			s.SortedSlice = append(s.SortedSlice, val)
		} else {
			s.SortedSlice = append(s.SortedSlice, top)
		}
	}
}

// Max return the max element of the stack
func (s *Stack) Max() int {
	if len(s.Slice) == 0 {
		return 0
	}
	return s.SortedSlice[len(s.SortedSlice)-1]
}

// PrintElements print elements
func (s *Stack) PrintElements() {
	fmt.Println(s.Slice)
}

// TODO:
// 栈为空的情况需要考虑返回error
// if len(s.Slice) == 0 {
// 	return 0
// }
