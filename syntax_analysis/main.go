/*
 * Revision History:
 *     Initial: 2018/01/21        Wang RiYu
 */

package main

import (
  "fmt"
  "strings"
  "os"
  "bufio"
)

func main() {
  var (
    input      string
    inputStack Stack
    stack      Stack
    parser     Parser
    /* 文法
      E -> E + T | T
      T -> T * F | F
      F -> P ↑ F | P
      P -> (E) | i
    */
    grammar = map[string]string{// 每一个 k, v 对应表达式 value -> key
      "E+T": "E",
      "T": "E",
      "T*F": "T",
      "F": "T",
      "P↑F": "F",
      "P": "F",
      "(E)": "P",
      "i": "P",
    }
    relation = [][]string{// 算符优先关系
      {Above, Below, Below, Below, Below, Above, Above},
      {Above, Above, Below, Below, Below, Above, Above},
      {Above, Above, Below, Below, Below, Above, Above},
      {Above, Above, Above, "   ", "   ", Above, Above},
      {Below, Below, Below, Below, Below, Equal, "   "},
      {Above, Above, Above, "   ", "   ", Above, Above},
      {Below, Below, Below, Below, Below, "   ", Equal},
    }
    Vt = []string{"+", "*", "↑", "i", "(", ")", "#"} // 终结符
    Vn = []string{"E", "T", "F", "P"}                // 非终结符
  )
  parser.Init(grammar, Vt, Vn, relation)
  parser.DisplayGrammar()
  parser.DisplayRelationTable()

  fmt.Println("\n输入语句, 以#结束, 每个终结符以空格隔开, 例如 i + i * i #")
  reader := bufio.NewReader(os.Stdin)
  strBytes, _, err := reader.ReadLine()
  if err != nil {
    fmt.Println(err)
  } else {
    input = string(strBytes)
    split := strings.Split(input, " ")
    if split[len(split)-1] != "#" {
      fmt.Println("unvalid input")
    } else {
      for i := range split {
        inputStack.Push(split[i])
      }
      stack.Push("#")

      result, err := parser.Analysis(&stack, &inputStack)
      if err != nil {
        fmt.Println(err)
      } else {
        if result {
          fmt.Println("归约分析成功")
        } else {
          fmt.Println("归约分析失败")
        }
      }
    }
  }
}
