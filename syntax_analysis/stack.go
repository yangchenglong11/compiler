/*
 * Revision History:
 *     Initial: 2018/01/18        Wang RiYu
 */

package main

import "bytes"

type Stack string

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
  return string(stack[len(stack)-1])
}

func (stack *Stack) Pop() string {
  theStack := *stack
  if theStack.IsEmpty() {
    return ""
  }
  value := theStack[len(theStack)-1]
  *stack = theStack[:len(theStack)-1]
  return string(value)
}

func (stack *Stack) Replace(start, end int, substring string) {
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

func (stack Stack) Reverse() string {
  var result Stack
  for i := stack.Len(); i > 0; i-- {
    result.Push(stack.Pop())
  }

  return string(result)
}

func (stack Stack) Index(index int) string {
  return string(string(stack)[index])
}
