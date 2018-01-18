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
  if len(stack) == 0 {
    return ""
  }
  return string(stack[len(stack)-1])
}

func (stack *Stack) Pop() string {
  theStack := *stack
  if len(theStack) == 0 {
    return ""
  }
  value := theStack[len(theStack)-1]
  *stack = theStack[:len(theStack)-1]
  return string(value)
}

