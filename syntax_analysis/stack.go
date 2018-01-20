/*
 * Revision History:
 *     Initial: 2018/01/18        Wang RiYu
 */

package main

import (
  "bytes"
  "fmt"
)

type Stack []rune

func (stack Stack) Len() int {
  return len(stack)
}

func (stack Stack) IsEmpty() bool {
  return len(stack) == 0
}

func (stack *Stack) Push(value string) {
  var buf bytes.Buffer
  buf.WriteString(string(*stack))
  buf.WriteString(value)
  *stack = Stack(buf.String())
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
  return string(value)
}

func (stack Stack) Left() string {
  if stack.IsEmpty() {
    return ""
  }
  return stack.Index(0)
}

func (stack *Stack) Shift() string {
  theStack := *stack
  if theStack.IsEmpty() {
    return ""
  }
  value := theStack[0:1]
  *stack = theStack[1:]
  return string(value)
}

func (stack *Stack) ToString() string {
  return string(*stack)
}

func (stack *Stack) Replace(start, end int, substring string) {
  if start > end || start < 0 || end < 0 {
    fmt.Println("unvalid params")
    return
  }
  var (
    buf      bytes.Buffer
    theStack = *stack
  )
  buf.WriteString(string(theStack[:start]))
  buf.WriteString(substring)
  if end < stack.Len()-1 {
    buf.WriteString(string(theStack[end+1:]))
  }
  *stack = Stack(buf.String())
}

func (stack Stack) Index(index int) string {
  return string(stack[index])
}
