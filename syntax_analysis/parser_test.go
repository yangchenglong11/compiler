/*
 * Revision History:
 *     Initial: 2018/01/21        Wang RiYu
 */

package syntax

import (
  "testing"
)

func TestParser(t *testing.T) {
  var (
    parser Parser
    /* 文法
      E’ -> #E#
      E -> E + T | T
      T -> T * F | F
      F -> P ↑ F | P
      P -> (E) | i
    */
    grammar = map[string]string{// 每一个 k, v 对应表达式 value -> key
      "E#E": "E’",
      "E + T": "E",
      "T": "E",
      "T * F": "T",
      "F": "T",
      "P ↑ F": "F",
      "P": "F",
      "(E)": "P",
      "i": "P",
    }
    relation = [][]string{
      {Ab, Be, Be, Be, Be, Ab, Ab},
      {Ab, Ab, Be, Be, Be, Ab, Ab},
      {Ab, Ab, Be, Be, Be, Ab, Ab},
      {Ab, Ab, Ab, "", "", Ab, Ab},
      {Be, Be, Be, Be, Be, Eq, ""},
      {Ab, Ab, Ab, "", "", Ab, Ab},
      {Be, Be, Be, Be, Be, "", Eq},
    }
    Vt = []string{"+", "*", "↑", "i", "(", ")", "#"}
    Vn = []string{"E’", "E", "T", "F", "P"}
  )
  parser.Init(grammar, Vt, Vn, relation)

  /* input1 */
  stack1 := Stack{"#"}
  input1 := Stack{}
  for _, v := range []string{"i", "+", "i", "*", "i", "#"} { // 可规约
    input1.Push(v)
  }
  result, err := parser.Analysis(&stack1, &input1)
  if err != nil {
    t.Error(err)
  }
  if !result {
    t.Error("归约失败")
  }

  /* input2 */
  stack2 := Stack{"#"}
  input2 := Stack{}
  for _, v := range []string{"(", "i", "*", "i", ")", "(", ")", "#"} { // 不可规约
    input2.Push(v)
  }
  result, _ = parser.Analysis(&stack2, &input2)
  if result {
    t.Error("归约出错")
  }
}
