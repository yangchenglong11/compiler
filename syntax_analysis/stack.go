/*
 * Revision History:
 *     Initial: 2018/01/18        Wang RiYu
 */

package syntax_analysis

import (
  "fmt"
  "strings"
)

var TokenN = Token{Output: "N"}

type Token struct {
  Label  int
  Name   string
  Code   int
  Addr   int
  Output string
}

type Stack []Token

func (stack Stack) Len() int {
  return len(stack)
}

func (stack Stack) IsEmpty() bool {
  return len(stack) == 0
}

func (stack Stack) ToString() string {
  var result []string
  for _, v := range stack {
    result = append(result, v.Output)
  }

  return strings.Join(result, " ")
}

func (stack *Stack) Index(index int) string {
  if index < 0 || index > stack.Len() {
    fmt.Println("out of range in stack")
    return ""
  }
  return Stack(*stack)[index].Output
}

func (stack *Stack) Push(t Token) {
  *stack = append(*stack, t)
}

func (stack Stack) Top() string {
  if stack.IsEmpty() {
    return ""
  }
  return stack[stack.Len()-1].Output
}

func (stack *Stack) Pop() Token {
  theStack := *stack
  if theStack.IsEmpty() {
    return Token{}
  }
  t := theStack[stack.Len()-1]
  *stack = theStack[:stack.Len()-1]
  return t
}

func (stack Stack) Left() string {
  if stack.IsEmpty() {
    return ""
  }
  return stack[0].Output
}

func (stack *Stack) Shift() Token {
  theStack := *stack
  if theStack.IsEmpty() {
    return Token{}
  }
  t := theStack[0]
  *stack = theStack[1:]
  return t
}

func (stack *Stack) Replace(start, end int, sub Token) {
  if start > end || start < 0 || end < 0 {
    fmt.Println("unvalid index")
    return
  }

  var theStack = *stack
  *stack = append([]Token{}, theStack[:start]...)
  *stack = append(*stack, sub)
  if end < theStack.Len()-1 {
    *stack = append(*stack, theStack[end+1:]...)
  }
}
