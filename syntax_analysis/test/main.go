/*
 * Revision History:
 *     Initial: 2018/01/22        Wang RiYu
 */

package main

import (
  "fmt"
  "strings"
  "os"
  "bufio"

  "github.com/yangchenglong11/compiler/syntax_analysis"
)

func main() {
  var (
    Ab = syntax.Ab
    Be = syntax.Be
    Eq = syntax.Eq
    input      string
    inputStack syntax.Stack
    stack      syntax.Stack
    parser     syntax.Parser
    /* 文法
      E’ -> #E#
      E -> E + T | T
      T -> T * F | F
      F -> P ↑ F | P
      P -> (E) | i
    */
    //grammar = map[string]string{// 每一个 k, v 对应表达式 value -> key
    //  "E+T": "E",
    //  "T": "E",
    //  "T*F": "T",
    //  "F": "T",
    //  "P↑F": "F",
    //  "P": "F",
    //  "(E)": "P",
    //  "i": "P",
    //}
    //relation = [][]string{// 算符优先关系
    //  {Ab, Be, Be, Be, Be, Ab, Ab},
    //  {Ab, Ab, Be, Be, Be, Ab, Ab},
    //  {Ab, Ab, Be, Be, Be, Ab, Ab},
    //  {Ab, Ab, Ab, "", "", Ab, Ab},
    //  {Be, Be, Be, Be, Be, Eq, ""},
    //  {Ab, Ab, Ab, "", "", Ab, Ab},
    //  {Be, Be, Be, Be, Be, "", Eq},
    //}
    //Vt = []string{"+", "*", "↑", "i", "(", ")", "#"} // 终结符
    //Vn = []string{"E’", "E", "T", "F", "P"}          // 非终结符

    /*
      P’ -> #P#
      P -> program L
      L -> S | id, L | id
      D -> L : K
      K -> int | bool | real
      S -> id := E | if B then S else S | while B do S | begin L end | var D;
      B -> id R id
      R -> < | >
      E -> I + I | I - I
      I -> i | id
    */
    grammar = map[string]string{
      "#P#": "P’",
      "program L": "P",
      "S":             "L",
      "id, L":           "L",
      "id":             "L",
      "id := E":         "S",
      "if B then S else S": "S",
      "while B do S":     "S",
      "begin L end":     "S",
      "var D;":          "S",
      "L: K":          "D",
      "integer":           "K",
      "bool":          "K",
      "real":          "K",
      "id R id": "B",
      "<": "R",
      ">": "R",
      "i": "I",
      "id’": "I",
      "I + I": "E",
      "I - I": "E",
    }
    relation = [][]string{
      {Eq, Be, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
      {Ab, "", Be, "", "", Be, "", "", Be, "", Be, "", Be, "", "", "", "", "", "", "", "", "", ""},
      {Ab, "", Eq, Eq, Eq, "", Ab, Ab, "", Ab, "", Ab, "", "", Ab, "", "", "", Be, Be, Ab, Ab, ""},
      {Ab, "", Be, "", "", Be, "", "", Be, "", Be, Ab, Be, "", Ab, "", "", "", "", "", "", "", ""},
      {Ab, "", Be, "", "", "", "", Ab, "", "", "", Ab, "", "", Ab, "", "", "", "", "", Be, Be, Be},
      {"", "", Be, "", "", "", Eq, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
      {"", "", Be, "", "", Be, "", Eq, Be, "", Be, "", Be, "", "", "", "", "", "", "", "", "", ""},
      {Ab, "", Be, "", "", Be, "", Ab, Be, "", Be, Ab, Be, "", Ab, "", "", "", "", "", "", "", ""},
      {"", "", Be, "", "", "", "", "", "", Eq, "", "", "", "", "", "", "", "", "", "", "", "", ""},
      {Ab, "", Be, "", "", Be, "", Ab, Be, "", Be, Ab, Be, "", Ab, "", "", "", "", "", "", "", ""},
      {"", "", Be, "", "", Be, "", "", Be, "", Be, Eq, Be, "", "", "", "", "", "", "", "", "", ""},
      {Ab, "", "", "", "", "", "", Ab, "", "", "", Ab, "", "", Ab, "", "", "", "", "", "", "", ""},
      {"", "", "", "", "", "", "", "", "", "", "", "", "", Eq, "", "", "", "", "", "", "", "", ""},
      {Ab, "", "", "", "", "", "", Ab, "", "", "", Ab, "", "", Ab, "", "", "", "", "", "", "", ""},
      {"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", Be, Be, Be, "", "", "", "", ""},
      {"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
      {"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
      {"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
      {"", "", Ab, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
      {"", "", Ab, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
      {Ab, "", Be, "", "", "", "", Ab, "", "", "", Ab, "", "", Ab, "", "", "", "", "", "", "", Be},
      {Ab, "", Be, "", "", "", "", Ab, "", "", "", Ab, "", "", Ab, "", "", "", "", "", "", "", Be},
      {Ab, "", "", "", "", "", "", Ab, "", "", "", Ab, "", "", Ab, "", "", "", "", "", Ab, Ab, ""},
    }
    Vt = []string{"#", "program", "id", ",", ":=", "if", "then", "else", "while", "do", "begin", "end", "var", ";", ":", "integer", "bool", "real", "<", ">", "+", "-", "i"}
    Vn = []string{"P’", "P", "L", "S", "D", "K", "B", "R", "E", "I"}
  )
  parser.Init(grammar, Vt, Vn, relation)
  parser.DisplayGrammar()
  parser.DisplayRelationTable()

  fmt.Println("\n输入语句, 以#结束, 每个终结符以空格隔开, 例如 program begin if id > id then id := id - i else id := i end #:")
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
