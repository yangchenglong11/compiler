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
      L’ -> #L#
      L -> S | i,L | i
      S -> id := E | if B then S else S | while B do S | begin L end | var D
      D -> L:K; | L:K;D
      K -> int | bool | real
    */
    grammar = map[string]string{
      "S": "L",
      "i,L": "L",
      "i": "L",
      "id:=E": "S",
      "ifBthenSelseS": "S",
      "whileBdoS": "S",
      "beginLend": "S",
      "varD": "S",
      "L:K;": "D",
      "L:K;D": "D",
      "int": "K",
      "bool": "K",
      "real": "K",
    }
    relation = [][]string{
      //# ; i , id := if then else while do begin end var : int bool real
      {Eq, "", Be, "", Be, "", Be, "", "", Be, "", Be, "", Be, "", "", "", ""}, // #
      {"", "", Be, "", Be, "", Be, "", Ab, Be, "", Be, "", Be, Be, "", "", ""}, // ;
      {Ab, "", "", Eq, "", "", "", "", "", "", "", "", Ab, "", Ab, "", "", ""}, // i
      {Ab, "", Be, "", Be, "", Be, "", "", Be, "", Be, Ab, Be, Ab, "", "", ""}, // ,
      {"", "", "", "", "", Eq, "", "", "", "", "", "", "", "", "", "", "", ""}, // id
      {Ab, "", "", "", "", "", "", "", Ab, "", "", "", Ab, "", Ab, "", "", ""}, // :=
      {"", "", "", "", "", "", "", Eq, "", "", "", "", "", "", "", "", "", ""}, // if
      {"", "", "", "", Be, "", Be, "", Eq, Be, "", Be, "", Be, "", "", "", ""}, // then
      {Ab, "", "", "", Be, "", Be, "", Ab, Be, "", Be, Ab, Be, Ab, "", "", ""}, // else
      {"", "", "", "", "", "", "", "", "", "", Eq, "", "", "", "", "", "", ""}, // while
      {Ab, "", "", "", Be, "", Be, "", Ab, Be, "", Be, Ab, Be, Ab, "", "", ""}, // do
      {"", "", Be, "", Be, "", Be, "", "", Be, "", Be, Eq, Be, "", "", "", ""}, // begin
      {Ab, "", "", "", "", "", "", "", Ab, "", "", "", Ab, "", Ab, "", "", ""}, // end
      {Ab, "", "", "", "", "", "", "", Ab, "", "", "", Ab, "", Ab, "", "", ""}, // var
      {"", Eq, "", "", "", "", "", "", "", "", "", "", "", "", "", Be, Be, Be}, // :
      {"", Ab, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}, // int
      {"", Ab, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}, // bool
      {"", Ab, "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""}, // real
    }
    Vt = []string{"#", ";", "i", ",", "id", ":=", "if", "then", "else", "while", "do", "begin", "end", "var", ":", "int", "bool", "real"}
    Vn = []string{"L’", "L", "S", "D", "K"}
  )
  parser.Init(grammar, Vt, Vn, relation)
  parser.DisplayGrammar()
  parser.DisplayRelationTable()

  fmt.Println("\n输入语句, 以#结束, 每个终结符以空格隔开, 例如 if a>b then a++ else b++ #:")
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
