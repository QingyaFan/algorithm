package miscellaneous

import (
	"fmt"
	"testing"
)

func TestStackPush(t *testing.T) {
	// var slice, sortedSlice []int
	stack := Stack{}
	stack.Push(2)
	fmt.Println(stack.Top())
}

func TestStackMax(t *testing.T) {
	stack := Stack{}
	stack.Push(2)
	stack.Push(1)
	stack.Push(9)
	fmt.Println(stack.Max())
}
