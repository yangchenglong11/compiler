/*
 * Revision History:
 *     Initial: 2018/01/18        Wang RiYu
 */

package syntax

import (
  "fmt"
  "strings"
)

type Stack []string

func (stack Stack) Len() int {
  return len(stack)
}

func (stack Stack) IsEmpty() bool {
  return len(stack) == 0
}

func (stack Stack) ToString() string {
  return strings.Join(stack, " ")
}

func (stack *Stack) Index(index int) string {
  if index < 0 || index > stack.Len() {
    fmt.Println("out of range in stack")
    return ""
  }
  return Stack(*stack)[index]
}

func (stack *Stack) Push(value string) {
  *stack = append(*stack, value)
}

func (stack Stack) Top() string {
  if stack.IsEmpty() {
    return ""
  }
  return string(stack[stack.Len()-1])
}

func (stack *Stack) Pop() string {
  theStack := *stack
  if theStack.IsEmpty() {
    return ""
  }
  value := theStack[stack.Len()-1]
  *stack = theStack[:stack.Len()-1]
  return value
}

func (stack Stack) Left() string {
  if stack.IsEmpty() {
    return ""
  }
  return stack[0]
}

func (stack *Stack) Shift() string {
  theStack := *stack
  if theStack.IsEmpty() {
    return ""
  }
  value := theStack[0]
  *stack = theStack[1:]
  return value
}

func (stack *Stack) Replace(start, end int, substring string) {
  if start > end || start < 0 || end < 0 {
    fmt.Println("unvalid params")
    return
  }

  var theStack = *stack
  *stack = append([]string{}, theStack[:start]...)
  *stack = append(*stack, substring)
  if end < theStack.Len()-1 {
    *stack = append(*stack, theStack[end+1:]...)
  }
}
